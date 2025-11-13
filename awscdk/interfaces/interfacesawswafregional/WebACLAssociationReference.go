package interfacesawswafregional


// A reference to a WebACLAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webACLAssociationReference := &WebACLAssociationReference{
//   	WebAclAssociationId: jsii.String("webAclAssociationId"),
//   }
//
type WebACLAssociationReference struct {
	// The Id of the WebACLAssociation resource.
	WebAclAssociationId *string `field:"required" json:"webAclAssociationId" yaml:"webAclAssociationId"`
}

