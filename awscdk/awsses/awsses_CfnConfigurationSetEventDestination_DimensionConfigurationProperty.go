package awsses


// Contains the dimension configuration to use when you publish email sending events to Amazon CloudWatch.
//
// For information about publishing email sending events to Amazon CloudWatch, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dimensionConfigurationProperty := &dimensionConfigurationProperty{
//   	defaultDimensionValue: jsii.String("defaultDimensionValue"),
//   	dimensionName: jsii.String("dimensionName"),
//   	dimensionValueSource: jsii.String("dimensionValueSource"),
//   }
//
type CfnConfigurationSetEventDestination_DimensionConfigurationProperty struct {
	// The default value of the dimension that is published to Amazon CloudWatch if you do not provide the value of the dimension when you send an email.
	//
	// The default value must meet the following requirements:
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), dashes (-), at signs (@), or periods (.).
	// - Contain 256 characters or fewer.
	DefaultDimensionValue *string `field:"required" json:"defaultDimensionValue" yaml:"defaultDimensionValue"`
	// The name of an Amazon CloudWatch dimension associated with an email sending metric.
	//
	// The name must meet the following requirements:
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), dashes (-), or colons (:).
	// - Contain 256 characters or fewer.
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
	// The place where Amazon SES finds the value of a dimension to publish to Amazon CloudWatch.
	//
	// To use the message tags that you specify using an `X-SES-MESSAGE-TAGS` header or a parameter to the `SendEmail` / `SendRawEmail` API, specify `messageTag` . To use your own email headers, specify `emailHeader` . To put a custom tag on any link included in your email, specify `linkTag` .
	DimensionValueSource *string `field:"required" json:"dimensionValueSource" yaml:"dimensionValueSource"`
}

