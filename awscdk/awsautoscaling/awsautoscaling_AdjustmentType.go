package awsautoscaling


// How adjustment numbers are interpreted.
//
// Example:
//   var autoScalingGroup autoScalingGroup
//
//
//   workerUtilizationMetric := cloudwatch.NewMetric(&metricProps{
//   	namespace: jsii.String("MyService"),
//   	metricName: jsii.String("WorkerUtilization"),
//   })
//
//   autoScalingGroup.scaleOnMetric(jsii.String("ScaleToCPU"), &basicStepScalingPolicyProps{
//   	metric: workerUtilizationMetric,
//   	scalingSteps: []scalingInterval{
//   		&scalingInterval{
//   			upper: jsii.Number(10),
//   			change: -jsii.Number(1),
//   		},
//   		&scalingInterval{
//   			lower: jsii.Number(50),
//   			change: +jsii.Number(1),
//   		},
//   		&scalingInterval{
//   			lower: jsii.Number(70),
//   			change: +jsii.Number(3),
//   		},
//   	},
//
//   	// Change this to AdjustmentType.PERCENT_CHANGE_IN_CAPACITY to interpret the
//   	// 'change' numbers before as percentages instead of capacity counts.
//   	adjustmentType: autoscaling.adjustmentType_CHANGE_IN_CAPACITY,
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

