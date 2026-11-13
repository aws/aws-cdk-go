package awscloudwatch


// Where the notification is sent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetProperty := &TargetProperty{
//   	Arn: jsii.String("arn"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Metadata: map[string]*string{
//   		"metadataKey": jsii.String("metadata"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-target.html
//
type CfnAlert_TargetProperty struct {
	// The ARN of the target.
	//
	// For sns, the ARN of the Amazon SNS topic. For slack and pagerduty, the ARN of a CloudWatch Omni integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-target.html#cfn-cloudwatch-alert-target-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The target type.
	//
	// Valid values: sns, slack, pagerduty.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-target.html#cfn-cloudwatch-alert-target-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Settings specific to the target, as key-value pairs (up to 20).
	//
	// The keys depend on the integration type. Not used for sns targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-target.html#cfn-cloudwatch-alert-target-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
}

