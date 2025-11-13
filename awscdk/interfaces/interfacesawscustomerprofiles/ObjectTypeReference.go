package interfacesawscustomerprofiles


// A reference to a ObjectType resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectTypeReference := &ObjectTypeReference{
//   	DomainName: jsii.String("domainName"),
//   	ObjectTypeName: jsii.String("objectTypeName"),
//   }
//
type ObjectTypeReference struct {
	// The DomainName of the ObjectType resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The ObjectTypeName of the ObjectType resource.
	ObjectTypeName *string `field:"required" json:"objectTypeName" yaml:"objectTypeName"`
}

