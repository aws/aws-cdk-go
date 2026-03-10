package previewawsdevopsguruevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.devopsguru@DevOpsGuruNewRecommendationCreated event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   devOpsGuruNewRecommendationCreatedProps := &DevOpsGuruNewRecommendationCreatedProps{
//   	AccountId: []*string{
//   		jsii.String("accountId"),
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
//   	InsightDescription: []*string{
//   		jsii.String("insightDescription"),
//   	},
//   	InsightId: []*string{
//   		jsii.String("insightId"),
//   	},
//   	InsightName: []*string{
//   		jsii.String("insightName"),
//   	},
//   	InsightType: []*string{
//   		jsii.String("insightType"),
//   	},
//   	InsightUrl: []*string{
//   		jsii.String("insightUrl"),
//   	},
//   	MessageType: []*string{
//   		jsii.String("messageType"),
//   	},
//   	Recommendations: []Recommendation{
//   		&Recommendation{
//   			Description: []*string{
//   				jsii.String("description"),
//   			},
//   			Link: []*string{
//   				jsii.String("link"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Reason: []*string{
//   				jsii.String("reason"),
//   			},
//   			RelatedAnomalies: []RelatedAnomaly{
//   				&RelatedAnomaly{
//   					AssociatedResourceArns: []*string{
//   						jsii.String("associatedResourceArns"),
//   					},
//   					Resources: []AnomalyResource{
//   						&AnomalyResource{
//   							Name: []*string{
//   								jsii.String("name"),
//   							},
//   							Type: []*string{
//   								jsii.String("type"),
//   							},
//   						},
//   					},
//   					SourceDetails: []RelatedAnomalySourceDetail{
//   						&RelatedAnomalySourceDetail{
//   							CloudWatchMetrics: []CloudWatchDimension{
//   								&CloudWatchDimension{
//   									MetricName: []*string{
//   										jsii.String("metricName"),
//   									},
//   									Namespace: []*string{
//   										jsii.String("namespace"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   			RelatedEvents: []RelatedEvent{
//   				&RelatedEvent{
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					Resources: []EventResource{
//   						&EventResource{
//   							Name: []*string{
//   								jsii.String("name"),
//   							},
//   							Type: []*string{
//   								jsii.String("type"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Region: []*string{
//   		jsii.String("region"),
//   	},
//   	ResourceCollection: &ResourceCollection{
//   		CloudFormation: &CloudFormation{
//   			StackNames: []*string{
//   				jsii.String("stackNames"),
//   			},
//   		},
//   		Tags: []TagType{
//   			&TagType{
//   				AppBoundaryKey: []*string{
//   					jsii.String("appBoundaryKey"),
//   				},
//   				TagValues: []*string{
//   					jsii.String("tagValues"),
//   				},
//   			},
//   		},
//   	},
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   	StartTimeIso: []*string{
//   		jsii.String("startTimeIso"),
//   	},
//   }
//
// Experimental.
type DevOpsGuruNewRecommendationCreated_DevOpsGuruNewRecommendationCreatedProps struct {
	// accountId property.
	//
	// Specify an array of string values to match this event if the actual value of accountId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// insightDescription property.
	//
	// Specify an array of string values to match this event if the actual value of insightDescription is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InsightDescription *[]*string `field:"optional" json:"insightDescription" yaml:"insightDescription"`
	// insightId property.
	//
	// Specify an array of string values to match this event if the actual value of insightId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InsightId *[]*string `field:"optional" json:"insightId" yaml:"insightId"`
	// insightName property.
	//
	// Specify an array of string values to match this event if the actual value of insightName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InsightName *[]*string `field:"optional" json:"insightName" yaml:"insightName"`
	// insightType property.
	//
	// Specify an array of string values to match this event if the actual value of insightType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InsightType *[]*string `field:"optional" json:"insightType" yaml:"insightType"`
	// insightUrl property.
	//
	// Specify an array of string values to match this event if the actual value of insightUrl is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InsightUrl *[]*string `field:"optional" json:"insightUrl" yaml:"insightUrl"`
	// messageType property.
	//
	// Specify an array of string values to match this event if the actual value of messageType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MessageType *[]*string `field:"optional" json:"messageType" yaml:"messageType"`
	// recommendations property.
	//
	// Specify an array of string values to match this event if the actual value of recommendations is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Recommendations *[]*DevOpsGuruNewRecommendationCreated_Recommendation `field:"optional" json:"recommendations" yaml:"recommendations"`
	// region property.
	//
	// Specify an array of string values to match this event if the actual value of region is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Region *[]*string `field:"optional" json:"region" yaml:"region"`
	// resourceCollection property.
	//
	// Specify an array of string values to match this event if the actual value of resourceCollection is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceCollection *DevOpsGuruNewRecommendationCreated_ResourceCollection `field:"optional" json:"resourceCollection" yaml:"resourceCollection"`
	// startTime property.
	//
	// Specify an array of string values to match this event if the actual value of startTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
	// startTimeISO property.
	//
	// Specify an array of string values to match this event if the actual value of startTimeISO is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTimeIso *[]*string `field:"optional" json:"startTimeIso" yaml:"startTimeIso"`
}

