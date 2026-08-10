package previewawstranscribeevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.transcribe@CallAnalyticsJobStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   callAnalyticsJobStateChangeProps := &CallAnalyticsJobStateChangeProps{
//   	AnalyticsJobDetails: &AnalyticsJobDetails{
//   		Skipped: []CallAnalyticsSkippedFeature{
//   			&CallAnalyticsSkippedFeature{
//   				Feature: []*string{
//   					jsii.String("feature"),
//   				},
//   				Message: []*string{
//   					jsii.String("message"),
//   				},
//   				ReasonCode: []*string{
//   					jsii.String("reasonCode"),
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
//   	JobName: []*string{
//   		jsii.String("jobName"),
//   	},
//   	JobStatus: []*string{
//   		jsii.String("jobStatus"),
//   	},
//   }
//
// Experimental.
type CallAnalyticsJobStateChange_CallAnalyticsJobStateChangeProps struct {
	// AnalyticsJobDetails property.
	//
	// Specify an array of string values to match this event if the actual value of AnalyticsJobDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AnalyticsJobDetails *CallAnalyticsJobStateChange_AnalyticsJobDetails `field:"optional" json:"analyticsJobDetails" yaml:"analyticsJobDetails"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// JobName property.
	//
	// Specify an array of string values to match this event if the actual value of JobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobName *[]*string `field:"optional" json:"jobName" yaml:"jobName"`
	// JobStatus property.
	//
	// Specify an array of string values to match this event if the actual value of JobStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobStatus *[]*string `field:"optional" json:"jobStatus" yaml:"jobStatus"`
}

