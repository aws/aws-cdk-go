package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RequestValidator.
// Experimental.
type IRequestValidatorRef interface {
	constructs.IConstruct
	// A reference to a RequestValidator resource.
	// Experimental.
	RequestValidatorRef() *RequestValidatorReference
}

// The jsii proxy for IRequestValidatorRef
type jsiiProxy_IRequestValidatorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRequestValidatorRef) RequestValidatorRef() *RequestValidatorReference {
	var returns *RequestValidatorReference
	_jsii_.Get(
		j,
		"requestValidatorRef",
		&returns,
	)
	return returns
}

