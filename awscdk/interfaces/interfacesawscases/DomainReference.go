package interfacesawscases


// A reference to a Domain resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainReference := &DomainReference{
//   	DomainArn: jsii.String("domainArn"),
//   }
//
type DomainReference struct {
	// The DomainArn of the Domain resource.
	DomainArn *string `field:"required" json:"domainArn" yaml:"domainArn"`
}

