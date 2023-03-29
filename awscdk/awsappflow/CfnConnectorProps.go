package awsappflow


// Properties for defining a `CfnConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProps := &CfnConnectorProps{
//   	ConnectorProvisioningConfig: &ConnectorProvisioningConfigProperty{
//   		Lambda: &LambdaConnectorProvisioningConfigProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   		},
//   	},
//   	ConnectorProvisioningType: jsii.String("connectorProvisioningType"),
//
//   	// the properties below are optional
//   	ConnectorLabel: jsii.String("connectorLabel"),
//   	Description: jsii.String("description"),
//   }
//
type CfnConnectorProps struct {
	// The configuration required for registering the connector.
	ConnectorProvisioningConfig interface{} `field:"required" json:"connectorProvisioningConfig" yaml:"connectorProvisioningConfig"`
	// The provisioning type used to register the connector.
	ConnectorProvisioningType *string `field:"required" json:"connectorProvisioningType" yaml:"connectorProvisioningType"`
	// The label used for registering the connector.
	ConnectorLabel *string `field:"optional" json:"connectorLabel" yaml:"connectorLabel"`
	// A description about the connector runtime setting.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

