package mixinsawspcaconnectorscep


// Properties for CfnConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectorMixinProps := &CfnConnectorMixinProps{
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
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
type CfnConnectorMixinProps struct {
	// The Amazon Resource Name (ARN) of the certificate authority associated with the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorscep-connector.html#cfn-pcaconnectorscep-connector-certificateauthorityarn
	//
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Contains settings relevant to the mobile device management system that you chose for the connector.
	//
	// If you didn't configure `MobileDeviceManagement` , then the connector is for general-purpose use and this object is empty.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorscep-connector.html#cfn-pcaconnectorscep-connector-mobiledevicemanagement
	//
	MobileDeviceManagement interface{} `field:"optional" json:"mobileDeviceManagement" yaml:"mobileDeviceManagement"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorscep-connector.html#cfn-pcaconnectorscep-connector-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

