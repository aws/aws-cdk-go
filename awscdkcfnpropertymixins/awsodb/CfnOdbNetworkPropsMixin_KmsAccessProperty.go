package awsodb


// The AWS Key Management Service (KMS) access configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   kmsAccessProperty := &KmsAccessProperty{
//   	DomainName: jsii.String("domainName"),
//   	Ipv4Addresses: []*string{
//   		jsii.String("ipv4Addresses"),
//   	},
//   	KmsPolicyDocument: jsii.String("kmsPolicyDocument"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-kmsaccess.html
//
type CfnOdbNetworkPropsMixin_KmsAccessProperty struct {
	// The domain name for the AWS KMS access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-kmsaccess.html#cfn-odb-odbnetwork-kmsaccess-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The IPv4 addresses for the AWS KMS access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-kmsaccess.html#cfn-odb-odbnetwork-kmsaccess-ipv4addresses
	//
	Ipv4Addresses *[]*string `field:"optional" json:"ipv4Addresses" yaml:"ipv4Addresses"`
	// The endpoint policy for the AWS KMS access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-kmsaccess.html#cfn-odb-odbnetwork-kmsaccess-kmspolicydocument
	//
	KmsPolicyDocument *string `field:"optional" json:"kmsPolicyDocument" yaml:"kmsPolicyDocument"`
	// The status of the managed resource access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-kmsaccess.html#cfn-odb-odbnetwork-kmsaccess-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

