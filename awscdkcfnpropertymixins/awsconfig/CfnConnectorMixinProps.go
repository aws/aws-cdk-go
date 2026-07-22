package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnConnectorMixinProps := &CfnConnectorMixinProps{
//   	ConnectorConfiguration: &ConnectorConfigurationProperty{
//   		Azure: &AzureConnectorConfigurationProperty{
//   			ClientIdentifier: jsii.String("clientIdentifier"),
//   			TenantIdentifier: jsii.String("tenantIdentifier"),
//   		},
//   	},
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
type CfnConnectorMixinProps struct {
	// The configuration for the connector.
	//
	// Specify the third-party cloud provider configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-connector.html#cfn-config-connector-connectorconfiguration
	//
	ConnectorConfiguration interface{} `field:"optional" json:"connectorConfiguration" yaml:"connectorConfiguration"`
	// The tags for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-connector.html#cfn-config-connector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

