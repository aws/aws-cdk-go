package previewawscloudwatchevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Alarm aws.cloudwatch@CloudWatchAlarmStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchAlarmStateChangeProps := &CloudWatchAlarmStateChangeProps{
//   	AlarmName: []*string{
//   		jsii.String("alarmName"),
//   	},
//   	Configuration: &Configuration{
//   		ActionsSuppressor: []*string{
//   			jsii.String("actionsSuppressor"),
//   		},
//   		ActionsSuppressorExtensionPeriod: []*string{
//   			jsii.String("actionsSuppressorExtensionPeriod"),
//   		},
//   		ActionsSuppressorWaitPeriod: []*string{
//   			jsii.String("actionsSuppressorWaitPeriod"),
//   		},
//   		AlarmRule: []*string{
//   			jsii.String("alarmRule"),
//   		},
//   		Description: []*string{
//   			jsii.String("description"),
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
//   	PreviousState: &State{
//   		ActionsSuppressedBy: []*string{
//   			jsii.String("actionsSuppressedBy"),
//   		},
//   		ActionsSuppressedReason: []*string{
//   			jsii.String("actionsSuppressedReason"),
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
//   	State: &State{
//   		ActionsSuppressedBy: []*string{
//   			jsii.String("actionsSuppressedBy"),
//   		},
//   		ActionsSuppressedReason: []*string{
//   			jsii.String("actionsSuppressedReason"),
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
type AlarmEvents_CloudWatchAlarmStateChange_CloudWatchAlarmStateChangeProps struct {
	// alarmName property.
	//
	// Specify an array of string values to match this event if the actual value of alarmName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Alarm reference.
	//
	// Experimental.
	AlarmName *[]*string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// configuration property.
	//
	// Specify an array of string values to match this event if the actual value of configuration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Configuration *AlarmEvents_CloudWatchAlarmStateChange_Configuration `field:"optional" json:"configuration" yaml:"configuration"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// previousState property.
	//
	// Specify an array of string values to match this event if the actual value of previousState is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousState *AlarmEvents_CloudWatchAlarmStateChange_State `field:"optional" json:"previousState" yaml:"previousState"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *AlarmEvents_CloudWatchAlarmStateChange_State `field:"optional" json:"state" yaml:"state"`
}

