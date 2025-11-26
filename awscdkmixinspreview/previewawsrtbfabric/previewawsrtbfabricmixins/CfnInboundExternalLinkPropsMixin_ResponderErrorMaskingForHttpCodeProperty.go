package previewawsrtbfabricmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responderErrorMaskingForHttpCodeProperty := &ResponderErrorMaskingForHttpCodeProperty{
//   	Action: jsii.String("action"),
//   	HttpCode: jsii.String("httpCode"),
//   	LoggingTypes: []*string{
//   		jsii.String("loggingTypes"),
//   	},
//   	ResponseLoggingPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-respondererrormaskingforhttpcode.html
//
type CfnInboundExternalLinkPropsMixin_ResponderErrorMaskingForHttpCodeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-respondererrormaskingforhttpcode.html#cfn-rtbfabric-inboundexternallink-respondererrormaskingforhttpcode-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-respondererrormaskingforhttpcode.html#cfn-rtbfabric-inboundexternallink-respondererrormaskingforhttpcode-httpcode
	//
	HttpCode *string `field:"optional" json:"httpCode" yaml:"httpCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-respondererrormaskingforhttpcode.html#cfn-rtbfabric-inboundexternallink-respondererrormaskingforhttpcode-loggingtypes
	//
	LoggingTypes *[]*string `field:"optional" json:"loggingTypes" yaml:"loggingTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-respondererrormaskingforhttpcode.html#cfn-rtbfabric-inboundexternallink-respondererrormaskingforhttpcode-responseloggingpercentage
	//
	ResponseLoggingPercentage *float64 `field:"optional" json:"responseLoggingPercentage" yaml:"responseLoggingPercentage"`
}

