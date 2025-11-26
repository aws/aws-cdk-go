package previewawscloudwatchevents


// Type definition for Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   configuration := &Configuration{
//   	ActionsEnabled: []*string{
//   		jsii.String("actionsEnabled"),
//   	},
//   	ActionsSuppressor: []*string{
//   		jsii.String("actionsSuppressor"),
//   	},
//   	ActionsSuppressorExtensionPeriod: []*string{
//   		jsii.String("actionsSuppressorExtensionPeriod"),
//   	},
//   	ActionsSuppressorWaitPeriod: []*string{
//   		jsii.String("actionsSuppressorWaitPeriod"),
//   	},
//   	AlarmActions: []*string{
//   		jsii.String("alarmActions"),
//   	},
//   	AlarmName: []*string{
//   		jsii.String("alarmName"),
//   	},
//   	AlarmRule: []*string{
//   		jsii.String("alarmRule"),
//   	},
//   	ComparisonOperator: []*string{
//   		jsii.String("comparisonOperator"),
//   	},
//   	DatapointsToAlarm: []*string{
//   		jsii.String("datapointsToAlarm"),
//   	},
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	EvaluateLowSampleCountPercentile: []*string{
//   		jsii.String("evaluateLowSampleCountPercentile"),
//   	},
//   	EvaluationPeriods: []*string{
//   		jsii.String("evaluationPeriods"),
//   	},
//   	InsufficientDataActions: []*string{
//   		jsii.String("insufficientDataActions"),
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
//   	OkActions: []*string{
//   		jsii.String("okActions"),
//   	},
//   	Threshold: []*string{
//   		jsii.String("threshold"),
//   	},
//   	Timestamp: []*string{
//   		jsii.String("timestamp"),
//   	},
//   	TreatMissingData: []*string{
//   		jsii.String("treatMissingData"),
//   	},
//   }
//
// Experimental.
type AlarmEvents_CloudWatchAlarmConfigurationChange_Configuration struct {
	// actionsEnabled property.
	//
	// Specify an array of string values to match this event if the actual value of actionsEnabled is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActionsEnabled *[]*string `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
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
	// alarmActions property.
	//
	// Specify an array of string values to match this event if the actual value of alarmActions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlarmActions *[]*string `field:"optional" json:"alarmActions" yaml:"alarmActions"`
	// alarmName property.
	//
	// Specify an array of string values to match this event if the actual value of alarmName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Alarm reference.
	//
	// Experimental.
	AlarmName *[]*string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// alarmRule property.
	//
	// Specify an array of string values to match this event if the actual value of alarmRule is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlarmRule *[]*string `field:"optional" json:"alarmRule" yaml:"alarmRule"`
	// comparisonOperator property.
	//
	// Specify an array of string values to match this event if the actual value of comparisonOperator is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ComparisonOperator *[]*string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// datapointsToAlarm property.
	//
	// Specify an array of string values to match this event if the actual value of datapointsToAlarm is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatapointsToAlarm *[]*string `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// evaluateLowSampleCountPercentile property.
	//
	// Specify an array of string values to match this event if the actual value of evaluateLowSampleCountPercentile is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EvaluateLowSampleCountPercentile *[]*string `field:"optional" json:"evaluateLowSampleCountPercentile" yaml:"evaluateLowSampleCountPercentile"`
	// evaluationPeriods property.
	//
	// Specify an array of string values to match this event if the actual value of evaluationPeriods is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EvaluationPeriods *[]*string `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// insufficientDataActions property.
	//
	// Specify an array of string values to match this event if the actual value of insufficientDataActions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InsufficientDataActions *[]*string `field:"optional" json:"insufficientDataActions" yaml:"insufficientDataActions"`
	// metrics property.
	//
	// Specify an array of string values to match this event if the actual value of metrics is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Metrics *[]*AlarmEvents_CloudWatchAlarmConfigurationChange_ConfigurationItem `field:"optional" json:"metrics" yaml:"metrics"`
	// okActions property.
	//
	// Specify an array of string values to match this event if the actual value of okActions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OkActions *[]*string `field:"optional" json:"okActions" yaml:"okActions"`
	// threshold property.
	//
	// Specify an array of string values to match this event if the actual value of threshold is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Threshold *[]*string `field:"optional" json:"threshold" yaml:"threshold"`
	// timestamp property.
	//
	// Specify an array of string values to match this event if the actual value of timestamp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Timestamp *[]*string `field:"optional" json:"timestamp" yaml:"timestamp"`
	// treatMissingData property.
	//
	// Specify an array of string values to match this event if the actual value of treatMissingData is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TreatMissingData *[]*string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
}

