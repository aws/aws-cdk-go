package awsdatazone


// A reference to a DomainUnit resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainUnitReference := &DomainUnitReference{
//   	DomainId: jsii.String("domainId"),
//   	DomainUnitId: jsii.String("domainUnitId"),
//   }
//
type DomainUnitReference struct {
	// The DomainId of the DomainUnit resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The Id of the DomainUnit resource.
	DomainUnitId *string `field:"required" json:"domainUnitId" yaml:"domainUnitId"`
}

