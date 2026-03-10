package previewawsautoscalingevents


// Type definition for TargetTrackingConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   targetTrackingConfiguration := &TargetTrackingConfiguration{
//   	CustomizedMetricSpecification: &CustomizedMetricSpecification{
//   		Dimensions: []CustomizedMetricSpecificationItem{
//   			&CustomizedMetricSpecificationItem{
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   		MetricName: []*string{
//   			jsii.String("metricName"),
//   		},
//   		Namespace: []*string{
//   			jsii.String("namespace"),
//   		},
//   		Statistic: []*string{
//   			jsii.String("statistic"),
//   		},
//   		Unit: []*string{
//   			jsii.String("unit"),
//   		},
//   	},
//   	PredefinedMetricSpecification: &PredefinedMetricSpecification{
//   		PredefinedMetricType: []*string{
//   			jsii.String("predefinedMetricType"),
//   		},
//   	},
//   	TargetValue: []*string{
//   		jsii.String("targetValue"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_TargetTrackingConfiguration struct {
	// customizedMetricSpecification property.
	//
	// Specify an array of string values to match this event if the actual value of customizedMetricSpecification is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CustomizedMetricSpecification *AWSAPICallViaCloudTrail_CustomizedMetricSpecification `field:"optional" json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// predefinedMetricSpecification property.
	//
	// Specify an array of string values to match this event if the actual value of predefinedMetricSpecification is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PredefinedMetricSpecification *AWSAPICallViaCloudTrail_PredefinedMetricSpecification `field:"optional" json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
	// targetValue property.
	//
	// Specify an array of string values to match this event if the actual value of targetValue is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetValue *[]*string `field:"optional" json:"targetValue" yaml:"targetValue"`
}

