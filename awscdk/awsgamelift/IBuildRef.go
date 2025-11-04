package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Build.
// Experimental.
type IBuildRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Build resource.
	// Experimental.
	BuildRef() *BuildReference
}

// The jsii proxy for IBuildRef
type jsiiProxy_IBuildRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IBuildRef) BuildRef() *BuildReference {
	var returns *BuildReference
	_jsii_.Get(
		j,
		"buildRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuildRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuildRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

