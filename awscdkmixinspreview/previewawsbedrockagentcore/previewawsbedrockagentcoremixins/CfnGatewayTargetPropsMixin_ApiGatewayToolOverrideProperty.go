package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   apiGatewayToolOverrideProperty := &ApiGatewayToolOverrideProperty{
//   	Description: jsii.String("description"),
//   	Method: jsii.String("method"),
//   	Name: jsii.String("name"),
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytooloverride.html
//
type CfnGatewayTargetPropsMixin_ApiGatewayToolOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytooloverride.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytooloverride-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytooloverride.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytooloverride-method
	//
	Method *string `field:"optional" json:"method" yaml:"method"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytooloverride.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytooloverride-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytooloverride.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytooloverride-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

