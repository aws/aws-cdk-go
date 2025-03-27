package awscdkscheduleralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import scheduler_alpha "github.com/aws/aws-cdk-go/awscdkscheduleralpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   groupProps := &GroupProps{
//   	GroupName: jsii.String("groupName"),
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   }
//
// Deprecated.
type GroupProps struct {
	// The name of the schedule group.
	//
	// Up to 64 letters (uppercase and lowercase), numbers, hyphens, underscores and dots are allowed.
	// Default: - A unique name will be generated.
	//
	// Deprecated.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The removal policy for the group.
	//
	// If the group is removed also all schedules are removed.
	// Default: RemovalPolicy.RETAIN
	//
	// Deprecated.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

