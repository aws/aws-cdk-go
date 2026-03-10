package previewawscodeguruprofilerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.codeguruprofiler@CodeGuruProfilerRecommendationStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeGuruProfilerRecommendationStateChangeProps := &CodeGuruProfilerRecommendationStateChangeProps{
//   	ComputeInstanceArns: []*string{
//   		jsii.String("computeInstanceArns"),
//   	},
//   	DeduplicationId: []*string{
//   		jsii.String("deduplicationId"),
//   	},
//   	EventEndTime: []*string{
//   		jsii.String("eventEndTime"),
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
//   	EventStartTime: []*string{
//   		jsii.String("eventStartTime"),
//   	},
//   	ExpiresOn: []*string{
//   		jsii.String("expiresOn"),
//   	},
//   	Recommendation: &Recommendation{
//   		Description: &Description{
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   		Name: &Name{
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   		Reason: &Reason{
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   		ResolutionSteps: &ResolutionSteps{
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	Schema: []*string{
//   		jsii.String("schema"),
//   	},
//   	Severity: []*string{
//   		jsii.String("severity"),
//   	},
//   	SourceUrl: []*string{
//   		jsii.String("sourceUrl"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	Title: &Title{
//   		Value: []*string{
//   			jsii.String("value"),
//   		},
//   	},
//   }
//
// Experimental.
type CodeGuruProfilerRecommendationStateChange_CodeGuruProfilerRecommendationStateChangeProps struct {
	// computeInstanceArns property.
	//
	// Specify an array of string values to match this event if the actual value of computeInstanceArns is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ComputeInstanceArns *[]*string `field:"optional" json:"computeInstanceArns" yaml:"computeInstanceArns"`
	// deduplicationId property.
	//
	// Specify an array of string values to match this event if the actual value of deduplicationId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeduplicationId *[]*string `field:"optional" json:"deduplicationId" yaml:"deduplicationId"`
	// eventEndTime property.
	//
	// Specify an array of string values to match this event if the actual value of eventEndTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventEndTime *[]*string `field:"optional" json:"eventEndTime" yaml:"eventEndTime"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// eventStartTime property.
	//
	// Specify an array of string values to match this event if the actual value of eventStartTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventStartTime *[]*string `field:"optional" json:"eventStartTime" yaml:"eventStartTime"`
	// expiresOn property.
	//
	// Specify an array of string values to match this event if the actual value of expiresOn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExpiresOn *[]*string `field:"optional" json:"expiresOn" yaml:"expiresOn"`
	// recommendation property.
	//
	// Specify an array of string values to match this event if the actual value of recommendation is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Recommendation *CodeGuruProfilerRecommendationStateChange_Recommendation `field:"optional" json:"recommendation" yaml:"recommendation"`
	// schema property.
	//
	// Specify an array of string values to match this event if the actual value of schema is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Schema *[]*string `field:"optional" json:"schema" yaml:"schema"`
	// severity property.
	//
	// Specify an array of string values to match this event if the actual value of severity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Severity *[]*string `field:"optional" json:"severity" yaml:"severity"`
	// sourceUrl property.
	//
	// Specify an array of string values to match this event if the actual value of sourceUrl is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceUrl *[]*string `field:"optional" json:"sourceUrl" yaml:"sourceUrl"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// title property.
	//
	// Specify an array of string values to match this event if the actual value of title is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Title *CodeGuruProfilerRecommendationStateChange_Title `field:"optional" json:"title" yaml:"title"`
}

