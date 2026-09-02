package awsssm


// The configuration for the cloud connector.
//
// Currently supports Azure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudConnectorConfigurationProperty := &CloudConnectorConfigurationProperty{
//   	AzureConfiguration: &AzureConfigurationProperty{
//   		ApplicationId: jsii.String("applicationId"),
//   		TenantId: jsii.String("tenantId"),
//
//   		// the properties below are optional
//   		ApplicationDisplayName: jsii.String("applicationDisplayName"),
//   		Targets: &ConfigurationTargetsProperty{
//   			Subscriptions: []interface{}{
//   				&AzureSubscriptionProperty{
//   					Id: jsii.String("id"),
//
//   					// the properties below are optional
//   					DisplayName: jsii.String("displayName"),
//   				},
//   			},
//   		},
//   		TenantDisplayName: jsii.String("tenantDisplayName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-cloudconnectorconfiguration.html
//
type CfnCloudConnector_CloudConnectorConfigurationProperty struct {
	// Configuration for connecting to Azure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-cloudconnectorconfiguration.html#cfn-ssm-cloudconnector-cloudconnectorconfiguration-azureconfiguration
	//
	AzureConfiguration interface{} `field:"required" json:"azureConfiguration" yaml:"azureConfiguration"`
}

