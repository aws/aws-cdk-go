package awssagemaker


// The benchmark tool configuration for an AI workload.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIWorkloadConfigsProperty := &AIWorkloadConfigsProperty{
//   	WorkloadSpec: &WorkloadSpecProperty{
//   		Inline: jsii.String("inline"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-aiworkloadconfigs.html
//
type CfnAIWorkloadConfig_AIWorkloadConfigsProperty struct {
	// The workload specification for benchmark tool configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-aiworkloadconfigs.html#cfn-sagemaker-aiworkloadconfig-aiworkloadconfigs-workloadspec
	//
	WorkloadSpec interface{} `field:"required" json:"workloadSpec" yaml:"workloadSpec"`
}

