package previewawsxrayevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.xray@AWSXRayInsightUpdate event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var names interface{}
//
//   aWSXRayInsightUpdateProps := &AWSXRayInsightUpdateProps{
//   	Categories: []*string{
//   		jsii.String("categories"),
//   	},
//   	ClientRequestImpactStatistics: &ClientRequestImpactStatistics{
//   		FaultCount: []*string{
//   			jsii.String("faultCount"),
//   		},
//   		OkCount: []*string{
//   			jsii.String("okCount"),
//   		},
//   		TotalCount: []*string{
//   			jsii.String("totalCount"),
//   		},
//   	},
//   	EndTime: []*string{
//   		jsii.String("endTime"),
//   	},
//   	Event: &Event{
//   		ClientRequestImpactStatistics: &ClientRequestImpactStatistics{
//   			FaultCount: []*string{
//   				jsii.String("faultCount"),
//   			},
//   			OkCount: []*string{
//   				jsii.String("okCount"),
//   			},
//   			TotalCount: []*string{
//   				jsii.String("totalCount"),
//   			},
//   		},
//   		EventTime: []*string{
//   			jsii.String("eventTime"),
//   		},
//   		RootCauseServiceRequestImpactStatistics: &ClientRequestImpactStatistics{
//   			FaultCount: []*string{
//   				jsii.String("faultCount"),
//   			},
//   			OkCount: []*string{
//   				jsii.String("okCount"),
//   			},
//   			TotalCount: []*string{
//   				jsii.String("totalCount"),
//   			},
//   		},
//   		Summary: []*string{
//   			jsii.String("summary"),
//   		},
//   		TopAnomalousServices: []AwsxRayInsightUpdateItem{
//   			&AwsxRayInsightUpdateItem{
//   				ServiceId: &ServiceId{
//   					AccountId: []*string{
//   						jsii.String("accountId"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					Names: []interface{}{
//   						names,
//   					},
//   					Type: []*string{
//   						jsii.String("type"),
//   					},
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
//   	GroupName: []*string{
//   		jsii.String("groupName"),
//   	},
//   	InsightId: []*string{
//   		jsii.String("insightId"),
//   	},
//   	RootCauseServiceId: &RootCauseServiceId{
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		Names: []interface{}{
//   			names,
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	RootCauseServiceRequestImpactStatistics: &ClientRequestImpactStatistics{
//   		FaultCount: []*string{
//   			jsii.String("faultCount"),
//   		},
//   		OkCount: []*string{
//   			jsii.String("okCount"),
//   		},
//   		TotalCount: []*string{
//   			jsii.String("totalCount"),
//   		},
//   	},
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   	Summary: []*string{
//   		jsii.String("summary"),
//   	},
//   	TopAnomalousServices: []AwsxRayInsightUpdateItem{
//   		&AwsxRayInsightUpdateItem{
//   			ServiceId: &ServiceId{
//   				AccountId: []*string{
//   					jsii.String("accountId"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Names: []interface{}{
//   					names,
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type AWSXRayInsightUpdate_AWSXRayInsightUpdateProps struct {
	// Categories property.
	//
	// Specify an array of string values to match this event if the actual value of Categories is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// ClientRequestImpactStatistics property.
	//
	// Specify an array of string values to match this event if the actual value of ClientRequestImpactStatistics is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientRequestImpactStatistics *AWSXRayInsightUpdate_ClientRequestImpactStatistics `field:"optional" json:"clientRequestImpactStatistics" yaml:"clientRequestImpactStatistics"`
	// EndTime property.
	//
	// Specify an array of string values to match this event if the actual value of EndTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTime *[]*string `field:"optional" json:"endTime" yaml:"endTime"`
	// Event property.
	//
	// Specify an array of string values to match this event if the actual value of Event is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Event *AWSXRayInsightUpdate_Event `field:"optional" json:"event" yaml:"event"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// GroupName property.
	//
	// Specify an array of string values to match this event if the actual value of GroupName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupName *[]*string `field:"optional" json:"groupName" yaml:"groupName"`
	// InsightId property.
	//
	// Specify an array of string values to match this event if the actual value of InsightId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InsightId *[]*string `field:"optional" json:"insightId" yaml:"insightId"`
	// RootCauseServiceId property.
	//
	// Specify an array of string values to match this event if the actual value of RootCauseServiceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RootCauseServiceId *AWSXRayInsightUpdate_RootCauseServiceId `field:"optional" json:"rootCauseServiceId" yaml:"rootCauseServiceId"`
	// RootCauseServiceRequestImpactStatistics property.
	//
	// Specify an array of string values to match this event if the actual value of RootCauseServiceRequestImpactStatistics is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RootCauseServiceRequestImpactStatistics *AWSXRayInsightUpdate_ClientRequestImpactStatistics `field:"optional" json:"rootCauseServiceRequestImpactStatistics" yaml:"rootCauseServiceRequestImpactStatistics"`
	// StartTime property.
	//
	// Specify an array of string values to match this event if the actual value of StartTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
	// State property.
	//
	// Specify an array of string values to match this event if the actual value of State is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
	// Summary property.
	//
	// Specify an array of string values to match this event if the actual value of Summary is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Summary *[]*string `field:"optional" json:"summary" yaml:"summary"`
	// TopAnomalousServices property.
	//
	// Specify an array of string values to match this event if the actual value of TopAnomalousServices is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TopAnomalousServices *[]*AWSXRayInsightUpdate_AwsxRayInsightUpdateItem `field:"optional" json:"topAnomalousServices" yaml:"topAnomalousServices"`
}

