package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   apiGatewayToolFilterProperty := &ApiGatewayToolFilterProperty{
//   	FilterPath: jsii.String("filterPath"),
//   	Methods: []*string{
//   		jsii.String("methods"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytoolfilter.html
//
type CfnGatewayTargetPropsMixin_ApiGatewayToolFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytoolfilter.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytoolfilter-filterpath
	//
	FilterPath *string `field:"optional" json:"filterPath" yaml:"filterPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytoolfilter.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytoolfilter-methods
	//
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
}

