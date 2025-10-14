package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceAction.
// Experimental.
type IServiceActionRef interface {
	constructs.IConstruct
	// A reference to a ServiceAction resource.
	// Experimental.
	ServiceActionRef() *ServiceActionReference
}

// The jsii proxy for IServiceActionRef
type jsiiProxy_IServiceActionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServiceActionRef) ServiceActionRef() *ServiceActionReference {
	var returns *ServiceActionReference
	_jsii_.Get(
		j,
		"serviceActionRef",
		&returns,
	)
	return returns
}

