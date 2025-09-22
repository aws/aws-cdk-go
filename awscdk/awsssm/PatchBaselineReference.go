package awsssm


// A reference to a PatchBaseline resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   patchBaselineReference := &PatchBaselineReference{
//   	PatchBaselineId: jsii.String("patchBaselineId"),
//   }
//
type PatchBaselineReference struct {
	// The Id of the PatchBaseline resource.
	PatchBaselineId *string `field:"required" json:"patchBaselineId" yaml:"patchBaselineId"`
}

