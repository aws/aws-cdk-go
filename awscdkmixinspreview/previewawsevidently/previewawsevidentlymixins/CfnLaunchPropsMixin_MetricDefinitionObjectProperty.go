package previewawsevidentlymixins


// This structure defines a metric that you want to use to evaluate the variations during a launch or experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricDefinitionObjectProperty := &MetricDefinitionObjectProperty{
//   	EntityIdKey: jsii.String("entityIdKey"),
//   	EventPattern: jsii.String("eventPattern"),
//   	MetricName: jsii.String("metricName"),
//   	UnitLabel: jsii.String("unitLabel"),
//   	ValueKey: jsii.String("valueKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-metricdefinitionobject.html
//
type CfnLaunchPropsMixin_MetricDefinitionObjectProperty struct {
	// The entity, such as a user or session, that does an action that causes a metric value to be recorded.
	//
	// An example is `userDetails.userID` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-metricdefinitionobject.html#cfn-evidently-launch-metricdefinitionobject-entityidkey
	//
	EntityIdKey *string `field:"optional" json:"entityIdKey" yaml:"entityIdKey"`
	// The EventBridge event pattern that defines how the metric is recorded.
	//
	// For more information about EventBridge event patterns, see [Amazon EventBridge event patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-patterns.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-metricdefinitionobject.html#cfn-evidently-launch-metricdefinitionobject-eventpattern
	//
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A name for the metric.
	//
	// It can include up to 255 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-metricdefinitionobject.html#cfn-evidently-launch-metricdefinitionobject-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// A label for the units that the metric is measuring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-metricdefinitionobject.html#cfn-evidently-launch-metricdefinitionobject-unitlabel
	//
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
	// The value that is tracked to produce the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-metricdefinitionobject.html#cfn-evidently-launch-metricdefinitionobject-valuekey
	//
	ValueKey *string `field:"optional" json:"valueKey" yaml:"valueKey"`
}

