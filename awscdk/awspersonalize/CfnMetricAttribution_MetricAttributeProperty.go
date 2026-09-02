package awspersonalize


// A metric attribute for the metric attribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricAttributeProperty := &MetricAttributeProperty{
//   	EventType: jsii.String("eventType"),
//   	Expression: jsii.String("expression"),
//   	MetricName: jsii.String("metricName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-metricattribute.html
//
type CfnMetricAttribution_MetricAttributeProperty struct {
	// The metric's event type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-metricattribute.html#cfn-personalize-metricattribution-metricattribute-eventtype
	//
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// The attribute's expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-metricattribute.html#cfn-personalize-metricattribution-metricattribute-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The metric's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-metricattribution-metricattribute.html#cfn-personalize-metricattribution-metricattribute-metricname
	//
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
}

