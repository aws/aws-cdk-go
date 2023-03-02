package awsappsync


// Properties for defining a `CfnDomainName`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainNameProps := &cfnDomainNameProps{
//   	certificateArn: jsii.String("certificateArn"),
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnDomainNameProps struct {
	// The Amazon Resource Name (ARN) of the certificate.
	//
	// This will be an AWS Certificate Manager certificate.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// The domain name.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The decription for your domain name.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

