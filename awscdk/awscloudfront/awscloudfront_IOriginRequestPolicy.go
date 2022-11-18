package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Origin Request Policy.
type IOriginRequestPolicy interface {
	// The ID of the origin request policy.
	OriginRequestPolicyId() *string
}

// The jsii proxy for IOriginRequestPolicy
type jsiiProxy_IOriginRequestPolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_IOriginRequestPolicy) OriginRequestPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyId",
		&returns,
	)
	return returns
}

