package awscloudwatch


// Where the notification is sent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   targetProperty := &TargetProperty{
//   	Arn: jsii.String("arn"),
//   	Metadata: map[string]*string{
//   		"metadataKey": jsii.String("metadata"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-target.html
//
type CfnAlertPropsMixin_TargetProperty struct {
	// The ARN of the target.
	//
	// For sns, the ARN of the Amazon SNS topic. For slack and pagerduty, the ARN of a CloudWatch Omni integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-target.html#cfn-cloudwatch-alert-target-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Settings specific to the target, as key-value pairs (up to 20).
	//
	// The keys depend on the integration type. Not used for sns targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-target.html#cfn-cloudwatch-alert-target-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The target type.
	//
	// Valid values: sns, slack, pagerduty.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-target.html#cfn-cloudwatch-alert-target-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

