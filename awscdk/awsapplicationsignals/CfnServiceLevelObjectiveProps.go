package awsapplicationsignals

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnServiceLevelObjective`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceLevelObjectiveProps := &CfnServiceLevelObjectiveProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Goal: &GoalProperty{
//   		AttainmentGoal: jsii.Number(123),
//   		Interval: &IntervalProperty{
//   			CalendarInterval: &CalendarIntervalProperty{
//   				Duration: jsii.Number(123),
//   				DurationUnit: jsii.String("durationUnit"),
//   				StartTime: jsii.Number(123),
//   			},
//   			RollingInterval: &RollingIntervalProperty{
//   				Duration: jsii.Number(123),
//   				DurationUnit: jsii.String("durationUnit"),
//   			},
//   		},
//   		WarningThreshold: jsii.Number(123),
//   	},
//   	RequestBasedSli: &RequestBasedSliProperty{
//   		RequestBasedSliMetric: &RequestBasedSliMetricProperty{
//   			KeyAttributes: map[string]*string{
//   				"keyAttributesKey": jsii.String("keyAttributes"),
//   			},
//   			MetricType: jsii.String("metricType"),
//   			MonitoredRequestCountMetric: &MonitoredRequestCountMetricProperty{
//   				BadCountMetric: []interface{}{
//   					&MetricDataQueryProperty{
//   						Id: jsii.String("id"),
//
//   						// the properties below are optional
//   						AccountId: jsii.String("accountId"),
//   						Expression: jsii.String("expression"),
//   						MetricStat: &MetricStatProperty{
//   							Metric: &MetricProperty{
//   								Dimensions: []interface{}{
//   									&DimensionProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//   							},
//   							Period: jsii.Number(123),
//   							Stat: jsii.String("stat"),
//
//   							// the properties below are optional
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   				GoodCountMetric: []interface{}{
//   					&MetricDataQueryProperty{
//   						Id: jsii.String("id"),
//
//   						// the properties below are optional
//   						AccountId: jsii.String("accountId"),
//   						Expression: jsii.String("expression"),
//   						MetricStat: &MetricStatProperty{
//   							Metric: &MetricProperty{
//   								Dimensions: []interface{}{
//   									&DimensionProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//   							},
//   							Period: jsii.Number(123),
//   							Stat: jsii.String("stat"),
//
//   							// the properties below are optional
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			OperationName: jsii.String("operationName"),
//   			TotalRequestCountMetric: []interface{}{
//   				&MetricDataQueryProperty{
//   					Id: jsii.String("id"),
//
//   					// the properties below are optional
//   					AccountId: jsii.String("accountId"),
//   					Expression: jsii.String("expression"),
//   					MetricStat: &MetricStatProperty{
//   						Metric: &MetricProperty{
//   							Dimensions: []interface{}{
//   								&DimensionProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MetricName: jsii.String("metricName"),
//   							Namespace: jsii.String("namespace"),
//   						},
//   						Period: jsii.Number(123),
//   						Stat: jsii.String("stat"),
//
//   						// the properties below are optional
//   						Unit: jsii.String("unit"),
//   					},
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		MetricThreshold: jsii.Number(123),
//   	},
//   	Sli: &SliProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		MetricThreshold: jsii.Number(123),
//   		SliMetric: &SliMetricProperty{
//   			KeyAttributes: map[string]*string{
//   				"keyAttributesKey": jsii.String("keyAttributes"),
//   			},
//   			MetricDataQueries: []interface{}{
//   				&MetricDataQueryProperty{
//   					Id: jsii.String("id"),
//
//   					// the properties below are optional
//   					AccountId: jsii.String("accountId"),
//   					Expression: jsii.String("expression"),
//   					MetricStat: &MetricStatProperty{
//   						Metric: &MetricProperty{
//   							Dimensions: []interface{}{
//   								&DimensionProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MetricName: jsii.String("metricName"),
//   							Namespace: jsii.String("namespace"),
//   						},
//   						Period: jsii.Number(123),
//   						Stat: jsii.String("stat"),
//
//   						// the properties below are optional
//   						Unit: jsii.String("unit"),
//   					},
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   			MetricType: jsii.String("metricType"),
//   			OperationName: jsii.String("operationName"),
//   			PeriodSeconds: jsii.Number(123),
//   			Statistic: jsii.String("statistic"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html
//
type CfnServiceLevelObjectiveProps struct {
	// A name for this SLO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional description for this SLO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-description
	//
	// Default: - "No description".
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// This structure contains the attributes that determine the goal of an SLO.
	//
	// This includes the time period for evaluation and the attainment threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-goal
	//
	Goal interface{} `field:"optional" json:"goal" yaml:"goal"`
	// A structure containing information about the performance metric that this SLO monitors, if this is a request-based SLO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-requestbasedsli
	//
	RequestBasedSli interface{} `field:"optional" json:"requestBasedSli" yaml:"requestBasedSli"`
	// A structure containing information about the performance metric that this SLO monitors, if this is a period-based SLO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-sli
	//
	Sli interface{} `field:"optional" json:"sli" yaml:"sli"`
	// A list of key-value pairs to associate with the SLO.
	//
	// You can associate as many as 50 tags with an SLO. To be able to associate tags with the SLO when you create the SLO, you must have the cloudwatch:TagResource permission.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

