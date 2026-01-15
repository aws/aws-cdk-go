package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiGatewayToolOverrideProperty := &ApiGatewayToolOverrideProperty{
//   	Method: jsii.String("method"),
//   	Name: jsii.String("name"),
//   	Path: jsii.String("path"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytooloverride.html
//
type CfnGatewayTarget_ApiGatewayToolOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytooloverride.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytooloverride-method
	//
	Method *string `field:"required" json:"method" yaml:"method"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytooloverride.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytooloverride-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytooloverride.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytooloverride-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytooloverride.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytooloverride-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

