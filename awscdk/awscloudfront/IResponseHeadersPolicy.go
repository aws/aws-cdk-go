package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront"
)

// Represents a response headers policy.
type IResponseHeadersPolicy interface {
	interfacesawscloudfront.IResponseHeadersPolicyRef
	// The ID of the response headers policy.
	ResponseHeadersPolicyId() *string
}

// The jsii proxy for IResponseHeadersPolicy
type jsiiProxy_IResponseHeadersPolicy struct {
	internal.Type__interfacesawscloudfrontIResponseHeadersPolicyRef
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

