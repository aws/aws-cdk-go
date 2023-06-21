package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSecurityProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityProfileProps := &CfnSecurityProfileProps{
//   	AdditionalMetricsToRetainV2: []interface{}{
//   		&MetricToRetainProperty{
//   			Metric: jsii.String("metric"),
//
//   			// the properties below are optional
//   			MetricDimension: &MetricDimensionProperty{
//   				DimensionName: jsii.String("dimensionName"),
//
//   				// the properties below are optional
//   				Operator: jsii.String("operator"),
//   			},
//   		},
//   	},
//   	AlertTargets: map[string]interface{}{
//   		"alertTargetsKey": &AlertTargetProperty{
//   			"alertTargetArn": jsii.String("alertTargetArn"),
//   			"roleArn": jsii.String("roleArn"),
//   		},
//   	},
//   	Behaviors: []interface{}{
//   		&BehaviorProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Criteria: &BehaviorCriteriaProperty{
//   				ComparisonOperator: jsii.String("comparisonOperator"),
//   				ConsecutiveDatapointsToAlarm: jsii.Number(123),
//   				ConsecutiveDatapointsToClear: jsii.Number(123),
//   				DurationSeconds: jsii.Number(123),
//   				MlDetectionConfig: &MachineLearningDetectionConfigProperty{
//   					ConfidenceLevel: jsii.String("confidenceLevel"),
//   				},
//   				StatisticalThreshold: &StatisticalThresholdProperty{
//   					Statistic: jsii.String("statistic"),
//   				},
//   				Value: &MetricValueProperty{
//   					Cidrs: []*string{
//   						jsii.String("cidrs"),
//   					},
//   					Count: jsii.String("count"),
//   					Number: jsii.Number(123),
//   					Numbers: []interface{}{
//   						jsii.Number(123),
//   					},
//   					Ports: []interface{}{
//   						jsii.Number(123),
//   					},
//   					Strings: []*string{
//   						jsii.String("strings"),
//   					},
//   				},
//   			},
//   			Metric: jsii.String("metric"),
//   			MetricDimension: &MetricDimensionProperty{
//   				DimensionName: jsii.String("dimensionName"),
//
//   				// the properties below are optional
//   				Operator: jsii.String("operator"),
//   			},
//   			SuppressAlerts: jsii.Boolean(false),
//   		},
//   	},
//   	SecurityProfileDescription: jsii.String("securityProfileDescription"),
//   	SecurityProfileName: jsii.String("securityProfileName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetArns: []*string{
//   		jsii.String("targetArns"),
//   	},
//   }
//
type CfnSecurityProfileProps struct {
	// A list of metrics whose data is retained (stored).
	//
	// By default, data is retained for any metric used in the profile's `behaviors` , but it's also retained for any metric specified here. Can be used with custom metrics; can't be used with dimensions.
	AdditionalMetricsToRetainV2 interface{} `field:"optional" json:"additionalMetricsToRetainV2" yaml:"additionalMetricsToRetainV2"`
	// Specifies the destinations to which alerts are sent.
	//
	// (Alerts are always sent to the console.) Alerts are generated when a device (thing) violates a behavior.
	AlertTargets interface{} `field:"optional" json:"alertTargets" yaml:"alertTargets"`
	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors interface{} `field:"optional" json:"behaviors" yaml:"behaviors"`
	// A description of the security profile.
	SecurityProfileDescription *string `field:"optional" json:"securityProfileDescription" yaml:"securityProfileDescription"`
	// The name you gave to the security profile.
	SecurityProfileName *string `field:"optional" json:"securityProfileName" yaml:"securityProfileName"`
	// Metadata that can be used to manage the security profile.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the target (thing group) to which the security profile is attached.
	TargetArns *[]*string `field:"optional" json:"targetArns" yaml:"targetArns"`
}

