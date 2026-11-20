package awsconfig


// The configuration for the connector.
//
// Specify exactly one third-party cloud provider configuration: Azure or Gcp.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   connectorConfigurationProperty := &ConnectorConfigurationProperty{
//   	Azure: &AzureConnectorConfigurationProperty{
//   		ClientIdentifier: jsii.String("clientIdentifier"),
//   		TenantIdentifier: jsii.String("tenantIdentifier"),
//   	},
//   	Gcp: &GcpConnectorConfigurationProperty{
//   		GcpOrganizationId: jsii.String("gcpOrganizationId"),
//   		ServiceAccountEmail: jsii.String("serviceAccountEmail"),
//   		WorkloadIdentityPoolProvider: jsii.String("workloadIdentityPoolProvider"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-connectorconfiguration.html
//
type CfnConnectorPropsMixin_ConnectorConfigurationProperty struct {
	// Configuration for connecting to Microsoft Azure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-connectorconfiguration.html#cfn-config-connector-connectorconfiguration-azure
	//
	Azure interface{} `field:"optional" json:"azure" yaml:"azure"`
	// Configuration for connecting to Google Cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-connectorconfiguration.html#cfn-config-connector-connectorconfiguration-gcp
	//
	Gcp interface{} `field:"optional" json:"gcp" yaml:"gcp"`
}

