package interfacesawsapigatewayv2


// A reference to a DomainName resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainNameReference := &DomainNameReference{
//   	DomainName: jsii.String("domainName"),
//   	DomainNameArn: jsii.String("domainNameArn"),
//   }
//
type DomainNameReference struct {
	// The DomainName of the DomainName resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The ARN of the DomainName resource.
	DomainNameArn *string `field:"required" json:"domainNameArn" yaml:"domainNameArn"`
}

