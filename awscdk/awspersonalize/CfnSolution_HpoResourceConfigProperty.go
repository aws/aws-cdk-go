package awspersonalize


// Describes the resource configuration for hyperparameter optimization (HPO).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hpoResourceConfigProperty := &HpoResourceConfigProperty{
//   	MaxNumberOfTrainingJobs: jsii.String("maxNumberOfTrainingJobs"),
//   	MaxParallelTrainingJobs: jsii.String("maxParallelTrainingJobs"),
//   }
//
type CfnSolution_HpoResourceConfigProperty struct {
	// The maximum number of training jobs when you create a solution version.
	//
	// The maximum value for `maxNumberOfTrainingJobs` is `40` .
	MaxNumberOfTrainingJobs *string `field:"optional" json:"maxNumberOfTrainingJobs" yaml:"maxNumberOfTrainingJobs"`
	// The maximum number of parallel training jobs when you create a solution version.
	//
	// The maximum value for `maxParallelTrainingJobs` is `10` .
	MaxParallelTrainingJobs *string `field:"optional" json:"maxParallelTrainingJobs" yaml:"maxParallelTrainingJobs"`
}

