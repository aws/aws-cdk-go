package awsappflow


// Properties for defining a `CfnConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProps := &cfnConnectorProps{
//   	connectorProvisioningConfig: &connectorProvisioningConfigProperty{
//   		lambda: &lambdaConnectorProvisioningConfigProperty{
//   			lambdaArn: jsii.String("lambdaArn"),
//   		},
//   	},
//   	connectorProvisioningType: jsii.String("connectorProvisioningType"),
//
//   	// the properties below are optional
//   	connectorLabel: jsii.String("connectorLabel"),
//   	description: jsii.String("description"),
//   }
//
type CfnConnectorProps struct {
	// `AWS::AppFlow::Connector.ConnectorProvisioningConfig`.
	ConnectorProvisioningConfig interface{} `field:"required" json:"connectorProvisioningConfig" yaml:"connectorProvisioningConfig"`
	// `AWS::AppFlow::Connector.ConnectorProvisioningType`.
	ConnectorProvisioningType *string `field:"required" json:"connectorProvisioningType" yaml:"connectorProvisioningType"`
	// `AWS::AppFlow::Connector.ConnectorLabel`.
	ConnectorLabel *string `field:"optional" json:"connectorLabel" yaml:"connectorLabel"`
	// `AWS::AppFlow::Connector.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

