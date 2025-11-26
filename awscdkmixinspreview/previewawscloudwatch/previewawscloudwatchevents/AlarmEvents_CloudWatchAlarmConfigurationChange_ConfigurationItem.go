package previewawscloudwatchevents


// Type definition for ConfigurationItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   configurationItem := &ConfigurationItem{
//   	Id: []*string{
//   		jsii.String("id"),
//   	},
//   	MetricStat: &MetricStat{
//   		Metric: &Metric{
//   			Dimensions: []*string{
//   				jsii.String("dimensions"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Namespace: []*string{
//   				jsii.String("namespace"),
//   			},
//   		},
//   		Period: []*string{
//   			jsii.String("period"),
//   		},
//   		Stat: []*string{
//   			jsii.String("stat"),
//   		},
//   	},
//   	ReturnData: []*string{
//   		jsii.String("returnData"),
//   	},
//   }
//
// Experimental.
type AlarmEvents_CloudWatchAlarmConfigurationChange_ConfigurationItem struct {
	// id property.
	//
	// Specify an array of string values to match this event if the actual value of id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
	// metricStat property.
	//
	// Specify an array of string values to match this event if the actual value of metricStat is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MetricStat *AlarmEvents_CloudWatchAlarmConfigurationChange_MetricStat `field:"optional" json:"metricStat" yaml:"metricStat"`
	// returnData property.
	//
	// Specify an array of string values to match this event if the actual value of returnData is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReturnData *[]*string `field:"optional" json:"returnData" yaml:"returnData"`
}

