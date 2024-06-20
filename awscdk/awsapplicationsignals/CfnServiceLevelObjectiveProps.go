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
	// The name of this SLO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// This structure contains information about the performance metric that an SLO monitors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-sli
	//
	Sli interface{} `field:"required" json:"sli" yaml:"sli"`
	// An optional description for this SLO.
	//
	// Default is 'No description'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-description
	//
	// Default: - "No description".
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A structure that contains the attributes that determine the goal of the SLO.
	//
	// This includes the time period for evaluation and the attainment threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-goal
	//
	Goal interface{} `field:"optional" json:"goal" yaml:"goal"`
	// The list of tag keys and values associated with the resource you specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html#cfn-applicationsignals-servicelevelobjective-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

