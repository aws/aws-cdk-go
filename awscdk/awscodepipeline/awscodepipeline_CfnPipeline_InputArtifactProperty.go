package awscodepipeline


// Represents information about an artifact to be worked on, such as a test or build artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputArtifactProperty := &inputArtifactProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnPipeline_InputArtifactProperty struct {
	// The name of the artifact to be worked on (for example, "My App").
	//
	// The input artifact of an action must exactly match the output artifact declared in a preceding action, but the input artifact does not have to be the next action in strict sequence from the action that provided the output artifact. Actions in parallel can declare different output artifacts, which are in turn consumed by different following actions.
	Name *string `field:"required" json:"name" yaml:"name"`
}

