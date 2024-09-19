package awspcaconnectorscep


// Contains configuration details for use with Microsoft Intune.
//
// For information about using Connector for SCEP for Microsoft Intune, see [Using Connector for SCEP for Microsoft Intune](https://docs.aws.amazon.com/privateca/latest/userguide/scep-connector.htmlconnector-for-scep-intune.html) .
//
// When you use Connector for SCEP for Microsoft Intune, certain functionalities are enabled by accessing Microsoft Intune through the Microsoft API. Your use of the Connector for SCEP and accompanying AWS services doesn't remove your need to have a valid license for your use of the Microsoft Intune service. You should also review the [Microsoft Intune® App Protection Policies](https://docs.aws.amazon.com/https://learn.microsoft.com/en-us/mem/intune/apps/app-protection-policy) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intuneConfigurationProperty := &IntuneConfigurationProperty{
//   	AzureApplicationId: jsii.String("azureApplicationId"),
//   	Domain: jsii.String("domain"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-intuneconfiguration.html
//
type CfnConnector_IntuneConfigurationProperty struct {
	// The directory (tenant) ID from your Microsoft Entra ID app registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-intuneconfiguration.html#cfn-pcaconnectorscep-connector-intuneconfiguration-azureapplicationid
	//
	AzureApplicationId *string `field:"required" json:"azureApplicationId" yaml:"azureApplicationId"`
	// The primary domain from your Microsoft Entra ID app registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-intuneconfiguration.html#cfn-pcaconnectorscep-connector-intuneconfiguration-domain
	//
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

