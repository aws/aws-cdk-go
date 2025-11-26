package previewawsappflowmixins


// Properties for CfnConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectorMixinProps := &CfnConnectorMixinProps{
//   	ConnectorLabel: jsii.String("connectorLabel"),
//   	ConnectorProvisioningConfig: &ConnectorProvisioningConfigProperty{
//   		Lambda: &LambdaConnectorProvisioningConfigProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   		},
//   	},
//   	ConnectorProvisioningType: jsii.String("connectorProvisioningType"),
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connector.html
//
type CfnConnectorMixinProps struct {
	// The label used for registering the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connector.html#cfn-appflow-connector-connectorlabel
	//
	ConnectorLabel *string `field:"optional" json:"connectorLabel" yaml:"connectorLabel"`
	// The configuration required for registering the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connector.html#cfn-appflow-connector-connectorprovisioningconfig
	//
	ConnectorProvisioningConfig interface{} `field:"optional" json:"connectorProvisioningConfig" yaml:"connectorProvisioningConfig"`
	// The provisioning type used to register the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connector.html#cfn-appflow-connector-connectorprovisioningtype
	//
	ConnectorProvisioningType *string `field:"optional" json:"connectorProvisioningType" yaml:"connectorProvisioningType"`
	// A description about the connector runtime setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connector.html#cfn-appflow-connector-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

