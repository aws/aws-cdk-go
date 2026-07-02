package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterValues interface{}
//
//   connectorTargetConfigurationProperty := &ConnectorTargetConfigurationProperty{
//   	Source: &ConnectorSourceProperty{
//   		ConnectorId: jsii.String("connectorId"),
//   	},
//
//   	// the properties below are optional
//   	Configurations: []interface{}{
//   		&ConnectorConfigurationProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			ParameterOverrides: []interface{}{
//   				&ConnectorParameterOverrideProperty{
//   					Path: jsii.String("path"),
//
//   					// the properties below are optional
//   					Description: jsii.String("description"),
//   					Visible: jsii.Boolean(false),
//   				},
//   			},
//   			ParameterValues: parameterValues,
//   		},
//   	},
//   	Enabled: []*string{
//   		jsii.String("enabled"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectortargetconfiguration.html
//
type CfnGatewayTarget_ConnectorTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectortargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-connectortargetconfiguration-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectortargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-connectortargetconfiguration-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectortargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-connectortargetconfiguration-enabled
	//
	Enabled *[]*string `field:"optional" json:"enabled" yaml:"enabled"`
}

