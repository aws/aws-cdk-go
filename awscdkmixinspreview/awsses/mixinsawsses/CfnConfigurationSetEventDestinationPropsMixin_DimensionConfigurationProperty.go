package mixinsawsses


// An object that defines the dimension configuration to use when you send email events to Amazon CloudWatch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dimensionConfigurationProperty := &DimensionConfigurationProperty{
//   	DefaultDimensionValue: jsii.String("defaultDimensionValue"),
//   	DimensionName: jsii.String("dimensionName"),
//   	DimensionValueSource: jsii.String("dimensionValueSource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-dimensionconfiguration.html
//
type CfnConfigurationSetEventDestinationPropsMixin_DimensionConfigurationProperty struct {
	// The default value of the dimension that is published to Amazon CloudWatch if you don't provide the value of the dimension when you send an email.
	//
	// This value has to meet the following criteria:
	//
	// - Can only contain ASCII letters (a–z, A–Z), numbers (0–9), underscores (_), or dashes (-), at signs (@), and periods (.).
	// - It can contain no more than 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-dimensionconfiguration.html#cfn-ses-configurationseteventdestination-dimensionconfiguration-defaultdimensionvalue
	//
	DefaultDimensionValue *string `field:"optional" json:"defaultDimensionValue" yaml:"defaultDimensionValue"`
	// The name of an Amazon CloudWatch dimension associated with an email sending metric.
	//
	// The name has to meet the following criteria:
	//
	// - It can only contain ASCII letters (a–z, A–Z), numbers (0–9), underscores (_), or dashes (-).
	// - It can contain no more than 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-dimensionconfiguration.html#cfn-ses-configurationseteventdestination-dimensionconfiguration-dimensionname
	//
	DimensionName *string `field:"optional" json:"dimensionName" yaml:"dimensionName"`
	// The location where the Amazon SES API v2 finds the value of a dimension to publish to Amazon CloudWatch.
	//
	// To use the message tags that you specify using an `X-SES-MESSAGE-TAGS` header or a parameter to the `SendEmail` or `SendRawEmail` API, choose `messageTag` . To use your own email headers, choose `emailHeader` . To use link tags, choose `linkTag` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-dimensionconfiguration.html#cfn-ses-configurationseteventdestination-dimensionconfiguration-dimensionvaluesource
	//
	DimensionValueSource *string `field:"optional" json:"dimensionValueSource" yaml:"dimensionValueSource"`
}

