package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceDefinition.
// Experimental.
type IResourceDefinitionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResourceDefinition resource.
	// Experimental.
	ResourceDefinitionRef() *ResourceDefinitionReference
}

// The jsii proxy for IResourceDefinitionRef
type jsiiProxy_IResourceDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IResourceDefinitionRef) ResourceDefinitionRef() *ResourceDefinitionReference {
	var returns *ResourceDefinitionReference
	_jsii_.Get(
		j,
		"resourceDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceDefinitionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

