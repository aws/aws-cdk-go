package awsodb


// The AWS Security Token Service (STS) access configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stsAccessProperty := &StsAccessProperty{
//   	DomainName: jsii.String("domainName"),
//   	Ipv4Addresses: []*string{
//   		jsii.String("ipv4Addresses"),
//   	},
//   	Status: jsii.String("status"),
//   	StsPolicyDocument: jsii.String("stsPolicyDocument"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-stsaccess.html
//
type CfnOdbNetwork_StsAccessProperty struct {
	// The domain name for the AWS STS access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-stsaccess.html#cfn-odb-odbnetwork-stsaccess-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The IPv4 addresses for the AWS STS access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-stsaccess.html#cfn-odb-odbnetwork-stsaccess-ipv4addresses
	//
	Ipv4Addresses *[]*string `field:"optional" json:"ipv4Addresses" yaml:"ipv4Addresses"`
	// The status of the managed resource access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-stsaccess.html#cfn-odb-odbnetwork-stsaccess-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The endpoint policy for the AWS STS access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-stsaccess.html#cfn-odb-odbnetwork-stsaccess-stspolicydocument
	//
	StsPolicyDocument *string `field:"optional" json:"stsPolicyDocument" yaml:"stsPolicyDocument"`
}

