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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-hporesourceconfig.html
//
type CfnSolution_HpoResourceConfigProperty struct {
	// The maximum number of training jobs when you create a solution version.
	//
	// The maximum value for `maxNumberOfTrainingJobs` is `40` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-hporesourceconfig.html#cfn-personalize-solution-hporesourceconfig-maxnumberoftrainingjobs
	//
	MaxNumberOfTrainingJobs *string `field:"optional" json:"maxNumberOfTrainingJobs" yaml:"maxNumberOfTrainingJobs"`
	// The maximum number of parallel training jobs when you create a solution version.
	//
	// The maximum value for `maxParallelTrainingJobs` is `10` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-hporesourceconfig.html#cfn-personalize-solution-hporesourceconfig-maxparalleltrainingjobs
	//
	MaxParallelTrainingJobs *string `field:"optional" json:"maxParallelTrainingJobs" yaml:"maxParallelTrainingJobs"`
}

