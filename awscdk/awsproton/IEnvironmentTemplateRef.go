package awsproton

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsproton/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentTemplate.
// Experimental.
type IEnvironmentTemplateRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EnvironmentTemplate resource.
	// Experimental.
	EnvironmentTemplateRef() *EnvironmentTemplateReference
}

// The jsii proxy for IEnvironmentTemplateRef
type jsiiProxy_IEnvironmentTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEnvironmentTemplateRef) EnvironmentTemplateRef() *EnvironmentTemplateReference {
	var returns *EnvironmentTemplateReference
	_jsii_.Get(
		j,
		"environmentTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentTemplateRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

