package awsconfig


// Configuration for connecting to Microsoft Azure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   azureConnectorConfigurationProperty := &AzureConnectorConfigurationProperty{
//   	ClientIdentifier: jsii.String("clientIdentifier"),
//   	TenantIdentifier: jsii.String("tenantIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-azureconnectorconfiguration.html
//
type CfnConnector_AzureConnectorConfigurationProperty struct {
	// The Azure client (application) identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-azureconnectorconfiguration.html#cfn-config-connector-azureconnectorconfiguration-clientidentifier
	//
	ClientIdentifier *string `field:"required" json:"clientIdentifier" yaml:"clientIdentifier"`
	// The Azure tenant identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-azureconnectorconfiguration.html#cfn-config-connector-azureconnectorconfiguration-tenantidentifier
	//
	TenantIdentifier *string `field:"required" json:"tenantIdentifier" yaml:"tenantIdentifier"`
}

