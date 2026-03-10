package previewawssecurityhubevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.securityhub@SecurityHubInsightResults event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   securityHubInsightResultsProps := &SecurityHubInsightResultsProps{
//   	ActionDescription: []*string{
//   		jsii.String("actionDescription"),
//   	},
//   	ActionName: []*string{
//   		jsii.String("actionName"),
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
//   	InsightArn: []*string{
//   		jsii.String("insightArn"),
//   	},
//   	InsightName: []*string{
//   		jsii.String("insightName"),
//   	},
//   	InsightResults: []map[string]*string{
//   		map[string]*string{
//   			"insightResultsKey": jsii.String("insightResults"),
//   		},
//   	},
//   	ResultType: []*string{
//   		jsii.String("resultType"),
//   	},
//   }
//
// Experimental.
type SecurityHubInsightResults_SecurityHubInsightResultsProps struct {
	// actionDescription property.
	//
	// Specify an array of string values to match this event if the actual value of actionDescription is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActionDescription *[]*string `field:"optional" json:"actionDescription" yaml:"actionDescription"`
	// actionName property.
	//
	// Specify an array of string values to match this event if the actual value of actionName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActionName *[]*string `field:"optional" json:"actionName" yaml:"actionName"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// insightArn property.
	//
	// Specify an array of string values to match this event if the actual value of insightArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InsightArn *[]*string `field:"optional" json:"insightArn" yaml:"insightArn"`
	// insightName property.
	//
	// Specify an array of string values to match this event if the actual value of insightName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InsightName *[]*string `field:"optional" json:"insightName" yaml:"insightName"`
	// insightResults property.
	//
	// Specify an array of string values to match this event if the actual value of insightResults is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InsightResults *[]*map[string]*string `field:"optional" json:"insightResults" yaml:"insightResults"`
	// resultType property.
	//
	// Specify an array of string values to match this event if the actual value of resultType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResultType *[]*string `field:"optional" json:"resultType" yaml:"resultType"`
}

