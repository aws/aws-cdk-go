package previewawsautoscalingevents


// Type definition for CustomizedMetricSpecification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customizedMetricSpecification := &CustomizedMetricSpecification{
//   	Dimensions: []CustomizedMetricSpecificationItem{
//   		&CustomizedMetricSpecificationItem{
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	MetricName: []*string{
//   		jsii.String("metricName"),
//   	},
//   	Namespace: []*string{
//   		jsii.String("namespace"),
//   	},
//   	Statistic: []*string{
//   		jsii.String("statistic"),
//   	},
//   	Unit: []*string{
//   		jsii.String("unit"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_CustomizedMetricSpecification struct {
	// dimensions property.
	//
	// Specify an array of string values to match this event if the actual value of dimensions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Dimensions *[]*AWSAPICallViaCloudTrail_CustomizedMetricSpecificationItem `field:"optional" json:"dimensions" yaml:"dimensions"`
	// metricName property.
	//
	// Specify an array of string values to match this event if the actual value of metricName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MetricName *[]*string `field:"optional" json:"metricName" yaml:"metricName"`
	// namespace property.
	//
	// Specify an array of string values to match this event if the actual value of namespace is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Namespace *[]*string `field:"optional" json:"namespace" yaml:"namespace"`
	// statistic property.
	//
	// Specify an array of string values to match this event if the actual value of statistic is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Statistic *[]*string `field:"optional" json:"statistic" yaml:"statistic"`
	// unit property.
	//
	// Specify an array of string values to match this event if the actual value of unit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Unit *[]*string `field:"optional" json:"unit" yaml:"unit"`
}

