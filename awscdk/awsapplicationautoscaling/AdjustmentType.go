package awsapplicationautoscaling


// How adjustment numbers are interpreted.
//
// Example:
//   var capacity scalableAttribute
//   var cpuUtilization metric
//
//
//   capacity.scaleOnMetric(jsii.String("ScaleToCPU"), &basicStepScalingPolicyProps{
//   	Metric: cpuUtilization,
//   	ScalingSteps: []scalingInterval{
//   		&scalingInterval{
//   			Upper: jsii.Number(10),
//   			Change: -jsii.Number(1),
//   		},
//   		&scalingInterval{
//   			Lower: jsii.Number(50),
//   			Change: +jsii.Number(1),
//   		},
//   		&scalingInterval{
//   			Lower: jsii.Number(70),
//   			Change: +jsii.Number(3),
//   		},
//   	},
//
//   	// Change this to AdjustmentType.PercentChangeInCapacity to interpret the
//   	// 'change' numbers before as percentages instead of capacity counts.
//   	AdjustmentType: appscaling.AdjustmentType_CHANGE_IN_CAPACITY,
//   })
//
type AdjustmentType string

const (
	// Add the adjustment number to the current capacity.
	//
	// A positive number increases capacity, a negative number decreases capacity.
	AdjustmentType_CHANGE_IN_CAPACITY AdjustmentType = "CHANGE_IN_CAPACITY"
	// Add this percentage of the current capacity to itself.
	//
	// The number must be between -100 and 100; a positive number increases
	// capacity and a negative number decreases it.
	AdjustmentType_PERCENT_CHANGE_IN_CAPACITY AdjustmentType = "PERCENT_CHANGE_IN_CAPACITY"
	// Make the capacity equal to the exact number given.
	AdjustmentType_EXACT_CAPACITY AdjustmentType = "EXACT_CAPACITY"
)

