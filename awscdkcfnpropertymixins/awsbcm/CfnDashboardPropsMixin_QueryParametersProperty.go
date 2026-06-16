package awsbcm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var costAndUsageExpressionProperty_ CostAndUsageExpressionProperty
//   var expressionProperty_ ExpressionProperty
//
//   queryParametersProperty := &QueryParametersProperty{
//   	CostAndUsage: &CostAndUsageQueryProperty{
//   		Filter: &CostAndUsageExpressionProperty{
//   			And: []interface{}{
//   				costAndUsageExpressionProperty_,
//   			},
//   			CostCategories: &CostCategoryValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Dimensions: &DimensionValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Not: costAndUsageExpressionProperty_,
//   			Or: []interface{}{
//   				costAndUsageExpressionProperty_,
//   			},
//   			Tags: &TagValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		Granularity: jsii.String("granularity"),
//   		GroupBy: []interface{}{
//   			&GroupDefinitionProperty{
//   				Key: jsii.String("key"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Metrics: []*string{
//   			jsii.String("metrics"),
//   		},
//   		TimeRange: &DateTimeRangeProperty{
//   			EndTime: &DateTimeValueProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   			StartTime: &DateTimeValueProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	ReservationCoverage: &ReservationCoverageQueryProperty{
//   		Filter: &ExpressionProperty{
//   			And: []interface{}{
//   				expressionProperty_,
//   			},
//   			CostCategories: &CostCategoryValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Dimensions: &DimensionValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Not: expressionProperty_,
//   			Tags: &TagValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		Granularity: jsii.String("granularity"),
//   		GroupBy: []interface{}{
//   			&GroupDefinitionProperty{
//   				Key: jsii.String("key"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Metrics: []*string{
//   			jsii.String("metrics"),
//   		},
//   		TimeRange: &DateTimeRangeProperty{
//   			EndTime: &DateTimeValueProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   			StartTime: &DateTimeValueProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	ReservationUtilization: &ReservationUtilizationQueryProperty{
//   		Filter: &ExpressionProperty{
//   			And: []interface{}{
//   				expressionProperty_,
//   			},
//   			CostCategories: &CostCategoryValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Dimensions: &DimensionValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Not: expressionProperty_,
//   			Tags: &TagValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		Granularity: jsii.String("granularity"),
//   		GroupBy: []interface{}{
//   			&GroupDefinitionProperty{
//   				Key: jsii.String("key"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		TimeRange: &DateTimeRangeProperty{
//   			EndTime: &DateTimeValueProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   			StartTime: &DateTimeValueProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	SavingsPlansCoverage: &SavingsPlansCoverageQueryProperty{
//   		Filter: &ExpressionProperty{
//   			And: []interface{}{
//   				expressionProperty_,
//   			},
//   			CostCategories: &CostCategoryValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Dimensions: &DimensionValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Not: expressionProperty_,
//   			Tags: &TagValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		Granularity: jsii.String("granularity"),
//   		GroupBy: []interface{}{
//   			&GroupDefinitionProperty{
//   				Key: jsii.String("key"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Metrics: []*string{
//   			jsii.String("metrics"),
//   		},
//   		TimeRange: &DateTimeRangeProperty{
//   			EndTime: &DateTimeValueProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   			StartTime: &DateTimeValueProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	SavingsPlansUtilization: &SavingsPlansUtilizationQueryProperty{
//   		Filter: &ExpressionProperty{
//   			And: []interface{}{
//   				expressionProperty_,
//   			},
//   			CostCategories: &CostCategoryValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Dimensions: &DimensionValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Not: expressionProperty_,
//   			Tags: &TagValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		Granularity: jsii.String("granularity"),
//   		TimeRange: &DateTimeRangeProperty{
//   			EndTime: &DateTimeValueProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   			StartTime: &DateTimeValueProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-queryparameters.html
//
type CfnDashboardPropsMixin_QueryParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-queryparameters.html#cfn-bcm-dashboard-queryparameters-costandusage
	//
	CostAndUsage interface{} `field:"optional" json:"costAndUsage" yaml:"costAndUsage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-queryparameters.html#cfn-bcm-dashboard-queryparameters-reservationcoverage
	//
	ReservationCoverage interface{} `field:"optional" json:"reservationCoverage" yaml:"reservationCoverage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-queryparameters.html#cfn-bcm-dashboard-queryparameters-reservationutilization
	//
	ReservationUtilization interface{} `field:"optional" json:"reservationUtilization" yaml:"reservationUtilization"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-queryparameters.html#cfn-bcm-dashboard-queryparameters-savingsplanscoverage
	//
	SavingsPlansCoverage interface{} `field:"optional" json:"savingsPlansCoverage" yaml:"savingsPlansCoverage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-queryparameters.html#cfn-bcm-dashboard-queryparameters-savingsplansutilization
	//
	SavingsPlansUtilization interface{} `field:"optional" json:"savingsPlansUtilization" yaml:"savingsPlansUtilization"`
}

