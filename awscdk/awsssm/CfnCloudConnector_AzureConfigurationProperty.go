package awsssm


// Configuration for connecting to Azure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   azureConfigurationProperty := &AzureConfigurationProperty{
//   	ApplicationId: jsii.String("applicationId"),
//   	TenantId: jsii.String("tenantId"),
//
//   	// the properties below are optional
//   	ApplicationDisplayName: jsii.String("applicationDisplayName"),
//   	Targets: &ConfigurationTargetsProperty{
//   		Subscriptions: []interface{}{
//   			&AzureSubscriptionProperty{
//   				Id: jsii.String("id"),
//
//   				// the properties below are optional
//   				DisplayName: jsii.String("displayName"),
//   			},
//   		},
//   	},
//   	TenantDisplayName: jsii.String("tenantDisplayName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azureconfiguration.html
//
type CfnCloudConnector_AzureConfigurationProperty struct {
	// The Azure AD application ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azureconfiguration.html#cfn-ssm-cloudconnector-azureconfiguration-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The Azure AD tenant ID.
	//
	// Cannot be changed after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azureconfiguration.html#cfn-ssm-cloudconnector-azureconfiguration-tenantid
	//
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// The display name of the Azure AD application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azureconfiguration.html#cfn-ssm-cloudconnector-azureconfiguration-applicationdisplayname
	//
	ApplicationDisplayName *string `field:"optional" json:"applicationDisplayName" yaml:"applicationDisplayName"`
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
}

