package previewawspcaconnectoradmixins


// Properties for CfnConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectorMixinProps := &CfnConnectorMixinProps{
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	DirectoryId: jsii.String("directoryId"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	VpcInformation: &VpcInformationProperty{
//   		IpAddressType: jsii.String("ipAddressType"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-connector.html
//
type CfnConnectorMixinProps struct {
	// The Amazon Resource Name (ARN) of the certificate authority being used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-connector.html#cfn-pcaconnectorad-connector-certificateauthorityarn
	//
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The identifier of the Active Directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-connector.html#cfn-pcaconnectorad-connector-directoryid
	//
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Metadata assigned to a connector consisting of a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-connector.html#cfn-pcaconnectorad-connector-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Information of the VPC and security group(s) used with the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-connector.html#cfn-pcaconnectorad-connector-vpcinformation
	//
	VpcInformation interface{} `field:"optional" json:"vpcInformation" yaml:"vpcInformation"`
}

