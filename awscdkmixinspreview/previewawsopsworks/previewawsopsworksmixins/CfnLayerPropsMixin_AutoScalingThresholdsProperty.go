package previewawsopsworksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autoScalingThresholdsProperty := &AutoScalingThresholdsProperty{
//   	CpuThreshold: jsii.Number(123),
//   	IgnoreMetricsTime: jsii.Number(123),
//   	InstanceCount: jsii.Number(123),
//   	LoadThreshold: jsii.Number(123),
//   	MemoryThreshold: jsii.Number(123),
//   	ThresholdsWaitTime: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-autoscalingthresholds.html
//
type CfnLayerPropsMixin_AutoScalingThresholdsProperty struct {
	// The CPU utilization threshold, as a percent of the available CPU.
	//
	// A value of -1 disables the threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-autoscalingthresholds.html#cfn-opsworks-layer-autoscalingthresholds-cputhreshold
	//
	CpuThreshold *float64 `field:"optional" json:"cpuThreshold" yaml:"cpuThreshold"`
	// The amount of time (in minutes) after a scaling event occurs that OpsWorks Stacks should ignore metrics and suppress additional scaling events.
	//
	// For example, OpsWorks Stacks adds new instances following an upscaling event but the instances won't start reducing the load until they have been booted and configured. There is no point in raising additional scaling events during that operation, which typically takes several minutes. `IgnoreMetricsTime` allows you to direct OpsWorks Stacks to suppress scaling events long enough to get the new instances online.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-autoscalingthresholds.html#cfn-opsworks-layer-autoscalingthresholds-ignoremetricstime
	//
	IgnoreMetricsTime *float64 `field:"optional" json:"ignoreMetricsTime" yaml:"ignoreMetricsTime"`
	// The number of instances to add or remove when the load exceeds a threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-autoscalingthresholds.html#cfn-opsworks-layer-autoscalingthresholds-instancecount
	//
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// The load threshold.
	//
	// A value of -1 disables the threshold. For more information about how load is computed, see [Load (computing)](https://docs.aws.amazon.com/http://en.wikipedia.org/wiki/Load_%28computing%29) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-autoscalingthresholds.html#cfn-opsworks-layer-autoscalingthresholds-loadthreshold
	//
	LoadThreshold *float64 `field:"optional" json:"loadThreshold" yaml:"loadThreshold"`
	// The memory utilization threshold, as a percent of the available memory.
	//
	// A value of -1 disables the threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-autoscalingthresholds.html#cfn-opsworks-layer-autoscalingthresholds-memorythreshold
	//
	MemoryThreshold *float64 `field:"optional" json:"memoryThreshold" yaml:"memoryThreshold"`
	// The amount of time, in minutes, that the load must exceed a threshold before more instances are added or removed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-autoscalingthresholds.html#cfn-opsworks-layer-autoscalingthresholds-thresholdswaittime
	//
	ThresholdsWaitTime *float64 `field:"optional" json:"thresholdsWaitTime" yaml:"thresholdsWaitTime"`
}

