package awsssm


// Configuration for connecting to Azure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   azureConfigurationProperty := &AzureConfigurationProperty{
//   	ApplicationDisplayName: jsii.String("applicationDisplayName"),
//   	ApplicationId: jsii.String("applicationId"),
//   	Targets: &ConfigurationTargetsProperty{
//   		Subscriptions: []interface{}{
//   			&AzureSubscriptionProperty{
//   				DisplayName: jsii.String("displayName"),
//   				Id: jsii.String("id"),
//   			},
//   		},
//   	},
//   	TenantDisplayName: jsii.String("tenantDisplayName"),
//   	TenantId: jsii.String("tenantId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azureconfiguration.html
//
type CfnCloudConnectorPropsMixin_AzureConfigurationProperty struct {
	// The display name of the Azure AD application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azureconfiguration.html#cfn-ssm-cloudconnector-azureconfiguration-applicationdisplayname
	//
	ApplicationDisplayName *string `field:"optional" json:"applicationDisplayName" yaml:"applicationDisplayName"`
	// The Azure AD application ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azureconfiguration.html#cfn-ssm-cloudconnector-azureconfiguration-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// The targets for the cloud connector.
	//
	// If omitted, the entire tenant is targeted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azureconfiguration.html#cfn-ssm-cloudconnector-azureconfiguration-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// The display name of the Azure AD tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azureconfiguration.html#cfn-ssm-cloudconnector-azureconfiguration-tenantdisplayname
	//
	TenantDisplayName *string `field:"optional" json:"tenantDisplayName" yaml:"tenantDisplayName"`
	// The Azure AD tenant ID.
	//
	// Cannot be changed after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azureconfiguration.html#cfn-ssm-cloudconnector-azureconfiguration-tenantid
	//
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

