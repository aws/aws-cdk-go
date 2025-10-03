package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RestoreTestingPlan.
// Experimental.
type IRestoreTestingPlanRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRestoreTestingPlanRef
type jsiiProxy_IRestoreTestingPlanRef struct {
	internal.Type__constructsIConstruct
}

