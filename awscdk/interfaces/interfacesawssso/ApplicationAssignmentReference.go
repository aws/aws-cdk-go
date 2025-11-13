package interfacesawssso


// A reference to a ApplicationAssignment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationAssignmentReference := &ApplicationAssignmentReference{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	PrincipalId: jsii.String("principalId"),
//   	PrincipalType: jsii.String("principalType"),
//   }
//
type ApplicationAssignmentReference struct {
	// The ApplicationArn of the ApplicationAssignment resource.
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// The PrincipalId of the ApplicationAssignment resource.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The PrincipalType of the ApplicationAssignment resource.
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
}

