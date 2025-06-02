package awsecs


// A floating-point percentage of the desired number of tasks to place and keep running in the task set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scaleProperty := &ScaleProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-scale.html
//
type CfnTaskSet_ScaleProperty struct {
	// The unit of measure for the scale value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-scale.html#cfn-ecs-taskset-scale-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// The value, specified as a percent total of a service's `desiredCount` , to scale the task set.
	//
	// Accepted values are numbers between 0 and 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-scale.html#cfn-ecs-taskset-scale-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

