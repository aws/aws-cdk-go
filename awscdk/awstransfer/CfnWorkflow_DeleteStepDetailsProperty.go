package awstransfer


// An object that contains the name and file location for a file being deleted by a workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deleteStepDetailsProperty := &DeleteStepDetailsProperty{
//   	Name: jsii.String("name"),
//   	SourceFileLocation: jsii.String("sourceFileLocation"),
//   }
//
type CfnWorkflow_DeleteStepDetailsProperty struct {
	// The name of the step, used as an identifier.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow.
	//
	// - To use the previous file as the input, enter `${previous.file}` . In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value.
	// - To use the originally uploaded file location as input for this step, enter `${original.file}` .
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
}

