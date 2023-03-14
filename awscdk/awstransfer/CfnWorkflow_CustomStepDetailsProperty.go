package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customStepDetailsProperty := &CustomStepDetailsProperty{
//   	Name: jsii.String("name"),
//   	SourceFileLocation: jsii.String("sourceFileLocation"),
//   	Target: jsii.String("target"),
//   	TimeoutSeconds: jsii.Number(123),
//   }
//
type CfnWorkflow_CustomStepDetailsProperty struct {
	// `CfnWorkflow.CustomStepDetailsProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnWorkflow.CustomStepDetailsProperty.SourceFileLocation`.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
	// `CfnWorkflow.CustomStepDetailsProperty.Target`.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// `CfnWorkflow.CustomStepDetailsProperty.TimeoutSeconds`.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

