package previewawsdevopsguruevents


// Type definition for Anomaly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   anomaly := &Anomaly{
//   	AnomalyResources: []AnomalyResource{
//   		&AnomalyResource{
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   	},
//   	AnomalySeverity: []*string{
//   		jsii.String("anomalySeverity"),
//   	},
//   	AssociatedResourceArns: []*string{
//   		jsii.String("associatedResourceArns"),
//   	},
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	EarliestImpactTime: []*string{
//   		jsii.String("earliestImpactTime"),
//   	},
//   	EarliestImpactTimeIso: []*string{
//   		jsii.String("earliestImpactTimeIso"),
//   	},
//   	EndTime: []*string{
//   		jsii.String("endTime"),
//   	},
//   	EndTimeIso: []*string{
//   		jsii.String("endTimeIso"),
//   	},
//   	Id: []*string{
//   		jsii.String("id"),
//   	},
//   	LatestImpactTime: []*string{
//   		jsii.String("latestImpactTime"),
//   	},
//   	LatestImpactTimeIso: []*string{
//   		jsii.String("latestImpactTimeIso"),
//   	},
//   	Limit: []*string{
//   		jsii.String("limit"),
//   	},
//   	OpenTime: []*string{
//   		jsii.String("openTime"),
//   	},
//   	OpenTimeIso: []*string{
//   		jsii.String("openTimeIso"),
//   	},
//   	SourceDetails: []SourceDetail{
//   		&SourceDetail{
//   			DataIdentifiers: &DataIdentifiers{
//   				AlarmCondition: &AlarmCondition{
//   					DetectionBand: []*string{
//   						jsii.String("detectionBand"),
//   					},
//   					ReferenceValue: &ReferenceValue{
//   						ComparisonOperator: []*string{
//   							jsii.String("comparisonOperator"),
//   						},
//   						DatapointsToAlarm: []*string{
//   							jsii.String("datapointsToAlarm"),
//   						},
//   						EvaluationPeriod: []*string{
//   							jsii.String("evaluationPeriod"),
//   						},
//   						Threshold: []*string{
//   							jsii.String("threshold"),
//   						},
//   					},
//   				},
//   				Dimensions: []CloudWatchDimension{
//   					&CloudWatchDimension{
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   						Value: []*string{
//   							jsii.String("value"),
//   						},
//   					},
//   				},
//   				MetricDisplayName: []*string{
//   					jsii.String("metricDisplayName"),
//   				},
//   				MetricQuery: &MetricQuery{
//   					GroupBy: &GroupBy{
//   						Dimensions: []*string{
//   							jsii.String("dimensions"),
//   						},
//   						Group: []*string{
//   							jsii.String("group"),
//   						},
//   					},
//   					Metric: []*string{
//   						jsii.String("metric"),
//   					},
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Namespace: []*string{
//   					jsii.String("namespace"),
//   				},
//   				Period: []*string{
//   					jsii.String("period"),
//   				},
//   				ResourceId: []*string{
//   					jsii.String("resourceId"),
//   				},
//   				ResourceType: []*string{
//   					jsii.String("resourceType"),
//   				},
//   				Stat: []*string{
//   					jsii.String("stat"),
//   				},
//   				Unit: []*string{
//   					jsii.String("unit"),
//   				},
//   			},
//   			DataSource: []*string{
//   				jsii.String("dataSource"),
//   			},
//   			DataSourceDetail: []*string{
//   				jsii.String("dataSourceDetail"),
//   			},
//   		},
//   	},
//   	SourceMetadata: &AnomalySourceMetadata{
//   		Source: []*string{
//   			jsii.String("source"),
//   		},
//   		SourceResourceArn: []*string{
//   			jsii.String("sourceResourceArn"),
//   		},
//   		SourceResourceType: []*string{
//   			jsii.String("sourceResourceType"),
//   		},
//   	},
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   	StartTimeIso: []*string{
//   		jsii.String("startTimeIso"),
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type DevOpsGuruNewAnomalyAssociation_Anomaly struct {
	// anomalyResources property.
	//
	// Specify an array of string values to match this event if the actual value of anomalyResources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AnomalyResources *[]*DevOpsGuruNewAnomalyAssociation_AnomalyResource `field:"optional" json:"anomalyResources" yaml:"anomalyResources"`
	// anomalySeverity property.
	//
	// Specify an array of string values to match this event if the actual value of anomalySeverity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AnomalySeverity *[]*string `field:"optional" json:"anomalySeverity" yaml:"anomalySeverity"`
	// associatedResourceArns property.
	//
	// Specify an array of string values to match this event if the actual value of associatedResourceArns is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AssociatedResourceArns *[]*string `field:"optional" json:"associatedResourceArns" yaml:"associatedResourceArns"`
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// earliestImpactTime property.
	//
	// Specify an array of string values to match this event if the actual value of earliestImpactTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EarliestImpactTime *[]*string `field:"optional" json:"earliestImpactTime" yaml:"earliestImpactTime"`
	// earliestImpactTimeISO property.
	//
	// Specify an array of string values to match this event if the actual value of earliestImpactTimeISO is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EarliestImpactTimeIso *[]*string `field:"optional" json:"earliestImpactTimeIso" yaml:"earliestImpactTimeIso"`
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
	// id property.
	//
	// Specify an array of string values to match this event if the actual value of id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
	// latestImpactTime property.
	//
	// Specify an array of string values to match this event if the actual value of latestImpactTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LatestImpactTime *[]*string `field:"optional" json:"latestImpactTime" yaml:"latestImpactTime"`
	// latestImpactTimeISO property.
	//
	// Specify an array of string values to match this event if the actual value of latestImpactTimeISO is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LatestImpactTimeIso *[]*string `field:"optional" json:"latestImpactTimeIso" yaml:"latestImpactTimeIso"`
	// limit property.
	//
	// Specify an array of string values to match this event if the actual value of limit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Limit *[]*string `field:"optional" json:"limit" yaml:"limit"`
	// openTime property.
	//
	// Specify an array of string values to match this event if the actual value of openTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OpenTime *[]*string `field:"optional" json:"openTime" yaml:"openTime"`
	// openTimeISO property.
	//
	// Specify an array of string values to match this event if the actual value of openTimeISO is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OpenTimeIso *[]*string `field:"optional" json:"openTimeIso" yaml:"openTimeIso"`
	// sourceDetails property.
	//
	// Specify an array of string values to match this event if the actual value of sourceDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceDetails *[]*DevOpsGuruNewAnomalyAssociation_SourceDetail `field:"optional" json:"sourceDetails" yaml:"sourceDetails"`
	// sourceMetadata property.
	//
	// Specify an array of string values to match this event if the actual value of sourceMetadata is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceMetadata *DevOpsGuruNewAnomalyAssociation_AnomalySourceMetadata `field:"optional" json:"sourceMetadata" yaml:"sourceMetadata"`
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
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

