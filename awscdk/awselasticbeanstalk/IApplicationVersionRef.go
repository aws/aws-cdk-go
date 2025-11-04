package awselasticbeanstalk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticbeanstalk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationVersion.
// Experimental.
type IApplicationVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ApplicationVersion resource.
	// Experimental.
	ApplicationVersionRef() *ApplicationVersionReference
}

// The jsii proxy for IApplicationVersionRef
type jsiiProxy_IApplicationVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IApplicationVersionRef) ApplicationVersionRef() *ApplicationVersionReference {
	var returns *ApplicationVersionReference
	_jsii_.Get(
		j,
		"applicationVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

