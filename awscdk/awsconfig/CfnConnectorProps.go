package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProps := &CfnConnectorProps{
//   	ConnectorConfiguration: &ConnectorConfigurationProperty{
//   		Azure: &AzureConnectorConfigurationProperty{
//   			ClientIdentifier: jsii.String("clientIdentifier"),
//   			TenantIdentifier: jsii.String("tenantIdentifier"),
//   		},
//   		Gcp: &GcpConnectorConfigurationProperty{
//   			ServiceAccountEmail: jsii.String("serviceAccountEmail"),
//   			WorkloadIdentityPoolProvider: jsii.String("workloadIdentityPoolProvider"),
//
//   			// the properties below are optional
//   			GcpOrganizationId: jsii.String("gcpOrganizationId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-connector.html
//
type CfnConnectorProps struct {
	// The configuration for the connector.
	//
	// Specify exactly one third-party cloud provider configuration: Azure or Gcp.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-connector.html#cfn-config-connector-connectorconfiguration
	//
	ConnectorConfiguration interface{} `field:"required" json:"connectorConfiguration" yaml:"connectorConfiguration"`
	// The tags for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-connector.html#cfn-config-connector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

