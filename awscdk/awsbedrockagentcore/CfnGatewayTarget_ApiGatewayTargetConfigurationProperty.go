package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//
//   		// the properties below are optional
//   		ToolOverrides: []interface{}{
//   			&ApiGatewayToolOverrideProperty{
//   				Method: jsii.String("method"),
//   				Name: jsii.String("name"),
//   				Path: jsii.String("path"),
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   			},
//   		},
//   	},
//   	RestApiId: jsii.String("restApiId"),
//   	Stage: jsii.String("stage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration.html
//
type CfnGatewayTarget_ApiGatewayTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration-apigatewaytoolconfiguration
	//
	ApiGatewayToolConfiguration interface{} `field:"required" json:"apiGatewayToolConfiguration" yaml:"apiGatewayToolConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration-restapiid
	//
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration-stage
	//
	Stage *string `field:"required" json:"stage" yaml:"stage"`
}

