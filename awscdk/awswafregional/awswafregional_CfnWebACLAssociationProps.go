package awswafregional


// Properties for defining a `CfnWebACLAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebACLAssociationProps := &cfnWebACLAssociationProps{
//   	resourceArn: jsii.String("resourceArn"),
//   	webAclId: jsii.String("webAclId"),
//   }
//
type CfnWebACLAssociationProps struct {
	// The Amazon Resource Name (ARN) of the resource to protect with the web ACL.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// A unique identifier (ID) for the web ACL.
	WebAclId *string `field:"required" json:"webAclId" yaml:"webAclId"`
}

