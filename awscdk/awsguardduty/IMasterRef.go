package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Master.
// Experimental.
type IMasterRef interface {
	constructs.IConstruct
	// A reference to a Master resource.
	// Experimental.
	MasterRef() *MasterReference
}

// The jsii proxy for IMasterRef
type jsiiProxy_IMasterRef struct {
	internal.Type__constructsIConstruct
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

