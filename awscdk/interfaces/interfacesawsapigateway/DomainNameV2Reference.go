package interfacesawsapigateway


// A reference to a DomainNameV2 resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainNameV2Reference := &DomainNameV2Reference{
//   	DomainNameArn: jsii.String("domainNameArn"),
//   }
//
type DomainNameV2Reference struct {
	// The DomainNameArn of the DomainNameV2 resource.
	DomainNameArn *string `field:"required" json:"domainNameArn" yaml:"domainNameArn"`
}

