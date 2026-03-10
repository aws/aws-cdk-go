package previewawsdevopsguruevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.devopsguru@DevOpsGuruInsightClosed event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   devOpsGuruInsightClosedProps := &DevOpsGuruInsightClosedProps{
//   	AccountId: []*string{
//   		jsii.String("accountId"),
//   	},
//   	Anomalies: []Anomaly{
//   		&Anomaly{
//   			AnomalyResources: []AnomalyResource{
//   				&AnomalyResource{
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					Type: []*string{
//   						jsii.String("type"),
//   					},
//   				},
//   			},
//   			AnomalySeverity: []*string{
//   				jsii.String("anomalySeverity"),
//   			},
//   			AssociatedResourceArns: []*string{
//   				jsii.String("associatedResourceArns"),
//   			},
//   			CloseTime: []*string{
//   				jsii.String("closeTime"),
//   			},
//   			CloseTimeIso: []*string{
//   				jsii.String("closeTimeIso"),
//   			},
//   			Description: []*string{
//   				jsii.String("description"),
//   			},
//   			EarliestImpactTime: []*string{
//   				jsii.String("earliestImpactTime"),
//   			},
//   			EarliestImpactTimeIso: []*string{
//   				jsii.String("earliestImpactTimeIso"),
//   			},
//   			EndTime: []*string{
//   				jsii.String("endTime"),
//   			},
//   			EndTimeIso: []*string{
//   				jsii.String("endTimeIso"),
//   			},
//   			Id: []*string{
//   				jsii.String("id"),
//   			},
//   			LatestImpactTime: []*string{
//   				jsii.String("latestImpactTime"),
//   			},
//   			LatestImpactTimeIso: []*string{
//   				jsii.String("latestImpactTimeIso"),
//   			},
//   			Limit: []*string{
//   				jsii.String("limit"),
//   			},
//   			OpenTime: []*string{
//   				jsii.String("openTime"),
//   			},
//   			OpenTimeIso: []*string{
//   				jsii.String("openTimeIso"),
//   			},
//   			SourceDetails: []SourceDetail{
//   				&SourceDetail{
//   					DataIdentifiers: &DataIdentifiers{
//   						AlarmCondition: &AlarmCondition{
//   							DetectionBand: []*string{
//   								jsii.String("detectionBand"),
//   							},
//   							ReferenceValue: &ReferenceValue{
//   								ComparisonOperator: []*string{
//   									jsii.String("comparisonOperator"),
//   								},
//   								DatapointsToAlarm: []*string{
//   									jsii.String("datapointsToAlarm"),
//   								},
//   								EvaluationPeriod: []*string{
//   									jsii.String("evaluationPeriod"),
//   								},
//   								Threshold: []*string{
//   									jsii.String("threshold"),
//   								},
//   							},
//   						},
//   						Dimensions: []CloudWatchDimension{
//   							&CloudWatchDimension{
//   								Name: []*string{
//   									jsii.String("name"),
//   								},
//   								Value: []*string{
//   									jsii.String("value"),
//   								},
//   							},
//   						},
//   						MetricDisplayName: []*string{
//   							jsii.String("metricDisplayName"),
//   						},
//   						MetricQuery: &MetricQuery{
//   							GroupBy: &GroupBy{
//   								Dimensions: []*string{
//   									jsii.String("dimensions"),
//   								},
//   								Group: []*string{
//   									jsii.String("group"),
//   								},
//   							},
//   							Metric: []*string{
//   								jsii.String("metric"),
//   							},
//   						},
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   						Namespace: []*string{
//   							jsii.String("namespace"),
//   						},
//   						Period: []*string{
//   							jsii.String("period"),
//   						},
//   						ResourceId: []*string{
//   							jsii.String("resourceId"),
//   						},
//   						ResourceType: []*string{
//   							jsii.String("resourceType"),
//   						},
//   						Stat: []*string{
//   							jsii.String("stat"),
//   						},
//   						Unit: []*string{
//   							jsii.String("unit"),
//   						},
//   					},
//   					DataSource: []*string{
//   						jsii.String("dataSource"),
//   					},
//   					DataSourceDetail: []*string{
//   						jsii.String("dataSourceDetail"),
//   					},
//   				},
//   			},
//   			SourceMetadata: &AnomalySourceMetadata{
//   				Source: []*string{
//   					jsii.String("source"),
//   				},
//   				SourceResourceArn: []*string{
//   					jsii.String("sourceResourceArn"),
//   				},
//   				SourceResourceType: []*string{
//   					jsii.String("sourceResourceType"),
//   				},
//   			},
//   			StartTime: []*string{
//   				jsii.String("startTime"),
//   			},
//   			StartTimeIso: []*string{
//   				jsii.String("startTimeIso"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   	},
//   	EndTime: []*string{
//   		jsii.String("endTime"),
//   	},
//   	EndTimeIso: []*string{
//   		jsii.String("endTimeIso"),
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
//   	InsightSeverity: []*string{
//   		jsii.String("insightSeverity"),
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
type DevOpsGuruInsightClosed_DevOpsGuruInsightClosedProps struct {
	// accountId property.
	//
	// Specify an array of string values to match this event if the actual value of accountId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// anomalies property.
	//
	// Specify an array of string values to match this event if the actual value of anomalies is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Anomalies *[]*DevOpsGuruInsightClosed_Anomaly `field:"optional" json:"anomalies" yaml:"anomalies"`
	// endTime property.
	//
	// Specify an array of string values to match this event if the actual value of endTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTime *[]*string `field:"optional" json:"endTime" yaml:"endTime"`
	// endTimeISO property.
	//
	// Specify an array of string values to match this event if the actual value of endTimeISO is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTimeIso *[]*string `field:"optional" json:"endTimeIso" yaml:"endTimeIso"`
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
	// insightSeverity property.
	//
	// Specify an array of string values to match this event if the actual value of insightSeverity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InsightSeverity *[]*string `field:"optional" json:"insightSeverity" yaml:"insightSeverity"`
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
	ResourceCollection *DevOpsGuruInsightClosed_ResourceCollection `field:"optional" json:"resourceCollection" yaml:"resourceCollection"`
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

