package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deleteStepDetailsProperty := &deleteStepDetailsProperty{
//   	name: jsii.String("name"),
//   	sourceFileLocation: jsii.String("sourceFileLocation"),
//   }
//
type CfnWorkflow_DeleteStepDetailsProperty struct {
	// `CfnWorkflow.DeleteStepDetailsProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnWorkflow.DeleteStepDetailsProperty.SourceFileLocation`.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
}

