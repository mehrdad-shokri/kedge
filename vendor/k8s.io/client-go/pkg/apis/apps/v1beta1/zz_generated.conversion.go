// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/client-go/pkg/api"
	api_v1 "k8s.io/client-go/pkg/api/v1"
	apps "k8s.io/client-go/pkg/apis/apps"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_StatefulSet_To_apps_StatefulSet,
		Convert_apps_StatefulSet_To_v1beta1_StatefulSet,
		Convert_v1beta1_StatefulSetList_To_apps_StatefulSetList,
		Convert_apps_StatefulSetList_To_v1beta1_StatefulSetList,
		Convert_v1beta1_StatefulSetSpec_To_apps_StatefulSetSpec,
		Convert_apps_StatefulSetSpec_To_v1beta1_StatefulSetSpec,
		Convert_v1beta1_StatefulSetStatus_To_apps_StatefulSetStatus,
		Convert_apps_StatefulSetStatus_To_v1beta1_StatefulSetStatus,
	)
}

func autoConvert_v1beta1_StatefulSet_To_apps_StatefulSet(in *StatefulSet, out *apps.StatefulSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_StatefulSetSpec_To_apps_StatefulSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_StatefulSetStatus_To_apps_StatefulSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1beta1_StatefulSet_To_apps_StatefulSet(in *StatefulSet, out *apps.StatefulSet, s conversion.Scope) error {
	return autoConvert_v1beta1_StatefulSet_To_apps_StatefulSet(in, out, s)
}

func autoConvert_apps_StatefulSet_To_v1beta1_StatefulSet(in *apps.StatefulSet, out *StatefulSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_apps_StatefulSetSpec_To_v1beta1_StatefulSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apps_StatefulSetStatus_To_v1beta1_StatefulSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_apps_StatefulSet_To_v1beta1_StatefulSet(in *apps.StatefulSet, out *StatefulSet, s conversion.Scope) error {
	return autoConvert_apps_StatefulSet_To_v1beta1_StatefulSet(in, out, s)
}

func autoConvert_v1beta1_StatefulSetList_To_apps_StatefulSetList(in *StatefulSetList, out *apps.StatefulSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]apps.StatefulSet, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_StatefulSet_To_apps_StatefulSet(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1beta1_StatefulSetList_To_apps_StatefulSetList(in *StatefulSetList, out *apps.StatefulSetList, s conversion.Scope) error {
	return autoConvert_v1beta1_StatefulSetList_To_apps_StatefulSetList(in, out, s)
}

func autoConvert_apps_StatefulSetList_To_v1beta1_StatefulSetList(in *apps.StatefulSetList, out *StatefulSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StatefulSet, len(*in))
		for i := range *in {
			if err := Convert_apps_StatefulSet_To_v1beta1_StatefulSet(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = make([]StatefulSet, 0)
	}
	return nil
}

func Convert_apps_StatefulSetList_To_v1beta1_StatefulSetList(in *apps.StatefulSetList, out *StatefulSetList, s conversion.Scope) error {
	return autoConvert_apps_StatefulSetList_To_v1beta1_StatefulSetList(in, out, s)
}

func autoConvert_v1beta1_StatefulSetSpec_To_apps_StatefulSetSpec(in *StatefulSetSpec, out *apps.StatefulSetSpec, s conversion.Scope) error {
	if err := v1.Convert_Pointer_int32_To_int32(&in.Replicas, &out.Replicas, s); err != nil {
		return err
	}
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	if err := api_v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.VolumeClaimTemplates = *(*[]api.PersistentVolumeClaim)(unsafe.Pointer(&in.VolumeClaimTemplates))
	out.ServiceName = in.ServiceName
	return nil
}

func autoConvert_apps_StatefulSetSpec_To_v1beta1_StatefulSetSpec(in *apps.StatefulSetSpec, out *StatefulSetSpec, s conversion.Scope) error {
	if err := v1.Convert_int32_To_Pointer_int32(&in.Replicas, &out.Replicas, s); err != nil {
		return err
	}
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	if err := api_v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.VolumeClaimTemplates = *(*[]api_v1.PersistentVolumeClaim)(unsafe.Pointer(&in.VolumeClaimTemplates))
	out.ServiceName = in.ServiceName
	return nil
}

func autoConvert_v1beta1_StatefulSetStatus_To_apps_StatefulSetStatus(in *StatefulSetStatus, out *apps.StatefulSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.Replicas = in.Replicas
	return nil
}

func Convert_v1beta1_StatefulSetStatus_To_apps_StatefulSetStatus(in *StatefulSetStatus, out *apps.StatefulSetStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_StatefulSetStatus_To_apps_StatefulSetStatus(in, out, s)
}

func autoConvert_apps_StatefulSetStatus_To_v1beta1_StatefulSetStatus(in *apps.StatefulSetStatus, out *StatefulSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.Replicas = in.Replicas
	return nil
}

func Convert_apps_StatefulSetStatus_To_v1beta1_StatefulSetStatus(in *apps.StatefulSetStatus, out *StatefulSetStatus, s conversion.Scope) error {
	return autoConvert_apps_StatefulSetStatus_To_v1beta1_StatefulSetStatus(in, out, s)
}
