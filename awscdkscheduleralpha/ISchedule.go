package awscdkscheduleralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2/internal"
)

// Interface representing a created or an imported `Schedule`.
// Experimental.
type ISchedule interface {
	awscdk.IResource
}

// The jsii proxy for ISchedule
type jsiiProxy_ISchedule struct {
	internal.Type__awscdkIResource
}

