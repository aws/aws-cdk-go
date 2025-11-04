package awsmemorydb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmemorydb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ParameterGroup.
// Experimental.
type IParameterGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ParameterGroup resource.
	// Experimental.
	ParameterGroupRef() *ParameterGroupReference
}

// The jsii proxy for IParameterGroupRef
type jsiiProxy_IParameterGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IParameterGroupRef) ParameterGroupRef() *ParameterGroupReference {
	var returns *ParameterGroupReference
	_jsii_.Get(
		j,
		"parameterGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameterGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameterGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

