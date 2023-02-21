package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Cache Policy.
type ICachePolicy interface {
	// The ID of the cache policy.
	CachePolicyId() *string
}

// The jsii proxy for ICachePolicy
type jsiiProxy_ICachePolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_ICachePolicy) CachePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyId",
		&returns,
	)
	return returns
}

