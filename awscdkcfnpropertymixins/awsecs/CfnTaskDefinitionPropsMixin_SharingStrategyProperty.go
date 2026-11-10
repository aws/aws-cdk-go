package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sharingStrategyProperty := &SharingStrategyProperty{
//   	NvidiaMps: &NvidiaMpsAllocationProperty{
//   		MaxComputePercent: jsii.Number(123),
//   		Memory: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-sharingstrategy.html
//
type CfnTaskDefinitionPropsMixin_SharingStrategyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-sharingstrategy.html#cfn-ecs-taskdefinition-sharingstrategy-nvidiamps
	//
	NvidiaMps interface{} `field:"optional" json:"nvidiaMps" yaml:"nvidiaMps"`
}

