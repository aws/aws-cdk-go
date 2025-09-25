package awsram


// A reference to a Permission resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   permissionReference := &PermissionReference{
//   	PermissionArn: jsii.String("permissionArn"),
//   }
//
type PermissionReference struct {
	// The Arn of the Permission resource.
	PermissionArn *string `field:"required" json:"permissionArn" yaml:"permissionArn"`
}

