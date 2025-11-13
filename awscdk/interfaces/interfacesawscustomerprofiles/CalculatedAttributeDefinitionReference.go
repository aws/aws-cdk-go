package interfacesawscustomerprofiles


// A reference to a CalculatedAttributeDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   calculatedAttributeDefinitionReference := &CalculatedAttributeDefinitionReference{
//   	CalculatedAttributeName: jsii.String("calculatedAttributeName"),
//   	DomainName: jsii.String("domainName"),
//   }
//
type CalculatedAttributeDefinitionReference struct {
	// The CalculatedAttributeName of the CalculatedAttributeDefinition resource.
	CalculatedAttributeName *string `field:"required" json:"calculatedAttributeName" yaml:"calculatedAttributeName"`
	// The DomainName of the CalculatedAttributeDefinition resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

