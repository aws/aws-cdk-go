package awsconfig


// The configuration for the connector.
//
// Specify the third-party cloud provider configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorConfigurationProperty := &ConnectorConfigurationProperty{
//   	Azure: &AzureConnectorConfigurationProperty{
//   		ClientIdentifier: jsii.String("clientIdentifier"),
//   		TenantIdentifier: jsii.String("tenantIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-connectorconfiguration.html
//
type CfnConnector_ConnectorConfigurationProperty struct {
	// Configuration for connecting to Microsoft Azure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-connectorconfiguration.html#cfn-config-connector-connectorconfiguration-azure
	//
	Azure interface{} `field:"optional" json:"azure" yaml:"azure"`
}

