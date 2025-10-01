package awssagemaker


// Defines the traffic routing strategy during an endpoint deployment to shift traffic from the old fleet to the new fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficRoutingConfigProperty := &TrafficRoutingConfigProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	CanarySize: &CapacitySizeProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	LinearStepSize: &CapacitySizeProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	WaitIntervalInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-trafficroutingconfig.html
//
type CfnEndpoint_TrafficRoutingConfigProperty struct {
	// Traffic routing strategy type.
	//
	// - `ALL_AT_ONCE` : Endpoint traffic shifts to the new fleet in a single step.
	// - `CANARY` : Endpoint traffic shifts to the new fleet in two steps. The first step is the canary, which is a small portion of the traffic. The second step is the remainder of the traffic.
	// - `LINEAR` : Endpoint traffic shifts to the new fleet in n steps of a configurable size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-trafficroutingconfig.html#cfn-sagemaker-endpoint-trafficroutingconfig-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Batch size for the first step to turn on traffic on the new endpoint fleet.
	//
	// `Value` must be less than or equal to 50% of the variant's total instance count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-trafficroutingconfig.html#cfn-sagemaker-endpoint-trafficroutingconfig-canarysize
	//
	CanarySize interface{} `field:"optional" json:"canarySize" yaml:"canarySize"`
	// Batch size for each step to turn on traffic on the new endpoint fleet.
	//
	// `Value` must be 10-50% of the variant's total instance count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-trafficroutingconfig.html#cfn-sagemaker-endpoint-trafficroutingconfig-linearstepsize
	//
	LinearStepSize interface{} `field:"optional" json:"linearStepSize" yaml:"linearStepSize"`
	// The waiting time (in seconds) between incremental steps to turn on traffic on the new endpoint fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-trafficroutingconfig.html#cfn-sagemaker-endpoint-trafficroutingconfig-waitintervalinseconds
	//
	WaitIntervalInSeconds *float64 `field:"optional" json:"waitIntervalInSeconds" yaml:"waitIntervalInSeconds"`
}

