package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProps := &cfnCertificateProps{
//   	certificateName: jsii.String("certificateName"),
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	subjectAlternativeNames: []*string{
//   		jsii.String("subjectAlternativeNames"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCertificateProps struct {
	// The name of the certificate.
	CertificateName *string `field:"required" json:"certificateName" yaml:"certificateName"`
	// The domain name of the certificate.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// An array of strings that specify the alternate domains (such as `example.org` ) and subdomains (such as `blog.example.com` ) of the certificate.
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

