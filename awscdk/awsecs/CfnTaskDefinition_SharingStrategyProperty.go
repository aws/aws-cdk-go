package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharingStrategyProperty := &SharingStrategyProperty{
//   	NvidiaMps: &NvidiaMpsAllocationProperty{
//   		Memory: jsii.Number(123),
//
//   		// the properties below are optional
//   		MaxComputePercent: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-sharingstrategy.html
//
type CfnTaskDefinition_SharingStrategyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-sharingstrategy.html#cfn-ecs-taskdefinition-sharingstrategy-nvidiamps
	//
	NvidiaMps interface{} `field:"required" json:"nvidiaMps" yaml:"nvidiaMps"`
}

