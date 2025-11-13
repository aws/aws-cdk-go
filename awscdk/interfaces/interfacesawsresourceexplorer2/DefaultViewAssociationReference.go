package interfacesawsresourceexplorer2


// A reference to a DefaultViewAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultViewAssociationReference := &DefaultViewAssociationReference{
//   	AssociatedAwsPrincipal: jsii.String("associatedAwsPrincipal"),
//   }
//
type DefaultViewAssociationReference struct {
	// The AssociatedAwsPrincipal of the DefaultViewAssociation resource.
	AssociatedAwsPrincipal *string `field:"required" json:"associatedAwsPrincipal" yaml:"associatedAwsPrincipal"`
}

