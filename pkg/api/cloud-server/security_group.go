/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - 混合云管理平台 (BlueKing - Hybrid Cloud Management System) available.
 * Copyright (C) 2022 THL A29 Limited,
 * a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 *
 * to the current version of the project delivered to anyone in the future.
 */

package cloudserver

import (
	"errors"
	"fmt"

	"hcm/pkg/api/core"
	"hcm/pkg/api/core/cloud"
	"hcm/pkg/criteria/constant"
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/criteria/validator"
	"hcm/pkg/dal/dao/types"
	"hcm/pkg/runtime/filter"
)

// -------------------------- List --------------------------

// SecurityGroupListReq security group list req.
type SecurityGroupListReq struct {
	Filter *filter.Expression `json:"filter" validate:"required"`
	Page   *types.BasePage    `json:"page" validate:"required"`
}

// Validate security group list request.
func (req *SecurityGroupListReq) Validate() error {
	return validator.Validate.Struct(req)
}

// SecurityGroupListResult define security group list result.
type SecurityGroupListResult struct {
	Count   uint64                    `json:"count,omitempty"`
	Details []cloud.BaseSecurityGroup `json:"details,omitempty"`
}

// SecurityGroupRuleListReq security group rule list req.
type SecurityGroupRuleListReq struct {
	Page *core.PageWithoutSort `json:"page" validate:"required"`
}

// Validate security group list request.
func (req *SecurityGroupRuleListReq) Validate() error {
	return validator.Validate.Struct(req)
}

// SecurityGroupRuleListResult define security group rule list result.
type SecurityGroupRuleListResult[T cloud.SecurityGroupRule] struct {
	Count   uint64 `json:"count,omitempty"`
	Details []T    `json:"details,omitempty"`
}

// -------------------------- Update --------------------------

// SecurityGroupUpdateReq security group update request.
type SecurityGroupUpdateReq struct {
	Name string  `json:"name"`
	Memo *string `json:"memo"`
}

// Validate security group update request.
func (req *SecurityGroupUpdateReq) Validate() error {
	if len(req.Name) == 0 && req.Memo == nil {
		return errors.New("name or memo is required")
	}

	if len(req.Name) != 0 {
		if err := validator.ValidateSecurityGroupName(req.Name); err != nil {
			return err
		}
	}

	if req.Memo != nil {
		if err := validator.ValidateSecurityGroupMemo(req.Memo); err != nil {
			return err
		}
	}

	return nil
}

// AssignSecurityGroupToBizReq define assign security group to biz req.
type AssignSecurityGroupToBizReq struct {
	BkBizID          int64    `json:"bk_biz_id" validate:"required"`
	SecurityGroupIDs []string `json:"security_group_ids" validate:"required"`
}

// Validate assign security group to biz request.
func (req *AssignSecurityGroupToBizReq) Validate() error {
	if err := validator.Validate.Struct(req); err != nil {
		return err
	}

	if req.BkBizID <= 0 {
		return errors.New("bk_biz_id should >= 0")
	}

	if len(req.SecurityGroupIDs) == 0 {
		return errors.New("security group ids is required")
	}

	if len(req.SecurityGroupIDs) > constant.BatchOperationMaxLimit {
		return fmt.Errorf("security group ids should <= %d", constant.BatchOperationMaxLimit)
	}

	return nil
}

// TCloudSGRuleUpdateReq define tcloud security group rule update req.
type TCloudSGRuleUpdateReq struct {
	Protocol                   *string `json:"protocol"`
	Port                       *string `json:"port"`
	IPv4Cidr                   *string `json:"ipv4_cidr"`
	IPv6Cidr                   *string `json:"ipv6_cidr"`
	CloudTargetSecurityGroupID *string `json:"cloud_target_security_group_id"`
	Action                     string  `json:"action"`
	Memo                       *string `json:"memo"`
}

// Validate tcloud security group rule update request.
func (req *TCloudSGRuleUpdateReq) Validate() error {
	return validator.Validate.Struct(req)
}

// AwsSGRuleUpdateReq define aws security group rule update req.
type AwsSGRuleUpdateReq struct {
	IPv4Cidr                   *string `json:"ipv4_cidr"`
	IPv6Cidr                   *string `json:"ipv6_cidr"`
	Memo                       *string `json:"memo"`
	FromPort                   int64   `json:"from_port"`
	ToPort                     int64   `json:"to_port"`
	Protocol                   *string `json:"protocol"`
	CloudTargetSecurityGroupID *string `json:"cloud_target_security_group_id"`
}

// Validate aws security group rule update request.
func (req *AwsSGRuleUpdateReq) Validate() error {
	return validator.Validate.Struct(req)
}

// AzureSGRuleUpdateReq define azure security group rule update req.
type AzureSGRuleUpdateReq struct {
	Name                             string    `json:"name"`
	Memo                             *string   `json:"memo"`
	DestinationAddressPrefix         *string   `json:"destination_address_prefix"`
	DestinationAddressPrefixes       []*string `json:"destination_address_prefixes"`
	CloudDestinationSecurityGroupIDs []*string `json:"cloud_destination_security_group_ids"`
	DestinationPortRange             *string   `json:"destination_port_range"`
	DestinationPortRanges            []*string `json:"destination_port_ranges"`
	Protocol                         string    `json:"protocol"`
	SourceAddressPrefix              *string   `json:"source_address_prefix"`
	SourceAddressPrefixes            []*string `json:"source_address_prefixes"`
	CloudSourceSecurityGroupIDs      []*string `json:"cloud_source_security_group_ids"`
	SourcePortRange                  *string   `json:"source_port_range"`
	SourcePortRanges                 []*string `json:"source_port_ranges"`
	Priority                         int32     `json:"priority"`
	Access                           string    `json:"access"`
}

// Validate azure security group rule update request.
func (req *AzureSGRuleUpdateReq) Validate() error {
	return validator.Validate.Struct(req)
}

// -------------------------- Delete --------------------------

// SecurityGroupBatchDeleteReq security group update request.
type SecurityGroupBatchDeleteReq struct {
	IDs []string `json:"ids" validate:"required"`
}

// Validate security group delete request.
func (req *SecurityGroupBatchDeleteReq) Validate() error {
	return validator.Validate.Struct(req)
}

// -------------------------- Create --------------------------

// SecurityGroupRuleCreateReq define security group rule create req.
type SecurityGroupRuleCreateReq[T SecurityGroupRule] struct {
	EgressRuleSet  []T `json:"egress_rule_set" validate:"omitempty"`
	IngressRuleSet []T `json:"ingress_rule_set" validate:"omitempty"`
}

// Validate security group rule create request.
func (req *SecurityGroupRuleCreateReq[T]) Validate() error {
	if len(req.EgressRuleSet) == 0 && len(req.IngressRuleSet) == 0 {
		return errors.New("egress rule or ingress rule is required")
	}

	if len(req.EgressRuleSet) != 0 && len(req.IngressRuleSet) != 0 {
		return errors.New("egress rule or ingress rule only one is allowed")
	}

	return nil
}

// SecurityGroupRule define security group rule when create.
type SecurityGroupRule interface {
	TCloudSecurityGroupRule | AwsSecurityGroupRule | HuaWeiSecurityGroupRule | AzureSecurityGroupRule
}

// TCloudSecurityGroupRule define tcloud security group rule spec.
type TCloudSecurityGroupRule struct {
	Protocol                   *string `json:"protocol"`
	Port                       *string `json:"port"`
	IPv4Cidr                   *string `json:"ipv4_cidr"`
	IPv6Cidr                   *string `json:"ipv6_cidr"`
	CloudTargetSecurityGroupID *string `json:"cloud_target_security_group_id"`
	Action                     string  `json:"action"`
	Memo                       *string `json:"memo"`
}

// AwsSecurityGroupRule define aws security group rule spec.
type AwsSecurityGroupRule struct {
	IPv4Cidr                   *string `json:"ipv4_cidr"`
	IPv6Cidr                   *string `json:"ipv6_cidr"`
	Memo                       *string `json:"memo"`
	FromPort                   int64   `json:"from_port"`
	ToPort                     int64   `json:"to_port"`
	Protocol                   *string `json:"protocol"`
	CloudTargetSecurityGroupID *string `json:"cloud_target_security_group_id"`
}

// HuaWeiSecurityGroupRule define huawei security group rule spec.
type HuaWeiSecurityGroupRule struct {
	Memo               *string `json:"memo"`
	Ethertype          *string `json:"ethertype"`
	Protocol           *string `json:"protocol"`
	RemoteIPPrefix     *string `json:"remote_ip_prefix"`
	CloudRemoteGroupID *string `json:"cloud_remote_group_id"`
	Port               *string `json:"port"`
	Action             *string `json:"action"`
	Priority           int64   `json:"priority"`
}

// AzureSecurityGroupRule define azure security group rule spec.
type AzureSecurityGroupRule struct {
	Name                             string    `json:"name"`
	Memo                             *string   `json:"memo"`
	DestinationAddressPrefix         *string   `json:"destination_address_prefix"`
	DestinationAddressPrefixes       []*string `json:"destination_address_prefixes"`
	CloudDestinationSecurityGroupIDs []*string `json:"cloud_destination_security_group_ids"`
	DestinationPortRange             *string   `json:"destination_port_range"`
	DestinationPortRanges            []*string `json:"destination_port_ranges"`
	Protocol                         string    `json:"protocol"`
	SourceAddressPrefix              *string   `json:"source_address_prefix"`
	SourceAddressPrefixes            []*string `json:"source_address_prefixes"`
	CloudSourceSecurityGroupIDs      []*string `json:"cloud_source_security_group_ids"`
	SourcePortRange                  *string   `json:"source_port_range"`
	SourcePortRanges                 []*string `json:"source_port_ranges"`
	Priority                         int32     `json:"priority"`
	// Type 更新时该字段无法更新。
	Type   enumor.SecurityGroupRuleType `json:"type"`
	Access string                       `json:"access"`
}
