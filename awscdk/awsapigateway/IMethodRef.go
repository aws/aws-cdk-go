package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Method.
// Experimental.
type IMethodRef interface {
	constructs.IConstruct
	// A reference to a Method resource.
	// Experimental.
	MethodRef() *MethodReference
}

// The jsii proxy for IMethodRef
type jsiiProxy_IMethodRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMethodRef) MethodRef() *MethodReference {
	var returns *MethodReference
	_jsii_.Get(
		j,
		"methodRef",
		&returns,
	)
	return returns
}

