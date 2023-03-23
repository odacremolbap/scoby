// Copyright 2023 TriggerMesh Inc.
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"

	corev1 "k8s.io/api/core/v1"
)

func TestNewContainer(t *testing.T) {
	testCases := map[string]struct {
		options  []ContainerOption
		expected corev1.Container
	}{
		"basic": {
			expected: corev1.Container{
				Name:  tName,
				Image: tImage,
			}},
		"with environment": {
			options: []ContainerOption{
				ContainerAddEnvFromValue("key", "value"),
			},
			expected: corev1.Container{
				Name:  tName,
				Image: tImage,
				Env: []corev1.EnvVar{
					{Name: "key",
						Value: "value",
					}},
			}},
		"with args": {
			options: []ContainerOption{
				ContainerAddArgs("--key1=value1 --key2 value2 --porcelain"),
			},
			expected: corev1.Container{
				Name:  tName,
				Image: tImage,
				Args:  []string{"--key1=value1", "--key2", "value2", "--porcelain"},
			}},
		"with port": {
			options: []ContainerOption{
				ContainerAddPort("port8888", 8888),
				ContainerAddPort("port9999", 9999),
			},
			expected: corev1.Container{
				Name:  tName,
				Image: tImage,
				Ports: []corev1.ContainerPort{
					{
						Name:          "port8888",
						ContainerPort: 8888,
					},
					{
						Name:          "port9999",
						ContainerPort: 9999,
					},
				},
			}},
		"with volume mount": {
			options: []ContainerOption{
				ContainerAddVolumeMount(
					NewVolumeMount(tName, tVolumeMountFile),
				),
			},
			expected: corev1.Container{
				Name:  tName,
				Image: tImage,
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:      tName,
						MountPath: tVolumeMountFile,
					},
				},
			}},
		"with environment variable": {
			options: []ContainerOption{
				ContainerAddEnv(
					&corev1.EnvVar{
						Name:  "test-env",
						Value: "test-value",
					}),
			},
			expected: corev1.Container{
				Name:  tName,
				Image: tImage,
				Env: []corev1.EnvVar{
					{
						Name:  "test-env",
						Value: "test-value",
					},
				},
			}},
		"with env from secret": {
			options: []ContainerOption{
				ContainerAddEnvVarFromSecret("test-env", "test-secret", "test-key"),
			},
			expected: corev1.Container{
				Name:  tName,
				Image: tImage,
				Env: []corev1.EnvVar{
					{
						Name: "test-env",
						ValueFrom: &corev1.EnvVarSource{
							SecretKeyRef: &corev1.SecretKeySelector{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "test-secret",
								},
								Key: "test-key",
							},
						},
					}},
			}},
		"with env from configmap": {
			options: []ContainerOption{
				ContainerAddEnvVarFromConfigMap("test-env", "test-secret", "test-key"),
			},
			expected: corev1.Container{
				Name:  tName,
				Image: tImage,
				Env: []corev1.EnvVar{
					{
						Name: "test-env",
						ValueFrom: &corev1.EnvVarSource{
							ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "test-secret",
								},
								Key: "test-key",
							},
						},
					}},
			}},
		"with image pull policy": {
			options: []ContainerOption{
				ContainerWithImagePullPolicy(corev1.PullAlways),
			},
			expected: corev1.Container{
				Name:            tName,
				Image:           tImage,
				ImagePullPolicy: corev1.PullAlways,
			}},
		"with termination message policy": {
			options: []ContainerOption{
				ContainerWithTerminationMessagePolicy(corev1.TerminationMessageReadFile),
			},
			expected: corev1.Container{
				Name:                     tName,
				Image:                    tImage,
				TerminationMessagePolicy: corev1.TerminationMessageReadFile,
			}},
		"with security context": {
			options: []ContainerOption{
				ContainerWithSecurityContext(
					NewSecurityContext(
						SecurityContextWithPrivilegeEscalation(false),
						SecurityContextWithReadOnlyRootFilesystem(true),
					)),
			},
			expected: corev1.Container{
				Name:  tName,
				Image: tImage,
				SecurityContext: &corev1.SecurityContext{
					AllowPrivilegeEscalation: &tFalse,
					ReadOnlyRootFilesystem:   &tTrue,
				},
			}},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := NewContainer(tName, tImage, tc.options...)
			assert.Equal(t, &tc.expected, got)
		})
	}
}
