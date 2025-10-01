package awssigner


// A reference to a ProfilePermission resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   profilePermissionReference := &ProfilePermissionReference{
//   	ProfileName: jsii.String("profileName"),
//   	StatementId: jsii.String("statementId"),
//   }
//
type ProfilePermissionReference struct {
	// The ProfileName of the ProfilePermission resource.
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// The StatementId of the ProfilePermission resource.
	StatementId *string `field:"required" json:"statementId" yaml:"statementId"`
}

