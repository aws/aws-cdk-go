package awspcaconnectorscep


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-mobiledevicemanagement.html#cfn-pcaconnectorscep-connector-mobiledevicemanagement-intune
	//
	Intune interface{} `field:"required" json:"intune" yaml:"intune"`
}

