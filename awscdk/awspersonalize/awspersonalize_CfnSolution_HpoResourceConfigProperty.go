package awspersonalize


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hpoResourceConfigProperty := &hpoResourceConfigProperty{
//   	maxNumberOfTrainingJobs: jsii.String("maxNumberOfTrainingJobs"),
//   	maxParallelTrainingJobs: jsii.String("maxParallelTrainingJobs"),
//   }
//
type CfnSolution_HpoResourceConfigProperty struct {
	// `CfnSolution.HpoResourceConfigProperty.MaxNumberOfTrainingJobs`.
	MaxNumberOfTrainingJobs *string `field:"optional" json:"maxNumberOfTrainingJobs" yaml:"maxNumberOfTrainingJobs"`
	// `CfnSolution.HpoResourceConfigProperty.MaxParallelTrainingJobs`.
	MaxParallelTrainingJobs *string `field:"optional" json:"maxParallelTrainingJobs" yaml:"maxParallelTrainingJobs"`
}

