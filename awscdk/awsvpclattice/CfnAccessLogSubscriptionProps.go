package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAccessLogSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessLogSubscriptionProps := &CfnAccessLogSubscriptionProps{
//   	DestinationArn: jsii.String("destinationArn"),
//
//   	// the properties below are optional
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAccessLogSubscriptionProps struct {
	// The Amazon Resource Name (ARN) of the destination.
	//
	// The supported destination types are CloudWatch Log groups, Kinesis Data Firehose delivery streams, and Amazon S3 buckets.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// The ID or Amazon Resource Name (ARN) of the service network or service.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// The tags for the access log subscription.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

