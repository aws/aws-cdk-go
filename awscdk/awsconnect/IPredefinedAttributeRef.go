package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PredefinedAttribute.
// Experimental.
type IPredefinedAttributeRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PredefinedAttribute resource.
	// Experimental.
	PredefinedAttributeRef() *PredefinedAttributeReference
}

// The jsii proxy for IPredefinedAttributeRef
type jsiiProxy_IPredefinedAttributeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPredefinedAttributeRef) PredefinedAttributeRef() *PredefinedAttributeReference {
	var returns *PredefinedAttributeReference
	_jsii_.Get(
		j,
		"predefinedAttributeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPredefinedAttributeRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPredefinedAttributeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

