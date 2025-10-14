package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RestoreTestingSelection.
// Experimental.
type IRestoreTestingSelectionRef interface {
	constructs.IConstruct
	// A reference to a RestoreTestingSelection resource.
	// Experimental.
	RestoreTestingSelectionRef() *RestoreTestingSelectionReference
}

// The jsii proxy for IRestoreTestingSelectionRef
type jsiiProxy_IRestoreTestingSelectionRef struct {
	internal.Type__constructsIConstruct
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

