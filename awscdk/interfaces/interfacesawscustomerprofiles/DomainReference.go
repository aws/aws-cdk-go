package interfacesawscustomerprofiles


// A reference to a Domain resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainReference := &DomainReference{
//   	DomainName: jsii.String("domainName"),
//   }
//
type DomainReference struct {
	// The DomainName of the Domain resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

