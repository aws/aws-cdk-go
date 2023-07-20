package awsappsync


// Properties for defining a `CfnDomainName`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainNameProps := &CfnDomainNameProps{
//   	CertificateArn: jsii.String("certificateArn"),
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainname.html
//
type CfnDomainNameProps struct {
	// The Amazon Resource Name (ARN) of the certificate.
	//
	// This will be an AWS Certificate Manager certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainname.html#cfn-appsync-domainname-certificatearn
	//
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// The domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainname.html#cfn-appsync-domainname-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The decription for your domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainname.html#cfn-appsync-domainname-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

