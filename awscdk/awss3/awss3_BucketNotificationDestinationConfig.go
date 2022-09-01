package awss3

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the properties of a notification destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var dependable iDependable
//
//   bucketNotificationDestinationConfig := &bucketNotificationDestinationConfig{
//   	arn: jsii.String("arn"),
//   	type: awscdk.Aws_s3.bucketNotificationDestinationType_LAMBDA,
//
//   	// the properties below are optional
//   	dependencies: []*iDependable{
//   		dependable,
//   	},
//   }
//
type BucketNotificationDestinationConfig struct {
	// The ARN of the destination (i.e. Lambda, SNS, SQS).
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The notification type.
	Type BucketNotificationDestinationType `field:"required" json:"type" yaml:"type"`
	// Any additional dependencies that should be resolved before the bucket notification can be configured (for example, the SNS Topic Policy resource).
	Dependencies *[]constructs.IDependable `field:"optional" json:"dependencies" yaml:"dependencies"`
}

