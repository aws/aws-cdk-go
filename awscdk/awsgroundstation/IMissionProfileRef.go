package awsgroundstation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgroundstation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MissionProfile.
// Experimental.
type IMissionProfileRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MissionProfile resource.
	// Experimental.
	MissionProfileRef() *MissionProfileReference
}

// The jsii proxy for IMissionProfileRef
type jsiiProxy_IMissionProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMissionProfileRef) MissionProfileRef() *MissionProfileReference {
	var returns *MissionProfileReference
	_jsii_.Get(
		j,
		"missionProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMissionProfileRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMissionProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

