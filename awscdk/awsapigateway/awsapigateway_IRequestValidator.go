package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigateway/internal"
)

// Experimental.
type IRequestValidator interface {
	awscdk.IResource
	// ID of the request validator, such as abc123.
	// Experimental.
	RequestValidatorId() *string
}

// The jsii proxy for IRequestValidator
type jsiiProxy_IRequestValidator struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRequestValidator) RequestValidatorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestValidatorId",
		&returns,
	)
	return returns
}

