package awssagemaker


// The workload specification for benchmark tool configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   workloadSpecProperty := &WorkloadSpecProperty{
//   	Inline: jsii.String("inline"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-workloadspec.html
//
type CfnAIWorkloadConfigPropsMixin_WorkloadSpecProperty struct {
	// An inline YAML or JSON string that defines benchmark parameters.
	//
	// The service validates the document against its own benchmark schema: it must declare a benchmark object whose type member matches the pattern ^(aiperf)$.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-workloadspec.html#cfn-sagemaker-aiworkloadconfig-workloadspec-inline
	//
	Inline *string `field:"optional" json:"inline" yaml:"inline"`
}

