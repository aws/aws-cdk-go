package awsefs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MountTarget.
// Experimental.
type IMountTargetRef interface {
	constructs.IConstruct
	// A reference to a MountTarget resource.
	// Experimental.
	MountTargetRef() *MountTargetReference
}

// The jsii proxy for IMountTargetRef
type jsiiProxy_IMountTargetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMountTargetRef) MountTargetRef() *MountTargetReference {
	var returns *MountTargetReference
	_jsii_.Get(
		j,
		"mountTargetRef",
		&returns,
	)
	return returns
}

