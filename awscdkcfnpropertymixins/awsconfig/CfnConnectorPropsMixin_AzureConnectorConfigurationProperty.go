package awsconfig


// Configuration for connecting to Microsoft Azure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   azureConnectorConfigurationProperty := &AzureConnectorConfigurationProperty{
//   	ClientIdentifier: jsii.String("clientIdentifier"),
//   	TenantIdentifier: jsii.String("tenantIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-azureconnectorconfiguration.html
//
type CfnConnectorPropsMixin_AzureConnectorConfigurationProperty struct {
	// The Azure client (application) identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-azureconnectorconfiguration.html#cfn-config-connector-azureconnectorconfiguration-clientidentifier
	//
	ClientIdentifier *string `field:"optional" json:"clientIdentifier" yaml:"clientIdentifier"`
	// The Azure tenant identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-azureconnectorconfiguration.html#cfn-config-connector-azureconnectorconfiguration-tenantidentifier
	//
	TenantIdentifier *string `field:"optional" json:"tenantIdentifier" yaml:"tenantIdentifier"`
}

