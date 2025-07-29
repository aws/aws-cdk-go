package awspcaconnectorad


// Properties for defining a `CfnConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProps := &CfnConnectorProps{
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	DirectoryId: jsii.String("directoryId"),
//   	VpcInformation: &VpcInformationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//
//   		// the properties below are optional
//   		IpAddressType: jsii.String("ipAddressType"),
//   	},
//
//   	// the properties below are optional
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-connector.html
//
type CfnConnectorProps struct {
	// The Amazon Resource Name (ARN) of the certificate authority being used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-connector.html#cfn-pcaconnectorad-connector-certificateauthorityarn
	//
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The identifier of the Active Directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-connector.html#cfn-pcaconnectorad-connector-directoryid
	//
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// Information of the VPC and security group(s) used with the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-connector.html#cfn-pcaconnectorad-connector-vpcinformation
	//
	VpcInformation interface{} `field:"required" json:"vpcInformation" yaml:"vpcInformation"`
	// Metadata assigned to a connector consisting of a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-connector.html#cfn-pcaconnectorad-connector-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

