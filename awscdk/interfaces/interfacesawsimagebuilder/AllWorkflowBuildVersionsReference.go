package interfacesawsimagebuilder


// A reference to a AllWorkflowBuildVersions resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allWorkflowBuildVersionsReference := &AllWorkflowBuildVersionsReference{
//   	AllWorkflowBuildVersionsArn: jsii.String("allWorkflowBuildVersionsArn"),
//   }
//
type AllWorkflowBuildVersionsReference struct {
	// The Arn of the AllWorkflowBuildVersions resource.
	AllWorkflowBuildVersionsArn *string `field:"required" json:"allWorkflowBuildVersionsArn" yaml:"allWorkflowBuildVersionsArn"`
}

