/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	coordinationv1 "k8s.io/api/coordination/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LeaseSpecApplyConfiguration represents a declarative configuration of the LeaseSpec type for use
// with apply.
type LeaseSpecApplyConfiguration struct {
	HolderIdentity       *string                                  `json:"holderIdentity,omitempty"`
	LeaseDurationSeconds *int32                                   `json:"leaseDurationSeconds,omitempty"`
	AcquireTime          *v1.MicroTime                            `json:"acquireTime,omitempty"`
	RenewTime            *v1.MicroTime                            `json:"renewTime,omitempty"`
	LeaseTransitions     *int32                                   `json:"leaseTransitions,omitempty"`
	Strategy             *coordinationv1.CoordinatedLeaseStrategy `json:"strategy,omitempty"`
	PreferredHolder      *string                                  `json:"preferredHolder,omitempty"`
}

// LeaseSpecApplyConfiguration constructs a declarative configuration of the LeaseSpec type for use with
// apply.
func LeaseSpec() *LeaseSpecApplyConfiguration {
	return &LeaseSpecApplyConfiguration{}
}

// WithHolderIdentity sets the HolderIdentity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HolderIdentity field is set to the value of the last call.
func (b *LeaseSpecApplyConfiguration) WithHolderIdentity(value string) *LeaseSpecApplyConfiguration {
	b.HolderIdentity = &value
	return b
}

// WithLeaseDurationSeconds sets the LeaseDurationSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LeaseDurationSeconds field is set to the value of the last call.
func (b *LeaseSpecApplyConfiguration) WithLeaseDurationSeconds(value int32) *LeaseSpecApplyConfiguration {
	b.LeaseDurationSeconds = &value
	return b
}

// WithAcquireTime sets the AcquireTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AcquireTime field is set to the value of the last call.
func (b *LeaseSpecApplyConfiguration) WithAcquireTime(value v1.MicroTime) *LeaseSpecApplyConfiguration {
	b.AcquireTime = &value
	return b
}

// WithRenewTime sets the RenewTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RenewTime field is set to the value of the last call.
func (b *LeaseSpecApplyConfiguration) WithRenewTime(value v1.MicroTime) *LeaseSpecApplyConfiguration {
	b.RenewTime = &value
	return b
}

// WithLeaseTransitions sets the LeaseTransitions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LeaseTransitions field is set to the value of the last call.
func (b *LeaseSpecApplyConfiguration) WithLeaseTransitions(value int32) *LeaseSpecApplyConfiguration {
	b.LeaseTransitions = &value
	return b
}

// WithStrategy sets the Strategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Strategy field is set to the value of the last call.
func (b *LeaseSpecApplyConfiguration) WithStrategy(value coordinationv1.CoordinatedLeaseStrategy) *LeaseSpecApplyConfiguration {
	b.Strategy = &value
	return b
}

// WithPreferredHolder sets the PreferredHolder field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreferredHolder field is set to the value of the last call.
func (b *LeaseSpecApplyConfiguration) WithPreferredHolder(value string) *LeaseSpecApplyConfiguration {
	b.PreferredHolder = &value
	return b
}
