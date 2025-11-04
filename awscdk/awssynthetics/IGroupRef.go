package awssynthetics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssynthetics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Group.
// Experimental.
type IGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Group resource.
	// Experimental.
	GroupRef() *GroupReference
}

// The jsii proxy for IGroupRef
type jsiiProxy_IGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IGroupRef) GroupRef() *GroupReference {
	var returns *GroupReference
	_jsii_.Get(
		j,
		"groupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

