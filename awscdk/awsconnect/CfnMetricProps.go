package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMetric`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMetricProps := &CfnMetricProps{
//   	InstanceArn: jsii.String("instanceArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	MetricCalculation: &MetricCalculationProperty{
//   		Calculation: jsii.String("calculation"),
//   		CalculationComponents: []interface{}{
//   			&CalculationComponentProperty{
//   				Alias: jsii.String("alias"),
//
//   				// the properties below are optional
//   				MetricFilters: []interface{}{
//   					&MetricFilterProperty{
//   						MetricFilterKey: jsii.String("metricFilterKey"),
//
//   						// the properties below are optional
//   						BooleanCondition: &MetricFilterBooleanConditionProperty{
//   							Comparison: jsii.String("comparison"),
//   						},
//   						Negate: jsii.Boolean(false),
//   						NumberCondition: &MetricFilterNumberConditionProperty{
//   							Comparison: jsii.String("comparison"),
//   							Values: []interface{}{
//   								jsii.Number(123),
//   							},
//   						},
//   						StringCondition: &MetricFilterStringConditionProperty{
//   							Comparison: jsii.String("comparison"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   				},
//   				MetricId: jsii.String("metricId"),
//   				MetricName: jsii.String("metricName"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	PositiveTrendIndicator: jsii.String("positiveTrendIndicator"),
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-metric.html
//
type CfnMetricProps struct {
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-metric.html#cfn-connect-metric-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The description of the custom metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-metric.html#cfn-connect-metric-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The calculation configuration for the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-metric.html#cfn-connect-metric-metriccalculation
	//
	MetricCalculation interface{} `field:"optional" json:"metricCalculation" yaml:"metricCalculation"`
	// The name of the custom metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-metric.html#cfn-connect-metric-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates how to classify a positive trend in metric data on the UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-metric.html#cfn-connect-metric-positivetrendindicator
	//
	PositiveTrendIndicator *string `field:"optional" json:"positiveTrendIndicator" yaml:"positiveTrendIndicator"`
	// The status of the custom metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-metric.html#cfn-connect-metric-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// One or more tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-metric.html#cfn-connect-metric-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Display unit for the metric data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-metric.html#cfn-connect-metric-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

