package mixinsawsapplicationsignals

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnServiceLevelObjectivePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceLevelObjectiveMixinProps := &CfnServiceLevelObjectiveMixinProps{
//   	BurnRateConfigurations: []interface{}{
//   		&BurnRateConfigurationProperty{
//   			LookBackWindowMinutes: jsii.Number(123),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ExclusionWindows: []interface{}{
//   		&ExclusionWindowProperty{
//   			Reason: jsii.String("reason"),
//   			RecurrenceRule: &RecurrenceRuleProperty{
//   				Expression: jsii.String("expression"),
//   			},
//   			StartTime: jsii.String("startTime"),
//   			Window: &WindowProperty{
//   				Duration: jsii.Number(123),
//   				DurationUnit: jsii.String("durationUnit"),
//   			},
//   		},
//   	},
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
//   	Name: jsii.String("name"),
//   	RequestBasedSli: &RequestBasedSliProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		MetricThreshold: jsii.Number(123),
//   		RequestBasedSliMetric: &RequestBasedSliMetricProperty{
//   			DependencyConfig: &DependencyConfigProperty{
//   				DependencyKeyAttributes: map[string]*string{
//   					"dependencyKeyAttributesKey": jsii.String("dependencyKeyAttributes"),
//   				},
//   				DependencyOperationName: jsii.String("dependencyOperationName"),
//   			},
//   			KeyAttributes: map[string]*string{
//   				"keyAttributesKey": jsii.String("keyAttributes"),
//   			},
//   			MetricType: jsii.String("metricType"),
//   			MonitoredRequestCountMetric: &MonitoredRequestCountMetricProperty{
//   				BadCountMetric: []interface{}{
//   					&MetricDataQueryProperty{
//   						AccountId: jsii.String("accountId"),
//   						Expression: jsii.String("expression"),
//   						Id: jsii.String("id"),
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
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   				GoodCountMetric: []interface{}{
//   					&MetricDataQueryProperty{
//   						AccountId: jsii.String("accountId"),
//   						Expression: jsii.String("expression"),
//   						Id: jsii.String("id"),
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
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			OperationName: jsii.String("operationName"),
//   			TotalRequestCountMetric: []interface{}{
//   				&MetricDataQueryProperty{
//   					AccountId: jsii.String("accountId"),
//   					Expression: jsii.String("expression"),
//   					Id: jsii.String("id"),
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
//   						Unit: jsii.String("unit"),
//   					},
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	Sli: &SliProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		MetricThreshold: jsii.Number(123),
//   		SliMetric: &SliMetricProperty{
//   			DependencyConfig: &DependencyConfigProperty{
//   				DependencyKeyAttributes: map[string]*string{
//   					"dependencyKeyAttributesKey": jsii.String("dependencyKeyAttributes"),
//   				},
//   				DependencyOperationName: jsii.String("dependencyOperationName"),
//   			},
//   			KeyAttributes: map[string]*string{
//   				"keyAttributesKey": jsii.String("keyAttributes"),
//   			},
//   			MetricDataQueries: []interface{}{
//   				&MetricDataQueryProperty{
//   					AccountId: jsii.String("accountId"),
//   					Expression: jsii.String("expression"),
//   					Id: jsii.String("id"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html
//
type CfnServiceLevelObjectiveMixinProps struct {
	// Each object in this array defines the length of the look-back window used to calculate one burn rate metric for this SLO.
	//
	// The burn rate measures how fast the service is consuming the error budget, relative to the attainment goal of the SLO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-burnrateconfigurations
	//
	BurnRateConfigurations interface{} `field:"optional" json:"burnRateConfigurations" yaml:"burnRateConfigurations"`
	// An optional description for this SLO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-description
	//
	// Default: - "No description".
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The time window to be excluded from the SLO performance metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-exclusionwindows
	//
	ExclusionWindows interface{} `field:"optional" json:"exclusionWindows" yaml:"exclusionWindows"`
	// This structure contains the attributes that determine the goal of an SLO.
	//
	// This includes the time period for evaluation and the attainment threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-goal
	//
	Goal interface{} `field:"optional" json:"goal" yaml:"goal"`
	// A name for this SLO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
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

