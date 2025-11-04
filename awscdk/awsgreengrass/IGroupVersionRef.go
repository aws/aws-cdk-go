package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupVersion.
// Experimental.
type IGroupVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a GroupVersion resource.
	// Experimental.
	GroupVersionRef() *GroupVersionReference
}

// The jsii proxy for IGroupVersionRef
type jsiiProxy_IGroupVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IGroupVersionRef) GroupVersionRef() *GroupVersionReference {
	var returns *GroupVersionReference
	_jsii_.Get(
		j,
		"groupVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroupVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroupVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

