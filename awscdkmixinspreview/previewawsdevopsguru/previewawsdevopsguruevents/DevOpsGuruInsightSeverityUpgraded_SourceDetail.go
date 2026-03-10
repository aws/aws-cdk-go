package previewawsdevopsguruevents


// Type definition for SourceDetail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceDetail := &SourceDetail{
//   	DataIdentifiers: &DataIdentifiers{
//   		AlarmCondition: &AlarmCondition{
//   			DetectionBand: []*string{
//   				jsii.String("detectionBand"),
//   			},
//   			ReferenceValue: &ReferenceValue{
//   				ComparisonOperator: []*string{
//   					jsii.String("comparisonOperator"),
//   				},
//   				DatapointsToAlarm: []*string{
//   					jsii.String("datapointsToAlarm"),
//   				},
//   				EvaluationPeriod: []*string{
//   					jsii.String("evaluationPeriod"),
//   				},
//   				Threshold: []*string{
//   					jsii.String("threshold"),
//   				},
//   			},
//   		},
//   		Dimensions: []CloudWatchDimension{
//   			&CloudWatchDimension{
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   		MetricDisplayName: []*string{
//   			jsii.String("metricDisplayName"),
//   		},
//   		MetricQuery: &MetricQuery{
//   			GroupBy: &GroupBy{
//   				Dimensions: []*string{
//   					jsii.String("dimensions"),
//   				},
//   				Group: []*string{
//   					jsii.String("group"),
//   				},
//   			},
//   			Metric: []*string{
//   				jsii.String("metric"),
//   			},
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		Namespace: []*string{
//   			jsii.String("namespace"),
//   		},
//   		Period: []*string{
//   			jsii.String("period"),
//   		},
//   		ResourceId: []*string{
//   			jsii.String("resourceId"),
//   		},
//   		ResourceType: []*string{
//   			jsii.String("resourceType"),
//   		},
//   		Stat: []*string{
//   			jsii.String("stat"),
//   		},
//   		Unit: []*string{
//   			jsii.String("unit"),
//   		},
//   	},
//   	DataSource: []*string{
//   		jsii.String("dataSource"),
//   	},
//   	DataSourceDetail: []*string{
//   		jsii.String("dataSourceDetail"),
//   	},
//   }
//
// Experimental.
type DevOpsGuruInsightSeverityUpgraded_SourceDetail struct {
	// dataIdentifiers property.
	//
	// Specify an array of string values to match this event if the actual value of dataIdentifiers is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DataIdentifiers *DevOpsGuruInsightSeverityUpgraded_DataIdentifiers `field:"optional" json:"dataIdentifiers" yaml:"dataIdentifiers"`
	// dataSource property.
	//
	// Specify an array of string values to match this event if the actual value of dataSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DataSource *[]*string `field:"optional" json:"dataSource" yaml:"dataSource"`
	// dataSourceDetail property.
	//
	// Specify an array of string values to match this event if the actual value of dataSourceDetail is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DataSourceDetail *[]*string `field:"optional" json:"dataSourceDetail" yaml:"dataSourceDetail"`
}

