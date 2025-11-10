package awsrtbfabric


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responderErrorMaskingForHttpCodeProperty := &ResponderErrorMaskingForHttpCodeProperty{
//   	Action: jsii.String("action"),
//   	HttpCode: jsii.String("httpCode"),
//   	LoggingTypes: []*string{
//   		jsii.String("loggingTypes"),
//   	},
//
//   	// the properties below are optional
//   	ResponseLoggingPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-respondererrormaskingforhttpcode.html
//
type CfnLink_ResponderErrorMaskingForHttpCodeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-respondererrormaskingforhttpcode.html#cfn-rtbfabric-link-respondererrormaskingforhttpcode-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-respondererrormaskingforhttpcode.html#cfn-rtbfabric-link-respondererrormaskingforhttpcode-httpcode
	//
	HttpCode *string `field:"required" json:"httpCode" yaml:"httpCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-respondererrormaskingforhttpcode.html#cfn-rtbfabric-link-respondererrormaskingforhttpcode-loggingtypes
	//
	LoggingTypes *[]*string `field:"required" json:"loggingTypes" yaml:"loggingTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-respondererrormaskingforhttpcode.html#cfn-rtbfabric-link-respondererrormaskingforhttpcode-responseloggingpercentage
	//
	ResponseLoggingPercentage *float64 `field:"optional" json:"responseLoggingPercentage" yaml:"responseLoggingPercentage"`
}

