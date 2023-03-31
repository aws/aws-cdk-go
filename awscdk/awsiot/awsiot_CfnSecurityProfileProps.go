package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnSecurityProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityProfileProps := &cfnSecurityProfileProps{
//   	additionalMetricsToRetainV2: []interface{}{
//   		&metricToRetainProperty{
//   			metric: jsii.String("metric"),
//
//   			// the properties below are optional
//   			metricDimension: &metricDimensionProperty{
//   				dimensionName: jsii.String("dimensionName"),
//
//   				// the properties below are optional
//   				operator: jsii.String("operator"),
//   			},
//   		},
//   	},
//   	alertTargets: map[string]interface{}{
//   		"alertTargetsKey": &AlertTargetProperty{
//   			"alertTargetArn": jsii.String("alertTargetArn"),
//   			"roleArn": jsii.String("roleArn"),
//   		},
//   	},
//   	behaviors: []interface{}{
//   		&behaviorProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			criteria: &behaviorCriteriaProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				consecutiveDatapointsToAlarm: jsii.Number(123),
//   				consecutiveDatapointsToClear: jsii.Number(123),
//   				durationSeconds: jsii.Number(123),
//   				mlDetectionConfig: &machineLearningDetectionConfigProperty{
//   					confidenceLevel: jsii.String("confidenceLevel"),
//   				},
//   				statisticalThreshold: &statisticalThresholdProperty{
//   					statistic: jsii.String("statistic"),
//   				},
//   				value: &metricValueProperty{
//   					cidrs: []*string{
//   						jsii.String("cidrs"),
//   					},
//   					count: jsii.String("count"),
//   					number: jsii.Number(123),
//   					numbers: []interface{}{
//   						jsii.Number(123),
//   					},
//   					ports: []interface{}{
//   						jsii.Number(123),
//   					},
//   					strings: []*string{
//   						jsii.String("strings"),
//   					},
//   				},
//   			},
//   			metric: jsii.String("metric"),
//   			metricDimension: &metricDimensionProperty{
//   				dimensionName: jsii.String("dimensionName"),
//
//   				// the properties below are optional
//   				operator: jsii.String("operator"),
//   			},
//   			suppressAlerts: jsii.Boolean(false),
//   		},
//   	},
//   	securityProfileDescription: jsii.String("securityProfileDescription"),
//   	securityProfileName: jsii.String("securityProfileName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetArns: []*string{
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

