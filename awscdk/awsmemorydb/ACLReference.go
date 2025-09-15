package awsmemorydb


// A reference to a ACL resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aCLReference := &ACLReference{
//   	AclArn: jsii.String("aclArn"),
//   	AclName: jsii.String("aclName"),
//   }
//
type ACLReference struct {
	// The ARN of the ACL resource.
	AclArn *string `field:"required" json:"aclArn" yaml:"aclArn"`
	// The ACLName of the ACL resource.
	AclName *string `field:"required" json:"aclName" yaml:"aclName"`
}

