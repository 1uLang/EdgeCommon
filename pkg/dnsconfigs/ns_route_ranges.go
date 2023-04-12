// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build plus

package dnsconfigs

import (
	"encoding/json"
	"errors"
	"github.com/1uLang/EdgeCommon/pkg/configutils"
	"github.com/1uLang/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/maps"
	"net"
)

type NSRouteRangeType = string

const (
	NSRouteRangeTypeIP     NSRouteRangeType = "ipRange" // IP范围
	NSRouteRangeTypeCIDR   NSRouteRangeType = "cidr"    // CIDR
	NSRouteRangeTypeRegion NSRouteRangeType = "region"  // 区域
)

func AllNSRouteRangeTypes() []*shared.Definition {
	return []*shared.Definition{
		{
			Name: "IP范围",
			Code: NSRouteRangeTypeIP,
		},
		{
			Name: "CIDR",
			Code: NSRouteRangeTypeCIDR,
		},
		{
			Name: "区域",
			Code: NSRouteRangeTypeRegion,
		},
	}
}

// NSRouteRegionResolver 解析IP接口
type NSRouteRegionResolver interface {
	Resolve(ip net.IP) (countryId int64, provinceId int64, cityId int64, providerId int64)
}

// NSRouteRangeInterface 线路范围接口
type NSRouteRangeInterface interface {
	// Init 初始化
	Init() error

	// Contains 判断是否包含
	Contains(ip net.IP) bool

	// SetRegionResolver 设置IP解析接口
	SetRegionResolver(resolver NSRouteRegionResolver)

	// IsExcluding 是否为排除
	IsExcluding() bool
}

type NSBaseRouteRange struct {
	IsReverse bool `json:"isReverse"`

	routeRegionResolver NSRouteRegionResolver
}

func (this *NSBaseRouteRange) SetRegionResolver(resolver NSRouteRegionResolver) {
	this.routeRegionResolver = resolver
}

func (this *NSBaseRouteRange) IsExcluding() bool {
	return this.IsReverse
}

// NSRouteRangeIPRange IP范围配置
// IPv4和IPv6不能混用
type NSRouteRangeIPRange struct {
	NSBaseRouteRange

	IPFrom string `json:"ipFrom"`
	IPTo   string `json:"ipTo"`

	ipFromLong uint64
	ipToLong   uint64

	ipVersion int // 4|6
}

func (this *NSRouteRangeIPRange) Init() error {
	var ipFrom = net.ParseIP(this.IPFrom)
	var ipTo = net.ParseIP(this.IPTo)
	if ipFrom == nil {
		return errors.New("invalid ipFrom '" + this.IPFrom + "'")
	}
	if ipTo == nil {
		return errors.New("invalid ipTo '" + this.IPTo + "'")
	}

	var ipFromVersion = configutils.IPVersion(ipFrom)
	var ipToVersion = configutils.IPVersion(ipTo)
	if ipFromVersion != ipToVersion {
		return errors.New("ipFrom and ipTo version are not same")
	}
	this.ipVersion = ipFromVersion

	this.ipFromLong = configutils.IP2Long(ipFrom)
	this.ipToLong = configutils.IP2Long(ipTo)

	if this.ipFromLong > this.ipToLong {
		this.ipFromLong, this.ipToLong = this.ipToLong, this.ipFromLong
	}

	return nil
}

func (this *NSRouteRangeIPRange) Contains(netIP net.IP) bool {
	if len(netIP) == 0 {
		return false
	}

	var version = configutils.IPVersion(netIP)
	if version != this.ipVersion {
		return false
	}

	var ipLong = configutils.IP2Long(netIP)
	return ipLong >= this.ipFromLong && ipLong <= this.ipToLong
}

// NSRouteRangeCIDR CIDR范围配置
type NSRouteRangeCIDR struct {
	NSBaseRouteRange

	CIDR string `json:"cidr"`

	cidr *net.IPNet
}

func (this *NSRouteRangeCIDR) Init() error {
	_, ipNet, err := net.ParseCIDR(this.CIDR)
	if err != nil {
		return errors.New("parse cidr failed: " + err.Error())
	}

	this.cidr = ipNet

	return nil
}

func (this *NSRouteRangeCIDR) Contains(netIP net.IP) bool {
	if netIP == nil {
		return false
	}

	if this.cidr == nil {
		return false
	}

	return this.cidr.Contains(netIP)
}

// NSRouteRangeRegion 区域范围
// country:ID, province:ID, city:ID, isp:ID
type NSRouteRangeRegion struct {
	NSBaseRouteRange

	Regions []*routeRegion `json:"regions"`
}

func (this *NSRouteRangeRegion) Init() error {
	return nil
}

func (this *NSRouteRangeRegion) Contains(netIP net.IP) bool {
	if this.routeRegionResolver == nil {
		return false
	}

	if len(this.Regions) == 0 {
		return false
	}

	countryId, provinceId, cityId, providerId := this.routeRegionResolver.Resolve(netIP)
	if countryId <= 0 && provinceId <= 0 && cityId <= 0 && providerId <= 0 {
		return false
	}

	for _, region := range this.Regions {
		if region.Id <= 0 {
			continue
		}

		switch region.Type {
		case "country":
			if region.Id == countryId {
				return true
			}
		case "province":
			if region.Id == provinceId {
				return true
			}
		case "city":
			if region.Id == cityId {
				return true
			}
		case "isp":
			if region.Id == providerId {
				return true
			}
		}
	}

	return false
}

type routeRegion struct {
	Type string `json:"type"` // country|province|city|isp
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// InitNSRangesFromJSON 从JSON中初始化线路范围
func InitNSRangesFromJSON(rangesJSON []byte) (ranges []NSRouteRangeInterface, err error) {
	if len(rangesJSON) == 0 {
		return
	}

	var rangeMaps = []maps.Map{}
	err = json.Unmarshal(rangesJSON, &rangeMaps)
	if err != nil {
		return nil, err
	}
	for _, rangeMap := range rangeMaps {
		var rangeType = rangeMap.GetString("type")
		paramsJSON, err := json.Marshal(rangeMap.Get("params"))
		if err != nil {
			return nil, err
		}

		var r NSRouteRangeInterface

		switch rangeType {
		case NSRouteRangeTypeIP:
			r = &NSRouteRangeIPRange{}
		case NSRouteRangeTypeCIDR:
			r = &NSRouteRangeCIDR{}
		case NSRouteRangeTypeRegion:
			r = &NSRouteRangeRegion{}
		default:
			return nil, errors.New("invalid route line type '" + rangeType + "'")
		}

		err = json.Unmarshal(paramsJSON, r)
		if err != nil {
			return nil, err
		}
		err = r.Init()
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, r)
	}
	return
}
