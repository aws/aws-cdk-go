package awswafv2


// A reference to a WebACL resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webACLReference := &WebACLReference{
//   	Scope: jsii.String("scope"),
//   	WebAclArn: jsii.String("webAclArn"),
//   	WebAclId: jsii.String("webAclId"),
//   	WebAclName: jsii.String("webAclName"),
//   }
//
type WebACLReference struct {
	// The Scope of the WebACL resource.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// The ARN of the WebACL resource.
	WebAclArn *string `field:"required" json:"webAclArn" yaml:"webAclArn"`
	// The Id of the WebACL resource.
	WebAclId *string `field:"required" json:"webAclId" yaml:"webAclId"`
	// The Name of the WebACL resource.
	WebAclName *string `field:"required" json:"webAclName" yaml:"webAclName"`
}

