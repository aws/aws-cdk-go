package awsvoiceid


// A reference to a Domain resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainReference := &DomainReference{
//   	DomainId: jsii.String("domainId"),
//   }
//
type DomainReference struct {
	// The DomainId of the Domain resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
}

