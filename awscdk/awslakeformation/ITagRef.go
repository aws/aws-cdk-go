package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Tag.
// Experimental.
type ITagRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Tag resource.
	// Experimental.
	TagRef() *TagReference
}

// The jsii proxy for ITagRef
type jsiiProxy_ITagRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITagRef) TagRef() *TagReference {
	var returns *TagReference
	_jsii_.Get(
		j,
		"tagRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITagRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITagRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

