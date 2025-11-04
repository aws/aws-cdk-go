package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RestoreTestingPlan.
// Experimental.
type IRestoreTestingPlanRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RestoreTestingPlan resource.
	// Experimental.
	RestoreTestingPlanRef() *RestoreTestingPlanReference
}

// The jsii proxy for IRestoreTestingPlanRef
type jsiiProxy_IRestoreTestingPlanRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IRestoreTestingPlanRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestoreTestingPlanRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

