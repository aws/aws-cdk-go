package awsiottwinmaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComponentType.
// Experimental.
type IComponentTypeRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ComponentType resource.
	// Experimental.
	ComponentTypeRef() *ComponentTypeReference
}

// The jsii proxy for IComponentTypeRef
type jsiiProxy_IComponentTypeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IComponentTypeRef) ComponentTypeRef() *ComponentTypeReference {
	var returns *ComponentTypeReference
	_jsii_.Get(
		j,
		"componentTypeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComponentTypeRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComponentTypeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

