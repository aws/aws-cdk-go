package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Represents the properties of a notification destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dependable iDependable
//
//   bucketNotificationDestinationConfig := &BucketNotificationDestinationConfig{
//   	Arn: jsii.String("arn"),
//   	Type: awscdk.Aws_s3.BucketNotificationDestinationType_LAMBDA,
//
//   	// the properties below are optional
//   	Dependencies: []*iDependable{
//   		dependable,
//   	},
//   }
//
// Experimental.
type BucketNotificationDestinationConfig struct {
	// The ARN of the destination (i.e. Lambda, SNS, SQS).
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The notification type.
	// Experimental.
	Type BucketNotificationDestinationType `field:"required" json:"type" yaml:"type"`
	// Any additional dependencies that should be resolved before the bucket notification can be configured (for example, the SNS Topic Policy resource).
	// Experimental.
	Dependencies *[]awscdk.IDependable `field:"optional" json:"dependencies" yaml:"dependencies"`
}

