// +build !ignore_autogenerated

// Copyright © 2019 Banzai Cloud
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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerConfig) DeepCopyInto(out *BrokerConfig) {
	*out = *in
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(v1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageConfigs != nil {
		in, out := &in.StorageConfigs, &out.StorageConfigs
		*out = make([]StorageConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerConfig.
func (in *BrokerConfig) DeepCopy() *BrokerConfig {
	if in == nil {
		return nil
	}
	out := new(BrokerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerState) DeepCopyInto(out *BrokerState) {
	*out = *in
	out.GracefulActionState = in.GracefulActionState
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerState.
func (in *BrokerState) DeepCopy() *BrokerState {
	if in == nil {
		return nil
	}
	out := new(BrokerState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CruiseControlConfig) DeepCopyInto(out *CruiseControlConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CruiseControlConfig.
func (in *CruiseControlConfig) DeepCopy() *CruiseControlConfig {
	if in == nil {
		return nil
	}
	out := new(CruiseControlConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyConfig) DeepCopyInto(out *EnvoyConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyConfig.
func (in *EnvoyConfig) DeepCopy() *EnvoyConfig {
	if in == nil {
		return nil
	}
	out := new(EnvoyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalListenerConfig) DeepCopyInto(out *ExternalListenerConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalListenerConfig.
func (in *ExternalListenerConfig) DeepCopy() *ExternalListenerConfig {
	if in == nil {
		return nil
	}
	out := new(ExternalListenerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GracefulActionState) DeepCopyInto(out *GracefulActionState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GracefulActionState.
func (in *GracefulActionState) DeepCopy() *GracefulActionState {
	if in == nil {
		return nil
	}
	out := new(GracefulActionState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalListenerConfig) DeepCopyInto(out *InternalListenerConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalListenerConfig.
func (in *InternalListenerConfig) DeepCopy() *InternalListenerConfig {
	if in == nil {
		return nil
	}
	out := new(InternalListenerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaCluster) DeepCopyInto(out *KafkaCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaCluster.
func (in *KafkaCluster) DeepCopy() *KafkaCluster {
	if in == nil {
		return nil
	}
	out := new(KafkaCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaClusterList) DeepCopyInto(out *KafkaClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaClusterList.
func (in *KafkaClusterList) DeepCopy() *KafkaClusterList {
	if in == nil {
		return nil
	}
	out := new(KafkaClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaClusterSpec) DeepCopyInto(out *KafkaClusterSpec) {
	*out = *in
	in.ListenersConfig.DeepCopyInto(&out.ListenersConfig)
	if in.ZKAddresses != nil {
		in, out := &in.ZKAddresses, &out.ZKAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RackAwareness != nil {
		in, out := &in.RackAwareness, &out.RackAwareness
		*out = new(RackAwareness)
		(*in).DeepCopyInto(*out)
	}
	if in.BrokerConfigs != nil {
		in, out := &in.BrokerConfigs, &out.BrokerConfigs
		*out = make([]BrokerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BrokerConfigGroups != nil {
		in, out := &in.BrokerConfigGroups, &out.BrokerConfigGroups
		*out = make(map[string]*BrokerConfig, len(*in))
		for key, val := range *in {
			var outVal *BrokerConfig
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(BrokerConfig)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	out.CruiseControlConfig = in.CruiseControlConfig
	out.EnvoyConfig = in.EnvoyConfig
	out.MonitoringConfig = in.MonitoringConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaClusterSpec.
func (in *KafkaClusterSpec) DeepCopy() *KafkaClusterSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaClusterStatus) DeepCopyInto(out *KafkaClusterStatus) {
	*out = *in
	if in.BrokersState != nil {
		in, out := &in.BrokersState, &out.BrokersState
		*out = make(map[int32]*BrokerState, len(*in))
		for key, val := range *in {
			var outVal *BrokerState
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(BrokerState)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaClusterStatus.
func (in *KafkaClusterStatus) DeepCopy() *KafkaClusterStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenersConfig) DeepCopyInto(out *ListenersConfig) {
	*out = *in
	if in.ExternalListeners != nil {
		in, out := &in.ExternalListeners, &out.ExternalListeners
		*out = make([]ExternalListenerConfig, len(*in))
		copy(*out, *in)
	}
	if in.InternalListeners != nil {
		in, out := &in.InternalListeners, &out.InternalListeners
		*out = make([]InternalListenerConfig, len(*in))
		copy(*out, *in)
	}
	if in.SSLSecrets != nil {
		in, out := &in.SSLSecrets, &out.SSLSecrets
		*out = new(SSLSecrets)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenersConfig.
func (in *ListenersConfig) DeepCopy() *ListenersConfig {
	if in == nil {
		return nil
	}
	out := new(ListenersConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConfig) DeepCopyInto(out *MonitoringConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConfig.
func (in *MonitoringConfig) DeepCopy() *MonitoringConfig {
	if in == nil {
		return nil
	}
	out := new(MonitoringConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackAwareness) DeepCopyInto(out *RackAwareness) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackAwareness.
func (in *RackAwareness) DeepCopy() *RackAwareness {
	if in == nil {
		return nil
	}
	out := new(RackAwareness)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSLSecrets) DeepCopyInto(out *SSLSecrets) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSLSecrets.
func (in *SSLSecrets) DeepCopy() *SSLSecrets {
	if in == nil {
		return nil
	}
	out := new(SSLSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfig) DeepCopyInto(out *StorageConfig) {
	*out = *in
	if in.PVCSpec != nil {
		in, out := &in.PVCSpec, &out.PVCSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfig.
func (in *StorageConfig) DeepCopy() *StorageConfig {
	if in == nil {
		return nil
	}
	out := new(StorageConfig)
	in.DeepCopyInto(out)
	return out
}
