package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelQualityJobDefinition.
// Experimental.
type IModelQualityJobDefinitionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ModelQualityJobDefinition resource.
	// Experimental.
	ModelQualityJobDefinitionRef() *ModelQualityJobDefinitionReference
}

// The jsii proxy for IModelQualityJobDefinitionRef
type jsiiProxy_IModelQualityJobDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IModelQualityJobDefinitionRef) ModelQualityJobDefinitionRef() *ModelQualityJobDefinitionReference {
	var returns *ModelQualityJobDefinitionReference
	_jsii_.Get(
		j,
		"modelQualityJobDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModelQualityJobDefinitionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModelQualityJobDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

