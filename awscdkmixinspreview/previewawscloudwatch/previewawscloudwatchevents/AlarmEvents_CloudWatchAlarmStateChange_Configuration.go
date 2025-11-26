package previewawscloudwatchevents


// Type definition for Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   configuration := &Configuration{
//   	ActionsSuppressor: []*string{
//   		jsii.String("actionsSuppressor"),
//   	},
//   	ActionsSuppressorExtensionPeriod: []*string{
//   		jsii.String("actionsSuppressorExtensionPeriod"),
//   	},
//   	ActionsSuppressorWaitPeriod: []*string{
//   		jsii.String("actionsSuppressorWaitPeriod"),
//   	},
//   	AlarmRule: []*string{
//   		jsii.String("alarmRule"),
//   	},
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	Metrics: []ConfigurationItem{
//   		&ConfigurationItem{
//   			Id: []*string{
//   				jsii.String("id"),
//   			},
//   			MetricStat: &MetricStat{
//   				Metric: &Metric{
//   					Dimensions: []*string{
//   						jsii.String("dimensions"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					Namespace: []*string{
//   						jsii.String("namespace"),
//   					},
//   				},
//   				Period: []*string{
//   					jsii.String("period"),
//   				},
//   				Stat: []*string{
//   					jsii.String("stat"),
//   				},
//   			},
//   			ReturnData: []*string{
//   				jsii.String("returnData"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type AlarmEvents_CloudWatchAlarmStateChange_Configuration struct {
	// actionsSuppressor property.
	//
	// Specify an array of string values to match this event if the actual value of actionsSuppressor is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActionsSuppressor *[]*string `field:"optional" json:"actionsSuppressor" yaml:"actionsSuppressor"`
	// actionsSuppressorExtensionPeriod property.
	//
	// Specify an array of string values to match this event if the actual value of actionsSuppressorExtensionPeriod is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActionsSuppressorExtensionPeriod *[]*string `field:"optional" json:"actionsSuppressorExtensionPeriod" yaml:"actionsSuppressorExtensionPeriod"`
	// actionsSuppressorWaitPeriod property.
	//
	// Specify an array of string values to match this event if the actual value of actionsSuppressorWaitPeriod is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActionsSuppressorWaitPeriod *[]*string `field:"optional" json:"actionsSuppressorWaitPeriod" yaml:"actionsSuppressorWaitPeriod"`
	// alarmRule property.
	//
	// Specify an array of string values to match this event if the actual value of alarmRule is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlarmRule *[]*string `field:"optional" json:"alarmRule" yaml:"alarmRule"`
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// metrics property.
	//
	// Specify an array of string values to match this event if the actual value of metrics is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Metrics *[]*AlarmEvents_CloudWatchAlarmStateChange_ConfigurationItem `field:"optional" json:"metrics" yaml:"metrics"`
}

