package interfacesawscustomerprofiles


// A reference to a DomainObjectType resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainObjectTypeReference := &DomainObjectTypeReference{
//   	DomainName: jsii.String("domainName"),
//   	ObjectTypeName: jsii.String("objectTypeName"),
//   }
//
type DomainObjectTypeReference struct {
	// The DomainName of the DomainObjectType resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The ObjectTypeName of the DomainObjectType resource.
	ObjectTypeName *string `field:"required" json:"objectTypeName" yaml:"objectTypeName"`
}

