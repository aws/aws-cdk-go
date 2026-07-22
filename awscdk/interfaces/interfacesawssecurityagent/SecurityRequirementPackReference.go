package interfacesawssecurityagent


// A reference to a SecurityRequirementPack resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityRequirementPackReference := &SecurityRequirementPackReference{
//   	PackId: jsii.String("packId"),
//   }
//
type SecurityRequirementPackReference struct {
	// The PackId of the SecurityRequirementPack resource.
	PackId *string `field:"required" json:"packId" yaml:"packId"`
}

