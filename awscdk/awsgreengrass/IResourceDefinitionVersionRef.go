package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceDefinitionVersion.
// Experimental.
type IResourceDefinitionVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResourceDefinitionVersion resource.
	// Experimental.
	ResourceDefinitionVersionRef() *ResourceDefinitionVersionReference
}

// The jsii proxy for IResourceDefinitionVersionRef
type jsiiProxy_IResourceDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IResourceDefinitionVersionRef) ResourceDefinitionVersionRef() *ResourceDefinitionVersionReference {
	var returns *ResourceDefinitionVersionReference
	_jsii_.Get(
		j,
		"resourceDefinitionVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceDefinitionVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceDefinitionVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

