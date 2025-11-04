package awsgreengrassv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrassv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComponentVersion.
// Experimental.
type IComponentVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ComponentVersion resource.
	// Experimental.
	ComponentVersionRef() *ComponentVersionReference
}

// The jsii proxy for IComponentVersionRef
type jsiiProxy_IComponentVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IComponentVersionRef) ComponentVersionRef() *ComponentVersionReference {
	var returns *ComponentVersionReference
	_jsii_.Get(
		j,
		"componentVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComponentVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComponentVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

