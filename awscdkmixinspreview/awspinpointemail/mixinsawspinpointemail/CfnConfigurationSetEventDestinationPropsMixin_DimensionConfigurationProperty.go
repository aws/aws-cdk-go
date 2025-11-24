package mixinsawspinpointemail


// An array of objects that define the dimensions to use when you send email events to Amazon CloudWatch.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationseteventdestination-dimensionconfiguration.html
//
type CfnConfigurationSetEventDestinationPropsMixin_DimensionConfigurationProperty struct {
	// The default value of the dimension that is published to Amazon CloudWatch if you don't provide the value of the dimension when you send an email.
	//
	// This value has to meet the following criteria:
	//
	// - It can only contain ASCII letters (a–z, A–Z), numbers (0–9), underscores (_), or dashes (-).
	// - It can contain no more than 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationseteventdestination-dimensionconfiguration.html#cfn-pinpointemail-configurationseteventdestination-dimensionconfiguration-defaultdimensionvalue
	//
	DefaultDimensionValue *string `field:"optional" json:"defaultDimensionValue" yaml:"defaultDimensionValue"`
	// The name of an Amazon CloudWatch dimension associated with an email sending metric.
	//
	// The name has to meet the following criteria:
	//
	// - It can only contain ASCII letters (a–z, A–Z), numbers (0–9), underscores (_), or dashes (-).
	// - It can contain no more than 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationseteventdestination-dimensionconfiguration.html#cfn-pinpointemail-configurationseteventdestination-dimensionconfiguration-dimensionname
	//
	DimensionName *string `field:"optional" json:"dimensionName" yaml:"dimensionName"`
	// The location where Amazon Pinpoint finds the value of a dimension to publish to Amazon CloudWatch.
	//
	// Acceptable values: `MESSAGE_TAG` , `EMAIL_HEADER` , and `LINK_TAG` .
	//
	// If you want Amazon Pinpoint to use the message tags that you specify using an `X-SES-MESSAGE-TAGS` header or a parameter to the `SendEmail` API, choose `MESSAGE_TAG` . If you want Amazon Pinpoint to use your own email headers, choose `EMAIL_HEADER` . If you want Amazon Pinpoint to use tags that are specified in your links, choose `LINK_TAG` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationseteventdestination-dimensionconfiguration.html#cfn-pinpointemail-configurationseteventdestination-dimensionconfiguration-dimensionvaluesource
	//
	DimensionValueSource *string `field:"optional" json:"dimensionValueSource" yaml:"dimensionValueSource"`
}

