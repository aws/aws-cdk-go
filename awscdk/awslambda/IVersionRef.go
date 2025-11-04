package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Version.
// Experimental.
type IVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Version resource.
	// Experimental.
	VersionRef() *VersionReference
}

// The jsii proxy for IVersionRef
type jsiiProxy_IVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IVersionRef) VersionRef() *VersionReference {
	var returns *VersionReference
	_jsii_.Get(
		j,
		"versionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

