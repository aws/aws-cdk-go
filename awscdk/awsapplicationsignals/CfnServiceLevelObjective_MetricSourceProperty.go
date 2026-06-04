package awsapplicationsignals


// Configuration for identifying the source of metrics for non-Application Signals services.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricSourceProperty := &MetricSourceProperty{
//   	MetricSourceKeyAttributes: map[string]*string{
//   		"metricSourceKeyAttributesKey": jsii.String("metricSourceKeyAttributes"),
//   	},
//
//   	// the properties below are optional
//   	MetricSourceAttributes: map[string]*string{
//   		"metricSourceAttributesKey": jsii.String("metricSourceAttributes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricsource.html
//
type CfnServiceLevelObjective_MetricSourceProperty struct {
	// Required attributes that identify the metric source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricsource.html#cfn-applicationsignals-servicelevelobjective-metricsource-metricsourcekeyattributes
	//
	MetricSourceKeyAttributes interface{} `field:"required" json:"metricSourceKeyAttributes" yaml:"metricSourceKeyAttributes"`
	// Optional additional attributes for the metric source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricsource.html#cfn-applicationsignals-servicelevelobjective-metricsource-metricsourceattributes
	//
	MetricSourceAttributes interface{} `field:"optional" json:"metricSourceAttributes" yaml:"metricSourceAttributes"`
}

