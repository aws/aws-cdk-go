package previewawsevidentlymixins


// Use this structure to tell Evidently whether higher or lower values are desired for a metric that is used in an experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricGoalObjectProperty := &MetricGoalObjectProperty{
//   	DesiredChange: jsii.String("desiredChange"),
//   	EntityIdKey: jsii.String("entityIdKey"),
//   	EventPattern: jsii.String("eventPattern"),
//   	MetricName: jsii.String("metricName"),
//   	UnitLabel: jsii.String("unitLabel"),
//   	ValueKey: jsii.String("valueKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-metricgoalobject.html
//
type CfnExperimentPropsMixin_MetricGoalObjectProperty struct {
	// `INCREASE` means that a variation with a higher number for this metric is performing better.
	//
	// `DECREASE` means that a variation with a lower number for this metric is performing better.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-metricgoalobject.html#cfn-evidently-experiment-metricgoalobject-desiredchange
	//
	DesiredChange *string `field:"optional" json:"desiredChange" yaml:"desiredChange"`
	// The entity, such as a user or session, that does an action that causes a metric value to be recorded.
	//
	// An example is `userDetails.userID` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-metricgoalobject.html#cfn-evidently-experiment-metricgoalobject-entityidkey
	//
	EntityIdKey *string `field:"optional" json:"entityIdKey" yaml:"entityIdKey"`
	// The EventBridge event pattern that defines how the metric is recorded.
	//
	// For more information about EventBridge event patterns, see [Amazon EventBridge event patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-patterns.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-metricgoalobject.html#cfn-evidently-experiment-metricgoalobject-eventpattern
	//
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A name for the metric.
	//
	// It can include up to 255 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-metricgoalobject.html#cfn-evidently-experiment-metricgoalobject-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// A label for the units that the metric is measuring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-metricgoalobject.html#cfn-evidently-experiment-metricgoalobject-unitlabel
	//
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
	// The JSON path to reference the numerical metric value in the event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-metricgoalobject.html#cfn-evidently-experiment-metricgoalobject-valuekey
	//
	ValueKey *string `field:"optional" json:"valueKey" yaml:"valueKey"`
}

