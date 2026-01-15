package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   apiGatewayToolConfigurationProperty := &ApiGatewayToolConfigurationProperty{
//   	ToolFilters: []interface{}{
//   		&ApiGatewayToolFilterProperty{
//   			FilterPath: jsii.String("filterPath"),
//   			Methods: []*string{
//   				jsii.String("methods"),
//   			},
//   		},
//   	},
//   	ToolOverrides: []interface{}{
//   		&ApiGatewayToolOverrideProperty{
//   			Description: jsii.String("description"),
//   			Method: jsii.String("method"),
//   			Name: jsii.String("name"),
//   			Path: jsii.String("path"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytoolconfiguration.html
//
type CfnGatewayTargetPropsMixin_ApiGatewayToolConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytoolconfiguration.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytoolconfiguration-toolfilters
	//
	ToolFilters interface{} `field:"optional" json:"toolFilters" yaml:"toolFilters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytoolconfiguration.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytoolconfiguration-tooloverrides
	//
	ToolOverrides interface{} `field:"optional" json:"toolOverrides" yaml:"toolOverrides"`
}

