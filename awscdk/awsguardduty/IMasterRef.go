package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Master.
// Experimental.
type IMasterRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Master resource.
	// Experimental.
	MasterRef() *MasterReference
}

// The jsii proxy for IMasterRef
type jsiiProxy_IMasterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMasterRef) MasterRef() *MasterReference {
	var returns *MasterReference
	_jsii_.Get(
		j,
		"masterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMasterRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMasterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

