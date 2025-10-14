package awsgroundstation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgroundstation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MissionProfile.
// Experimental.
type IMissionProfileRef interface {
	constructs.IConstruct
	// A reference to a MissionProfile resource.
	// Experimental.
	MissionProfileRef() *MissionProfileReference
}

// The jsii proxy for IMissionProfileRef
type jsiiProxy_IMissionProfileRef struct {
	internal.Type__constructsIConstruct
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

