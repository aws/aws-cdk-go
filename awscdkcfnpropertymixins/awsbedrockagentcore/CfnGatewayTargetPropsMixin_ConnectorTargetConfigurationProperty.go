package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var parameterValues interface{}
//
//   connectorTargetConfigurationProperty := &ConnectorTargetConfigurationProperty{
//   	Configurations: []interface{}{
//   		&ConnectorConfigurationProperty{
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   			ParameterOverrides: []interface{}{
//   				&ConnectorParameterOverrideProperty{
//   					Description: jsii.String("description"),
//   					Path: jsii.String("path"),
//   					Visible: jsii.Boolean(false),
//   				},
//   			},
//   			ParameterValues: parameterValues,
//   		},
//   	},
//   	Enabled: []*string{
//   		jsii.String("enabled"),
//   	},
//   	Source: &ConnectorSourceProperty{
//   		ConnectorId: jsii.String("connectorId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectortargetconfiguration.html
//
type CfnGatewayTargetPropsMixin_ConnectorTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectortargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-connectortargetconfiguration-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectortargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-connectortargetconfiguration-enabled
	//
	Enabled *[]*string `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectortargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-connectortargetconfiguration-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

