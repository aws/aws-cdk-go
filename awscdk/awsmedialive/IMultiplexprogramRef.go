package awsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Multiplexprogram.
// Experimental.
type IMultiplexprogramRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Multiplexprogram resource.
	// Experimental.
	MultiplexprogramRef() *MultiplexprogramReference
}

// The jsii proxy for IMultiplexprogramRef
type jsiiProxy_IMultiplexprogramRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMultiplexprogramRef) MultiplexprogramRef() *MultiplexprogramReference {
	var returns *MultiplexprogramReference
	_jsii_.Get(
		j,
		"multiplexprogramRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiplexprogramRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiplexprogramRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

