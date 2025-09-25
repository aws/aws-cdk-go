package awslambda


// A reference to a Permission resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   permissionReference := &PermissionReference{
//   	FunctionName: jsii.String("functionName"),
//   	PermissionId: jsii.String("permissionId"),
//   }
//
type PermissionReference struct {
	// The FunctionName of the Permission resource.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The Id of the Permission resource.
	PermissionId *string `field:"required" json:"permissionId" yaml:"permissionId"`
}

