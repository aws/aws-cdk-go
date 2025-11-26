package previewawscloudwatchevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Alarm aws.cloudwatch@CloudWatchAlarmConfigurationChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchAlarmConfigurationChangeProps := &CloudWatchAlarmConfigurationChangeProps{
//   	AlarmName: []*string{
//   		jsii.String("alarmName"),
//   	},
//   	Configuration: &Configuration{
//   		ActionsEnabled: []*string{
//   			jsii.String("actionsEnabled"),
//   		},
//   		ActionsSuppressor: []*string{
//   			jsii.String("actionsSuppressor"),
//   		},
//   		ActionsSuppressorExtensionPeriod: []*string{
//   			jsii.String("actionsSuppressorExtensionPeriod"),
//   		},
//   		ActionsSuppressorWaitPeriod: []*string{
//   			jsii.String("actionsSuppressorWaitPeriod"),
//   		},
//   		AlarmActions: []*string{
//   			jsii.String("alarmActions"),
//   		},
//   		AlarmName: []*string{
//   			jsii.String("alarmName"),
//   		},
//   		AlarmRule: []*string{
//   			jsii.String("alarmRule"),
//   		},
//   		ComparisonOperator: []*string{
//   			jsii.String("comparisonOperator"),
//   		},
//   		DatapointsToAlarm: []*string{
//   			jsii.String("datapointsToAlarm"),
//   		},
//   		Description: []*string{
//   			jsii.String("description"),
//   		},
//   		EvaluateLowSampleCountPercentile: []*string{
//   			jsii.String("evaluateLowSampleCountPercentile"),
//   		},
//   		EvaluationPeriods: []*string{
//   			jsii.String("evaluationPeriods"),
//   		},
//   		InsufficientDataActions: []*string{
//   			jsii.String("insufficientDataActions"),
//   		},
//   		Metrics: []ConfigurationItem{
//   			&ConfigurationItem{
//   				Id: []*string{
//   					jsii.String("id"),
//   				},
//   				MetricStat: &MetricStat{
//   					Metric: &Metric{
//   						Dimensions: []*string{
//   							jsii.String("dimensions"),
//   						},
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   						Namespace: []*string{
//   							jsii.String("namespace"),
//   						},
//   					},
//   					Period: []*string{
//   						jsii.String("period"),
//   					},
//   					Stat: []*string{
//   						jsii.String("stat"),
//   					},
//   				},
//   				ReturnData: []*string{
//   					jsii.String("returnData"),
//   				},
//   			},
//   		},
//   		OkActions: []*string{
//   			jsii.String("okActions"),
//   		},
//   		Threshold: []*string{
//   			jsii.String("threshold"),
//   		},
//   		Timestamp: []*string{
//   			jsii.String("timestamp"),
//   		},
//   		TreatMissingData: []*string{
//   			jsii.String("treatMissingData"),
//   		},
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	Operation: []*string{
//   		jsii.String("operation"),
//   	},
//   	State: &State{
//   		ActionsSuppressedBy: []*string{
//   			jsii.String("actionsSuppressedBy"),
//   		},
//   		EvaluationState: []*string{
//   			jsii.String("evaluationState"),
//   		},
//   		Reason: []*string{
//   			jsii.String("reason"),
//   		},
//   		ReasonData: []*string{
//   			jsii.String("reasonData"),
//   		},
//   		Timestamp: []*string{
//   			jsii.String("timestamp"),
//   		},
//   		Value: []*string{
//   			jsii.String("value"),
//   		},
//   	},
//   }
//
// Experimental.
type AlarmEvents_CloudWatchAlarmConfigurationChange_CloudWatchAlarmConfigurationChangeProps struct {
	// alarmName property.
	//
	// Specify an array of string values to match this event if the actual value of alarmName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlarmName *[]*string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// configuration property.
	//
	// Specify an array of string values to match this event if the actual value of configuration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Configuration *AlarmEvents_CloudWatchAlarmConfigurationChange_Configuration `field:"optional" json:"configuration" yaml:"configuration"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// operation property.
	//
	// Specify an array of string values to match this event if the actual value of operation is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Operation *[]*string `field:"optional" json:"operation" yaml:"operation"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *AlarmEvents_CloudWatchAlarmConfigurationChange_State `field:"optional" json:"state" yaml:"state"`
}

