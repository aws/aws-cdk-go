package awsevidently


// Use this structure to tell Evidently whether higher or lower values are desired for a metric that is used in an experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricGoalObjectProperty := &metricGoalObjectProperty{
//   	desiredChange: jsii.String("desiredChange"),
//   	entityIdKey: jsii.String("entityIdKey"),
//   	metricName: jsii.String("metricName"),
//   	valueKey: jsii.String("valueKey"),
//
//   	// the properties below are optional
//   	eventPattern: jsii.String("eventPattern"),
//   	unitLabel: jsii.String("unitLabel"),
//   }
//
type CfnExperiment_MetricGoalObjectProperty struct {
	// `INCREASE` means that a variation with a higher number for this metric is performing better.
	//
	// `DECREASE` means that a variation with a lower number for this metric is performing better.
	DesiredChange *string `field:"required" json:"desiredChange" yaml:"desiredChange"`
	// The entity, such as a user or session, that does an action that causes a metric value to be recorded.
	//
	// An example is `userDetails.userID` .
	EntityIdKey *string `field:"required" json:"entityIdKey" yaml:"entityIdKey"`
	// A name for the metric.
	//
	// It can include up to 255 characters.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The JSON path to reference the numerical metric value in the event.
	ValueKey *string `field:"required" json:"valueKey" yaml:"valueKey"`
	// The EventBridge event pattern that defines how the metric is recorded.
	//
	// For more information about EventBridge event patterns, see [Amazon EventBridge event patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-patterns.html) .
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A label for the units that the metric is measuring.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
}

