package awspcaconnectorscep


// If you don't supply a value, by default Connector for SCEP creates a connector for general-purpose use.
//
// A general-purpose connector is designed to work with clients or endpoints that support the SCEP protocol, except Connector for SCEP for Microsoft Intune. For information about considerations and limitations with using Connector for SCEP, see [Considerations and Limitations](https://docs.aws.amazon.com/privateca/latest/userguide/scep-connector.htmlc4scep-considerations-limitations.html) .
//
// If you provide an `IntuneConfiguration` , Connector for SCEP creates a connector for use with Microsoft Intune, and you manage the challenge passwords using Microsoft Intune. For more information, see [Using Connector for SCEP for Microsoft Intune](https://docs.aws.amazon.com/privateca/latest/userguide/scep-connector.htmlconnector-for-scep-intune.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mobileDeviceManagementProperty := &MobileDeviceManagementProperty{
//   	Intune: &IntuneConfigurationProperty{
//   		AzureApplicationId: jsii.String("azureApplicationId"),
//   		Domain: jsii.String("domain"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-mobiledevicemanagement.html
//
type CfnConnector_MobileDeviceManagementProperty struct {
	// Configuration settings for use with Microsoft Intune.
	//
	// For information about using Connector for SCEP for Microsoft Intune, see [Using Connector for SCEP for Microsoft Intune](https://docs.aws.amazon.com/privateca/latest/userguide/scep-connector.htmlconnector-for-scep-intune.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-mobiledevicemanagement.html#cfn-pcaconnectorscep-connector-mobiledevicemanagement-intune
	//
	Intune interface{} `field:"required" json:"intune" yaml:"intune"`
}

