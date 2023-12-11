/*
Copyright 2023 The Godel Scheduler Authors.

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

package securitycontext

import (
	v1 "k8s.io/api/core/v1"
)

// HasPrivilegedRequest returns the value of SecurityContext.Privileged, taking into account
// the possibility of nils
func HasPrivilegedRequest(container *v1.Container) bool {
	if container.SecurityContext == nil {
		return false
	}
	if container.SecurityContext.Privileged == nil {
		return false
	}
	return *container.SecurityContext.Privileged
}

// HasCapabilitiesRequest returns true if Adds or Drops are defined in the security context
// capabilities, taking into account nils
func HasCapabilitiesRequest(container *v1.Container) bool {
	if container.SecurityContext == nil {
		return false
	}
	if container.SecurityContext.Capabilities == nil {
		return false
	}
	return len(container.SecurityContext.Capabilities.Add) > 0 || len(container.SecurityContext.Capabilities.Drop) > 0
}

// DetermineEffectiveSecurityContext returns a synthesized SecurityContext for reading effective configurations
// from the provided pod's and container's security context. Container's fields take precedence in cases where both
// are set
func DetermineEffectiveSecurityContext(pod *v1.Pod, container *v1.Container) *v1.SecurityContext {
	effectiveSc := securityContextFromPodSecurityContext(pod)
	containerSc := container.SecurityContext

	if effectiveSc == nil && containerSc == nil {
		return &v1.SecurityContext{}
	}
	if effectiveSc != nil && containerSc == nil {
		return effectiveSc
	}
	if effectiveSc == nil && containerSc != nil {
		return containerSc
	}

	if containerSc.SELinuxOptions != nil {
		effectiveSc.SELinuxOptions = new(v1.SELinuxOptions)
		*effectiveSc.SELinuxOptions = *containerSc.SELinuxOptions
	}

	if containerSc.WindowsOptions != nil {
		// only override fields that are set at the container level, not the whole thing
		if effectiveSc.WindowsOptions == nil {
			effectiveSc.WindowsOptions = &v1.WindowsSecurityContextOptions{}
		}
		if containerSc.WindowsOptions.GMSACredentialSpecName != nil || containerSc.WindowsOptions.GMSACredentialSpec != nil {
			// both GMSA fields go hand in hand
			effectiveSc.WindowsOptions.GMSACredentialSpecName = containerSc.WindowsOptions.GMSACredentialSpecName
			effectiveSc.WindowsOptions.GMSACredentialSpec = containerSc.WindowsOptions.GMSACredentialSpec
		}
		if containerSc.WindowsOptions.RunAsUserName != nil {
			effectiveSc.WindowsOptions.RunAsUserName = containerSc.WindowsOptions.RunAsUserName
		}
	}

	if containerSc.Capabilities != nil {
		effectiveSc.Capabilities = new(v1.Capabilities)
		*effectiveSc.Capabilities = *containerSc.Capabilities
	}

	if containerSc.Privileged != nil {
		effectiveSc.Privileged = new(bool)
		*effectiveSc.Privileged = *containerSc.Privileged
	}

	if containerSc.RunAsUser != nil {
		effectiveSc.RunAsUser = new(int64)
		*effectiveSc.RunAsUser = *containerSc.RunAsUser
	}

	if containerSc.RunAsGroup != nil {
		effectiveSc.RunAsGroup = new(int64)
		*effectiveSc.RunAsGroup = *containerSc.RunAsGroup
	}

	if containerSc.RunAsNonRoot != nil {
		effectiveSc.RunAsNonRoot = new(bool)
		*effectiveSc.RunAsNonRoot = *containerSc.RunAsNonRoot
	}

	if containerSc.ReadOnlyRootFilesystem != nil {
		effectiveSc.ReadOnlyRootFilesystem = new(bool)
		*effectiveSc.ReadOnlyRootFilesystem = *containerSc.ReadOnlyRootFilesystem
	}

	if containerSc.AllowPrivilegeEscalation != nil {
		effectiveSc.AllowPrivilegeEscalation = new(bool)
		*effectiveSc.AllowPrivilegeEscalation = *containerSc.AllowPrivilegeEscalation
	}

	if containerSc.ProcMount != nil {
		effectiveSc.ProcMount = new(v1.ProcMountType)
		*effectiveSc.ProcMount = *containerSc.ProcMount
	}

	return effectiveSc
}

func securityContextFromPodSecurityContext(pod *v1.Pod) *v1.SecurityContext {
	if pod.Spec.SecurityContext == nil {
		return nil
	}

	synthesized := &v1.SecurityContext{}

	if pod.Spec.SecurityContext.SELinuxOptions != nil {
		synthesized.SELinuxOptions = &v1.SELinuxOptions{}
		*synthesized.SELinuxOptions = *pod.Spec.SecurityContext.SELinuxOptions
	}

	if pod.Spec.SecurityContext.WindowsOptions != nil {
		synthesized.WindowsOptions = &v1.WindowsSecurityContextOptions{}
		*synthesized.WindowsOptions = *pod.Spec.SecurityContext.WindowsOptions
	}

	if pod.Spec.SecurityContext.RunAsUser != nil {
		synthesized.RunAsUser = new(int64)
		*synthesized.RunAsUser = *pod.Spec.SecurityContext.RunAsUser
	}

	if pod.Spec.SecurityContext.RunAsGroup != nil {
		synthesized.RunAsGroup = new(int64)
		*synthesized.RunAsGroup = *pod.Spec.SecurityContext.RunAsGroup
	}

	if pod.Spec.SecurityContext.RunAsNonRoot != nil {
		synthesized.RunAsNonRoot = new(bool)
		*synthesized.RunAsNonRoot = *pod.Spec.SecurityContext.RunAsNonRoot
	}

	return synthesized
}

// AddNoNewPrivileges returns if we should add the no_new_privs option.
func AddNoNewPrivileges(sc *v1.SecurityContext) bool {
	if sc == nil {
		return false
	}

	// handle the case where the user did not set the default and did not explicitly set allowPrivilegeEscalation
	if sc.AllowPrivilegeEscalation == nil {
		return false
	}

	// handle the case where defaultAllowPrivilegeEscalation is false or the user explicitly set allowPrivilegeEscalation to true/false
	return !*sc.AllowPrivilegeEscalation
}

var (
	// These *must* be kept in sync with moby/moby.
	// https://github.com/moby/moby/blob/master/oci/defaults.go#L116-L134
	// @jessfraz will watch changes to those files upstream.
	defaultMaskedPaths = []string{
		"/proc/acpi",
		"/proc/kcore",
		"/proc/keys",
		"/proc/latency_stats",
		"/proc/timer_list",
		"/proc/timer_stats",
		"/proc/sched_debug",
		"/proc/scsi",
		"/sys/firmware",
	}
	defaultReadonlyPaths = []string{
		"/proc/asound",
		"/proc/bus",
		"/proc/fs",
		"/proc/irq",
		"/proc/sys",
		"/proc/sysrq-trigger",
	}
)

// ConvertToRuntimeMaskedPaths converts the ProcMountType to the specified or default
// masked paths.
func ConvertToRuntimeMaskedPaths(opt *v1.ProcMountType) []string {
	if opt != nil && *opt == v1.UnmaskedProcMount {
		// Unmasked proc mount should have no paths set as masked.
		return []string{}
	}

	// Otherwise, add the default masked paths to the runtime security context.
	return defaultMaskedPaths
}

// ConvertToRuntimeReadonlyPaths converts the ProcMountType to the specified or default
// readonly paths.
func ConvertToRuntimeReadonlyPaths(opt *v1.ProcMountType) []string {
	if opt != nil && *opt == v1.UnmaskedProcMount {
		// Unmasked proc mount should have no paths set as readonly.
		return []string{}
	}

	// Otherwise, add the default readonly paths to the runtime security context.
	return defaultReadonlyPaths
}
