package awswafregional


// A reference to a WebACL resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webACLReference := &WebACLReference{
//   	WebAclId: jsii.String("webAclId"),
//   }
//
type WebACLReference struct {
	// The Id of the WebACL resource.
	WebAclId *string `field:"required" json:"webAclId" yaml:"webAclId"`
}

