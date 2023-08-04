package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties returned by a Subscription destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   logSubscriptionDestinationConfig := &LogSubscriptionDestinationConfig{
//   	Arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	Role: role,
//   }
//
type LogSubscriptionDestinationConfig struct {
	// The ARN of the subscription's destination.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The role to assume to write log events to the destination.
	// Default: No role assumed.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

