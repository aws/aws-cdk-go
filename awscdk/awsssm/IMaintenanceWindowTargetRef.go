package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MaintenanceWindowTarget.
// Experimental.
type IMaintenanceWindowTargetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMaintenanceWindowTargetRef
type jsiiProxy_IMaintenanceWindowTargetRef struct {
	internal.Type__constructsIConstruct
}

