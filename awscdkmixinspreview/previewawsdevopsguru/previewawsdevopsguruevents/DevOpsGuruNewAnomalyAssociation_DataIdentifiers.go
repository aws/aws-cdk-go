package previewawsdevopsguruevents


// Type definition for DataIdentifiers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataIdentifiers := &DataIdentifiers{
//   	AlarmCondition: &AlarmCondition{
//   		DetectionBand: []*string{
//   			jsii.String("detectionBand"),
//   		},
//   		ReferenceValue: &ReferenceValue{
//   			ComparisonOperator: []*string{
//   				jsii.String("comparisonOperator"),
//   			},
//   			DatapointsToAlarm: []*string{
//   				jsii.String("datapointsToAlarm"),
//   			},
//   			EvaluationPeriod: []*string{
//   				jsii.String("evaluationPeriod"),
//   			},
//   			Threshold: []*string{
//   				jsii.String("threshold"),
//   			},
//   		},
//   	},
//   	Dimensions: []CloudWatchDimension{
//   		&CloudWatchDimension{
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	MetricDisplayName: []*string{
//   		jsii.String("metricDisplayName"),
//   	},
//   	MetricQuery: &MetricQuery{
//   		GroupBy: &GroupBy{
//   			Dimensions: []*string{
//   				jsii.String("dimensions"),
//   			},
//   			Group: []*string{
//   				jsii.String("group"),
//   			},
//   		},
//   		Metric: []*string{
//   			jsii.String("metric"),
//   		},
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Namespace: []*string{
//   		jsii.String("namespace"),
//   	},
//   	Period: []*string{
//   		jsii.String("period"),
//   	},
//   	ResourceId: []*string{
//   		jsii.String("resourceId"),
//   	},
//   	ResourceType: []*string{
//   		jsii.String("resourceType"),
//   	},
//   	Stat: []*string{
//   		jsii.String("stat"),
//   	},
//   	Unit: []*string{
//   		jsii.String("unit"),
//   	},
//   }
//
// Experimental.
type DevOpsGuruNewAnomalyAssociation_DataIdentifiers struct {
	// alarmCondition property.
	//
	// Specify an array of string values to match this event if the actual value of alarmCondition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlarmCondition *DevOpsGuruNewAnomalyAssociation_AlarmCondition `field:"optional" json:"alarmCondition" yaml:"alarmCondition"`
	// dimensions property.
	//
	// Specify an array of string values to match this event if the actual value of dimensions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Dimensions *[]*DevOpsGuruNewAnomalyAssociation_CloudWatchDimension `field:"optional" json:"dimensions" yaml:"dimensions"`
	// metricDisplayName property.
	//
	// Specify an array of string values to match this event if the actual value of metricDisplayName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MetricDisplayName *[]*string `field:"optional" json:"metricDisplayName" yaml:"metricDisplayName"`
	// metricQuery property.
	//
	// Specify an array of string values to match this event if the actual value of metricQuery is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MetricQuery *DevOpsGuruNewAnomalyAssociation_MetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// namespace property.
	//
	// Specify an array of string values to match this event if the actual value of namespace is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Namespace *[]*string `field:"optional" json:"namespace" yaml:"namespace"`
	// period property.
	//
	// Specify an array of string values to match this event if the actual value of period is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Period *[]*string `field:"optional" json:"period" yaml:"period"`
	// ResourceId property.
	//
	// Specify an array of string values to match this event if the actual value of ResourceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceId *[]*string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// ResourceType property.
	//
	// Specify an array of string values to match this event if the actual value of ResourceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceType *[]*string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// stat property.
	//
	// Specify an array of string values to match this event if the actual value of stat is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Stat *[]*string `field:"optional" json:"stat" yaml:"stat"`
	// unit property.
	//
	// Specify an array of string values to match this event if the actual value of unit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Unit *[]*string `field:"optional" json:"unit" yaml:"unit"`
}

