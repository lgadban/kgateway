// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// RateLimitApplyConfiguration represents a declarative configuration of the RateLimit type for use
// with apply.
type RateLimitApplyConfiguration struct {
	Local  *LocalRateLimitPolicyApplyConfiguration `json:"local,omitempty"`
	Global *RateLimitPolicyApplyConfiguration      `json:"global,omitempty"`
}

// RateLimitApplyConfiguration constructs a declarative configuration of the RateLimit type for use with
// apply.
func RateLimit() *RateLimitApplyConfiguration {
	return &RateLimitApplyConfiguration{}
}

// WithLocal sets the Local field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Local field is set to the value of the last call.
func (b *RateLimitApplyConfiguration) WithLocal(value *LocalRateLimitPolicyApplyConfiguration) *RateLimitApplyConfiguration {
	b.Local = value
	return b
}

// WithGlobal sets the Global field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Global field is set to the value of the last call.
func (b *RateLimitApplyConfiguration) WithGlobal(value *RateLimitPolicyApplyConfiguration) *RateLimitApplyConfiguration {
	b.Global = value
	return b
}
