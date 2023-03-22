package awsevidently


// This structure defines a metric that you want to use to evaluate the variations during a launch or experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDefinitionObjectProperty := &metricDefinitionObjectProperty{
//   	entityIdKey: jsii.String("entityIdKey"),
//   	metricName: jsii.String("metricName"),
//   	valueKey: jsii.String("valueKey"),
//
//   	// the properties below are optional
//   	eventPattern: jsii.String("eventPattern"),
//   	unitLabel: jsii.String("unitLabel"),
//   }
//
type CfnLaunch_MetricDefinitionObjectProperty struct {
	// The entity, such as a user or session, that does an action that causes a metric value to be recorded.
	//
	// An example is `userDetails.userID` .
	EntityIdKey *string `field:"required" json:"entityIdKey" yaml:"entityIdKey"`
	// A name for the metric.
	//
	// It can include up to 255 characters.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The value that is tracked to produce the metric.
	ValueKey *string `field:"required" json:"valueKey" yaml:"valueKey"`
	// The EventBridge event pattern that defines how the metric is recorded.
	//
	// For more information about EventBridge event patterns, see [Amazon EventBridge event patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-patterns.html) .
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A label for the units that the metric is measuring.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
}

