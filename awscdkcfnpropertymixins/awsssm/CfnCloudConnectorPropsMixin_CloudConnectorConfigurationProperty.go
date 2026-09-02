package awsssm


// The configuration for the cloud connector.
//
// Currently supports Azure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cloudConnectorConfigurationProperty := &CloudConnectorConfigurationProperty{
//   	AzureConfiguration: &AzureConfigurationProperty{
//   		ApplicationDisplayName: jsii.String("applicationDisplayName"),
//   		ApplicationId: jsii.String("applicationId"),
//   		Targets: &ConfigurationTargetsProperty{
//   			Subscriptions: []interface{}{
//   				&AzureSubscriptionProperty{
//   					DisplayName: jsii.String("displayName"),
//   					Id: jsii.String("id"),
//   				},
//   			},
//   		},
//   		TenantDisplayName: jsii.String("tenantDisplayName"),
//   		TenantId: jsii.String("tenantId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-cloudconnectorconfiguration.html
//
type CfnCloudConnectorPropsMixin_CloudConnectorConfigurationProperty struct {
	// Configuration for connecting to Azure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-cloudconnectorconfiguration.html#cfn-ssm-cloudconnector-cloudconnectorconfiguration-azureconfiguration
	//
	AzureConfiguration interface{} `field:"optional" json:"azureConfiguration" yaml:"azureConfiguration"`
}

