package awslightsail


// `PublicDomainName` is a property of the [AWS::Lightsail::Container](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html) resource. It describes the public domain names to use with a container service, such as `example.com` and `www.example.com` . It also describes the certificates to use with a container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicDomainNameProperty := &publicDomainNameProperty{
//   	certificateName: jsii.String("certificateName"),
//   	domainNames: []*string{
//   		jsii.String("domainNames"),
//   	},
//   }
//
type CfnContainer_PublicDomainNameProperty struct {
	// The name of the certificate for the public domains.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// The public domain names to use with the container service.
	DomainNames *[]*string `field:"optional" json:"domainNames" yaml:"domainNames"`
}

