package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   apiGatewayTargetConfigurationProperty := &ApiGatewayTargetConfigurationProperty{
//   	ApiGatewayToolConfiguration: &ApiGatewayToolConfigurationProperty{
//   		ToolFilters: []interface{}{
//   			&ApiGatewayToolFilterProperty{
//   				FilterPath: jsii.String("filterPath"),
//   				Methods: []*string{
//   					jsii.String("methods"),
//   				},
//   			},
//   		},
//   		ToolOverrides: []interface{}{
//   			&ApiGatewayToolOverrideProperty{
//   				Description: jsii.String("description"),
//   				Method: jsii.String("method"),
//   				Name: jsii.String("name"),
//   				Path: jsii.String("path"),
//   			},
//   		},
//   	},
//   	RestApiId: jsii.String("restApiId"),
//   	Stage: jsii.String("stage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration.html
//
type CfnGatewayTargetPropsMixin_ApiGatewayTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration-apigatewaytoolconfiguration
	//
	ApiGatewayToolConfiguration interface{} `field:"optional" json:"apiGatewayToolConfiguration" yaml:"apiGatewayToolConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration-restapiid
	//
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration-stage
	//
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

