package awsservicediscovery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Service.
// Experimental.
type IServiceRef interface {
	constructs.IConstruct
	// A reference to a Service resource.
	// Experimental.
	ServiceRef() *ServiceReference
}

// The jsii proxy for IServiceRef
type jsiiProxy_IServiceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServiceRef) ServiceRef() *ServiceReference {
	var returns *ServiceReference
	_jsii_.Get(
		j,
		"serviceRef",
		&returns,
	)
	return returns
}

