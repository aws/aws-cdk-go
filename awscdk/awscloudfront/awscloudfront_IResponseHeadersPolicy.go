package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a response headers policy.
type IResponseHeadersPolicy interface {
	// The ID of the response headers policy.
	ResponseHeadersPolicyId() *string
}

// The jsii proxy for IResponseHeadersPolicy
type jsiiProxy_IResponseHeadersPolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_IResponseHeadersPolicy) ResponseHeadersPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyId",
		&returns,
	)
	return returns
}

