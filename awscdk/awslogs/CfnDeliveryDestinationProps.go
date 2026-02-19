package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDeliveryDestination`.
//
// Example:
//   import _ "github.com/aws-samples/dummy/awscdkmixinspreview/with"
//   import cloudfrontMixins "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   // Create CloudFront distribution
//   var bucket Bucket
//
//   distribution := cloudfront.NewDistribution(scope, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.S3BucketOrigin_WithOriginAccessControl(bucket),
//   	},
//   })
//
//   // Create destination bucket
//   destBucket := s3.NewBucket(scope, jsii.String("DeliveryBucket"))
//   // Add permissions to bucket to facilitate log delivery
//   bucketPolicy := s3.NewBucketPolicy(scope, jsii.String("DeliveryBucketPolicy"), &BucketPolicyProps{
//   	Bucket: destBucket,
//   	Document: iam.NewPolicyDocument(),
//   })
//   // Create S3 delivery destination for logs
//   destination := logs.NewCfnDeliveryDestination(scope, jsii.String("Destination"), &CfnDeliveryDestinationProps{
//   	DestinationResourceArn: destBucket.bucketArn,
//   	Name: jsii.String("unique-destination-name"),
//   	DeliveryDestinationType: jsii.String("S3"),
//   })
//
//   distribution.With(cloudfrontMixins.CfnDistributionLogsMixin_CONNECTION_LOGS().ToDestination(destination))
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
	// Displays whether this delivery destination is CloudWatch Logs, Amazon S3, Firehose, or X-Ray.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverydestination.html#cfn-logs-deliverydestination-deliverydestinationtype
	//
	DeliveryDestinationType *string `field:"optional" json:"deliveryDestinationType" yaml:"deliveryDestinationType"`
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

