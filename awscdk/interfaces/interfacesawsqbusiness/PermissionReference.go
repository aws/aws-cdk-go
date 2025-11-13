package interfacesawsqbusiness


// A reference to a Permission resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   permissionReference := &PermissionReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	StatementId: jsii.String("statementId"),
//   }
//
type PermissionReference struct {
	// The ApplicationId of the Permission resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The StatementId of the Permission resource.
	StatementId *string `field:"required" json:"statementId" yaml:"statementId"`
}

