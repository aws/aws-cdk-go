package awsemr


// `ScalingRule` is a subproperty of the `AutoScalingPolicy` property type.
//
// `ScalingRule` defines the scale-in or scale-out rules for scaling activity, including the CloudWatch metric alarm that triggers activity, how EC2 instances are added or removed, and the periodicity of adjustments. The automatic scaling policy for an instance group can comprise one or more automatic scaling rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingRuleProperty := &scalingRuleProperty{
//   	action: &scalingActionProperty{
//   		simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   			scalingAdjustment: jsii.Number(123),
//
//   			// the properties below are optional
//   			adjustmentType: jsii.String("adjustmentType"),
//   			coolDown: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		market: jsii.String("market"),
//   	},
//   	name: jsii.String("name"),
//   	trigger: &scalingTriggerProperty{
//   		cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			metricName: jsii.String("metricName"),
//   			period: jsii.Number(123),
//   			threshold: jsii.Number(123),
//
//   			// the properties below are optional
//   			dimensions: []interface{}{
//   				&metricDimensionProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			evaluationPeriods: jsii.Number(123),
//   			namespace: jsii.String("namespace"),
//   			statistic: jsii.String("statistic"),
//   			unit: jsii.String("unit"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnInstanceGroupConfig_ScalingRuleProperty struct {
	// The conditions that trigger an automatic scaling activity.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// The name used to identify an automatic scaling rule.
	//
	// Rule names must be unique within a scaling policy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The CloudWatch alarm definition that determines when automatic scaling activity is triggered.
	Trigger interface{} `field:"required" json:"trigger" yaml:"trigger"`
	// A friendly, more verbose description of the automatic scaling rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

