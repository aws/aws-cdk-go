package interfacesawsdatazone


// A reference to a Owner resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ownerReference := &OwnerReference{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EntityIdentifier: jsii.String("entityIdentifier"),
//   	EntityType: jsii.String("entityType"),
//   	OwnerIdentifier: jsii.String("ownerIdentifier"),
//   	OwnerType: jsii.String("ownerType"),
//   }
//
type OwnerReference struct {
	// The DomainIdentifier of the Owner resource.
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The EntityIdentifier of the Owner resource.
	EntityIdentifier *string `field:"required" json:"entityIdentifier" yaml:"entityIdentifier"`
	// The EntityType of the Owner resource.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// The OwnerIdentifier of the Owner resource.
	OwnerIdentifier *string `field:"required" json:"ownerIdentifier" yaml:"ownerIdentifier"`
	// The OwnerType of the Owner resource.
	OwnerType *string `field:"required" json:"ownerType" yaml:"ownerType"`
}

