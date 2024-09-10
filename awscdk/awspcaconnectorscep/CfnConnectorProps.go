package awspcaconnectorscep


// Properties for defining a `CfnConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProps := &CfnConnectorProps{
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//
//   	// the properties below are optional
//   	MobileDeviceManagement: &MobileDeviceManagementProperty{
//   		Intune: &IntuneConfigurationProperty{
//   			AzureApplicationId: jsii.String("azureApplicationId"),
//   			Domain: jsii.String("domain"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorscep-connector.html
//
type CfnConnectorProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorscep-connector.html#cfn-pcaconnectorscep-connector-certificateauthorityarn
	//
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorscep-connector.html#cfn-pcaconnectorscep-connector-mobiledevicemanagement
	//
	MobileDeviceManagement interface{} `field:"optional" json:"mobileDeviceManagement" yaml:"mobileDeviceManagement"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorscep-connector.html#cfn-pcaconnectorscep-connector-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

