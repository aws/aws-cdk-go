package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RestoreTestingSelection.
// Experimental.
type IRestoreTestingSelectionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RestoreTestingSelection resource.
	// Experimental.
	RestoreTestingSelectionRef() *RestoreTestingSelectionReference
}

// The jsii proxy for IRestoreTestingSelectionRef
type jsiiProxy_IRestoreTestingSelectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRestoreTestingSelectionRef) RestoreTestingSelectionRef() *RestoreTestingSelectionReference {
	var returns *RestoreTestingSelectionReference
	_jsii_.Get(
		j,
		"restoreTestingSelectionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestoreTestingSelectionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestoreTestingSelectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

