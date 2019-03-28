// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import (
	"encoding/json"
	"strings"
)

// Response[]UAP from controller
var response struct {
	Data []json.RawMessage
	Meta meta
}

type StaMap map[string]Sta

// Station data
type Sta struct {
	u                *Unifi
	ID               string `json:"_id"`
	IsGuestByUsw     bool   `json:"_is_guest_by_usw,omitempty"`
	LastSeenByUsw    int    `json:"_last_seen_by_usw,omitempty"`
	UptimeByUsw      int    `json:"_uptime_by_usw,omitempty"`
	AssocTime        int    `json:"assoc_time"`
	FirstSeen        int64  `json:"first_seen"`
	IP               string `json:"ip"`
	IsGuest          bool   `json:"is_guest"`
	IsWired          bool   `json:"is_wired"`
	LastSeen         int64  `json:"last_seen"`
	LatestAssocTime  int    `json:"latest_assoc_time"`
	Mac              string `json:"mac"`
	Network          string `json:"network,omitempty"`
	NetworkID        string `json:"network_id,omitempty"`
	Oui              string `json:"oui"`
	SiteID           string `json:"site_id"`
	SwDepth          int    `json:"sw_depth,omitempty"`
	SwMac            string `json:"sw_mac,omitempty"`
	SwPort           int    `json:"sw_port,omitempty"`
	Uptime           int    `json:"uptime"`
	UserID           string `json:"user_id"`
	Hostname         string `json:"hostname,omitempty"`
	IsGuestByUap     bool   `json:"_is_guest_by_uap,omitempty"`
	LastSeenByUap    int    `json:"_last_seen_by_uap,omitempty"`
	RoamCount        int    `json:"roam_count,omitempty"`
	UptimeByUap      int    `json:"_uptime_by_uap,omitempty"`
	ApMac            string `json:"ap_mac,omitempty"`
	Authorized       bool   `json:"authorized,omitempty"`
	BSSID            string `json:"bssid,omitempty"`
	BytesR           int    `json:"bytes-r,omitempty"`
	Ccq              int    `json:"ccq,omitempty"`
	Channel          int    `json:"channel,omitempty"`
	ESSID            string `json:"essid,omitempty"`
	Idletime         int    `json:"idletime,omitempty"`
	Is11R            bool   `json:"is_11r,omitempty"`
	Noise            int    `json:"noise,omitempty"`
	PowersaveEnabled bool   `json:"powersave_enabled,omitempty"`
	QosPolicyApplied bool   `json:"qos_policy_applied,omitempty"`
	Radio            string `json:"radio,omitempty"`
	RadioProto       string `json:"radio_proto,omitempty"`
	Rssi             int    `json:"rssi,omitempty"`
	RxBytes          int    `json:"rx_bytes,omitempty"`
	RxBytesR         int    `json:"rx_bytes-r,omitempty"`
	RxPackets        int    `json:"rx_packets,omitempty"`
	RxRate           int    `json:"rx_rate,omitempty"`
	Signal           int    `json:"signal,omitempty"`
	TxBytes          int    `json:"tx_bytes,omitempty"`
	TxBytesR         int    `json:"tx_bytes-r,omitempty"`
	TxPackets        int    `json:"tx_packets,omitempty"`
	TxPower          int    `json:"tx_power,omitempty"`
	TxRate           int    `json:"tx_rate,omitempty"`
	Vlan             int    `json:"vlan,omitempty"`
}

// Returns a station name
func (s Sta) Name() string {
	if s.Hostname != "" {
		return s.Hostname
	}
	if s.IP != "" {
		return s.IP
	}
	return s.Mac
}

func (s Sta) Block(site *Site) error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.parse(site, "block-sta", command{Mac: s.Mac}, &response)
}

func (s Sta) UnBlock(site *Site) error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.parse(site, "unblock-sta", command{Mac: s.Mac}, &response)
}

func (s Sta) Disconnect(site *Site) error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.parse(site, "kick-sta", command{Mac: s.Mac}, &response)
}

func (s Sta) AuthorizeGuest(site *Site, minutes, down, up, mbytes *int64, apMac *string) error {
	if s.u == nil {
		return ErrLoginFirst
	}

	// Prepare command
	payload := command{Mac: s.Mac}

	if minutes != nil {
		payload.Minutes = *minutes
	}
	if down != nil {
		payload.Down = *down
	}
	if up != nil {
		payload.Up = *up

	}
	if mbytes != nil {
		payload.MBytes = *mbytes
	}
	if apMac != nil {
		payload.ApMac = strings.ToLower(*apMac)
	}

	return s.u.parse(site, "authorize-guest", payload, &response)
}

func (s Sta) UnauthorizeGuest(site *Site) error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.parse(site, "unauthorize-guest", command{Mac: s.Mac}, &response)
}
