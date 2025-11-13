package interfacesawswafv2


// A reference to a WebACLAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webACLAssociationReference := &WebACLAssociationReference{
//   	ResourceArn: jsii.String("resourceArn"),
//   	WebAclArn: jsii.String("webAclArn"),
//   }
//
type WebACLAssociationReference struct {
	// The ResourceArn of the WebACLAssociation resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The WebACLArn of the WebACLAssociation resource.
	WebAclArn *string `field:"required" json:"webAclArn" yaml:"webAclArn"`
}

