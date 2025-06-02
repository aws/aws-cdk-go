package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDeliveryDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var deliveryDestinationPolicy interface{}
//
//   cfnDeliveryDestinationProps := &CfnDeliveryDestinationProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DeliveryDestinationPolicy: &DestinationPolicyProperty{
//   		DeliveryDestinationName: jsii.String("deliveryDestinationName"),
//   		DeliveryDestinationPolicy: deliveryDestinationPolicy,
//   	},
//   	DestinationResourceArn: jsii.String("destinationResourceArn"),
//   	OutputFormat: jsii.String("outputFormat"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverydestination.html
//
type CfnDeliveryDestinationProps struct {
	// The name of this delivery destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverydestination.html#cfn-logs-deliverydestination-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.
	//
	// For examples of this policy, see [Examples](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestinationPolicy.html#API_PutDeliveryDestinationPolicy_Examples) in the CloudWatch Logs API Reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverydestination.html#cfn-logs-deliverydestination-deliverydestinationpolicy
	//
	DeliveryDestinationPolicy interface{} `field:"optional" json:"deliveryDestinationPolicy" yaml:"deliveryDestinationPolicy"`
	// The ARN of the AWS destination that this delivery destination represents.
	//
	// That AWS destination can be a log group in CloudWatch Logs , an Amazon S3 bucket, or a Firehose stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverydestination.html#cfn-logs-deliverydestination-destinationresourcearn
	//
	DestinationResourceArn *string `field:"optional" json:"destinationResourceArn" yaml:"destinationResourceArn"`
	// The format of the logs that are sent to this delivery destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverydestination.html#cfn-logs-deliverydestination-outputformat
	//
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// An array of key-value pairs to apply to the delivery destination.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverydestination.html#cfn-logs-deliverydestination-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

