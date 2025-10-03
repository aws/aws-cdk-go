package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RestoreTestingSelection.
// Experimental.
type IRestoreTestingSelectionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRestoreTestingSelectionRef
type jsiiProxy_IRestoreTestingSelectionRef struct {
	internal.Type__constructsIConstruct
}

