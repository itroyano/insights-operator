// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	time "time"

	operatorv1 "github.com/openshift/api/operator/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ImagePrunerSpecApplyConfiguration represents a declarative configuration of the ImagePrunerSpec type for use
// with apply.
type ImagePrunerSpecApplyConfiguration struct {
	Schedule                     *string                      `json:"schedule,omitempty"`
	Suspend                      *bool                        `json:"suspend,omitempty"`
	KeepTagRevisions             *int                         `json:"keepTagRevisions,omitempty"`
	KeepYoungerThan              *time.Duration               `json:"keepYoungerThan,omitempty"`
	KeepYoungerThanDuration      *v1.Duration                 `json:"keepYoungerThanDuration,omitempty"`
	Resources                    *corev1.ResourceRequirements `json:"resources,omitempty"`
	Affinity                     *corev1.Affinity             `json:"affinity,omitempty"`
	NodeSelector                 map[string]string            `json:"nodeSelector,omitempty"`
	Tolerations                  []corev1.Toleration          `json:"tolerations,omitempty"`
	SuccessfulJobsHistoryLimit   *int32                       `json:"successfulJobsHistoryLimit,omitempty"`
	FailedJobsHistoryLimit       *int32                       `json:"failedJobsHistoryLimit,omitempty"`
	IgnoreInvalidImageReferences *bool                        `json:"ignoreInvalidImageReferences,omitempty"`
	LogLevel                     *operatorv1.LogLevel         `json:"logLevel,omitempty"`
}

// ImagePrunerSpecApplyConfiguration constructs a declarative configuration of the ImagePrunerSpec type for use with
// apply.
func ImagePrunerSpec() *ImagePrunerSpecApplyConfiguration {
	return &ImagePrunerSpecApplyConfiguration{}
}

// WithSchedule sets the Schedule field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Schedule field is set to the value of the last call.
func (b *ImagePrunerSpecApplyConfiguration) WithSchedule(value string) *ImagePrunerSpecApplyConfiguration {
	b.Schedule = &value
	return b
}

// WithSuspend sets the Suspend field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Suspend field is set to the value of the last call.
func (b *ImagePrunerSpecApplyConfiguration) WithSuspend(value bool) *ImagePrunerSpecApplyConfiguration {
	b.Suspend = &value
	return b
}

// WithKeepTagRevisions sets the KeepTagRevisions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeepTagRevisions field is set to the value of the last call.
func (b *ImagePrunerSpecApplyConfiguration) WithKeepTagRevisions(value int) *ImagePrunerSpecApplyConfiguration {
	b.KeepTagRevisions = &value
	return b
}

// WithKeepYoungerThan sets the KeepYoungerThan field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeepYoungerThan field is set to the value of the last call.
func (b *ImagePrunerSpecApplyConfiguration) WithKeepYoungerThan(value time.Duration) *ImagePrunerSpecApplyConfiguration {
	b.KeepYoungerThan = &value
	return b
}

// WithKeepYoungerThanDuration sets the KeepYoungerThanDuration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeepYoungerThanDuration field is set to the value of the last call.
func (b *ImagePrunerSpecApplyConfiguration) WithKeepYoungerThanDuration(value v1.Duration) *ImagePrunerSpecApplyConfiguration {
	b.KeepYoungerThanDuration = &value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *ImagePrunerSpecApplyConfiguration) WithResources(value corev1.ResourceRequirements) *ImagePrunerSpecApplyConfiguration {
	b.Resources = &value
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *ImagePrunerSpecApplyConfiguration) WithAffinity(value corev1.Affinity) *ImagePrunerSpecApplyConfiguration {
	b.Affinity = &value
	return b
}

// WithNodeSelector puts the entries into the NodeSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeSelector field,
// overwriting an existing map entries in NodeSelector field with the same key.
func (b *ImagePrunerSpecApplyConfiguration) WithNodeSelector(entries map[string]string) *ImagePrunerSpecApplyConfiguration {
	if b.NodeSelector == nil && len(entries) > 0 {
		b.NodeSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeSelector[k] = v
	}
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *ImagePrunerSpecApplyConfiguration) WithTolerations(values ...corev1.Toleration) *ImagePrunerSpecApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}

// WithSuccessfulJobsHistoryLimit sets the SuccessfulJobsHistoryLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SuccessfulJobsHistoryLimit field is set to the value of the last call.
func (b *ImagePrunerSpecApplyConfiguration) WithSuccessfulJobsHistoryLimit(value int32) *ImagePrunerSpecApplyConfiguration {
	b.SuccessfulJobsHistoryLimit = &value
	return b
}

// WithFailedJobsHistoryLimit sets the FailedJobsHistoryLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailedJobsHistoryLimit field is set to the value of the last call.
func (b *ImagePrunerSpecApplyConfiguration) WithFailedJobsHistoryLimit(value int32) *ImagePrunerSpecApplyConfiguration {
	b.FailedJobsHistoryLimit = &value
	return b
}

// WithIgnoreInvalidImageReferences sets the IgnoreInvalidImageReferences field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IgnoreInvalidImageReferences field is set to the value of the last call.
func (b *ImagePrunerSpecApplyConfiguration) WithIgnoreInvalidImageReferences(value bool) *ImagePrunerSpecApplyConfiguration {
	b.IgnoreInvalidImageReferences = &value
	return b
}

// WithLogLevel sets the LogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogLevel field is set to the value of the last call.
func (b *ImagePrunerSpecApplyConfiguration) WithLogLevel(value operatorv1.LogLevel) *ImagePrunerSpecApplyConfiguration {
	b.LogLevel = &value
	return b
}
