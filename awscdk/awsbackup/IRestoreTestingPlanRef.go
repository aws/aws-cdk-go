package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RestoreTestingPlan.
// Experimental.
type IRestoreTestingPlanRef interface {
	constructs.IConstruct
	// A reference to a RestoreTestingPlan resource.
	// Experimental.
	RestoreTestingPlanRef() *RestoreTestingPlanReference
}

// The jsii proxy for IRestoreTestingPlanRef
type jsiiProxy_IRestoreTestingPlanRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRestoreTestingPlanRef) RestoreTestingPlanRef() *RestoreTestingPlanReference {
	var returns *RestoreTestingPlanReference
	_jsii_.Get(
		j,
		"restoreTestingPlanRef",
		&returns,
	)
	return returns
}

