package main

// Provides output for use as an Ansible inventory plugin

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/goamz/ec2"
)

func newInventory(instances *ec2.InstancesResp) (*inventory, error) {
	meta := make(map[string]map[string]map[string]string)
	hostvars := make(map[string]map[string]string)
	meta["hostvars"] = hostvars
	var hosts []string

	inv := inventory{Meta: meta, Hosts: &hosts}

	for _, reservation := range instances.Reservations {
		for _, instance := range reservation.Instances {
			tags := tagsToMap(instance.Tags)
			label, ok := tags["Name"]
			if !ok {
				return &inv, fmt.Errorf("instance %s does not have a 'Name' tag", instance.InstanceId)
			}
			displayGroup := tags["DisplayGroup"]
			hosts = append(hosts, label)
			hostvars[label] = map[string]string{
				"ansible_ssh_host":    instance.PublicIpAddress,
				"host_label":          label,
				"host_aws_id":         instance.InstanceId,
				"host_display_group":  displayGroup,
				"host_arch":           instance.Architecture,
				"host_aws_type":       instance.InstanceType,
				"host_aws_avail_zone": instance.AvailZone,
				"host_private_ip":     instance.PrivateIpAddress,
				"host_public_ip":      instance.PublicIpAddress,
				"host_public_dns":     instance.DNSName,
				"host_private_dns":    instance.PrivateDNSName,
			}
		}
	}
	return &inv, nil
}

func tagsToMap(tags []ec2.Tag) map[string]string {
	m := make(map[string]string, len(tags))
	for _, tag := range tags {
		m[tag.Key] = tag.Value
	}
	return m
}

type inventory struct {
	Meta  map[string]map[string]map[string]string `json:"_meta"`
	Hosts *[]string                               `json:"hosts"`
}

func (i *inventory) toJSON() ([]byte, error) {
	return json.MarshalIndent(i, " ", "  ")
}
