// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ThanosRulerStatusApplyConfiguration represents an declarative configuration of the ThanosRulerStatus type for use
// with apply.
type ThanosRulerStatusApplyConfiguration struct {
	Paused              *bool                         `json:"paused,omitempty"`
	Replicas            *int32                        `json:"replicas,omitempty"`
	UpdatedReplicas     *int32                        `json:"updatedReplicas,omitempty"`
	AvailableReplicas   *int32                        `json:"availableReplicas,omitempty"`
	UnavailableReplicas *int32                        `json:"unavailableReplicas,omitempty"`
	Conditions          []ConditionApplyConfiguration `json:"conditions,omitempty"`
}

// ThanosRulerStatusApplyConfiguration constructs an declarative configuration of the ThanosRulerStatus type for use with
// apply.
func ThanosRulerStatus() *ThanosRulerStatusApplyConfiguration {
	return &ThanosRulerStatusApplyConfiguration{}
}

// WithPaused sets the Paused field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Paused field is set to the value of the last call.
func (b *ThanosRulerStatusApplyConfiguration) WithPaused(value bool) *ThanosRulerStatusApplyConfiguration {
	b.Paused = &value
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *ThanosRulerStatusApplyConfiguration) WithReplicas(value int32) *ThanosRulerStatusApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithUpdatedReplicas sets the UpdatedReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UpdatedReplicas field is set to the value of the last call.
func (b *ThanosRulerStatusApplyConfiguration) WithUpdatedReplicas(value int32) *ThanosRulerStatusApplyConfiguration {
	b.UpdatedReplicas = &value
	return b
}

// WithAvailableReplicas sets the AvailableReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AvailableReplicas field is set to the value of the last call.
func (b *ThanosRulerStatusApplyConfiguration) WithAvailableReplicas(value int32) *ThanosRulerStatusApplyConfiguration {
	b.AvailableReplicas = &value
	return b
}

// WithUnavailableReplicas sets the UnavailableReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UnavailableReplicas field is set to the value of the last call.
func (b *ThanosRulerStatusApplyConfiguration) WithUnavailableReplicas(value int32) *ThanosRulerStatusApplyConfiguration {
	b.UnavailableReplicas = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ThanosRulerStatusApplyConfiguration) WithConditions(values ...*ConditionApplyConfiguration) *ThanosRulerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
