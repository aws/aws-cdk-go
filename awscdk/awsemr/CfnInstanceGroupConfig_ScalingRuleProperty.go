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
//   scalingRuleProperty := &ScalingRuleProperty{
//   	Action: &ScalingActionProperty{
//   		SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   			ScalingAdjustment: jsii.Number(123),
//
//   			// the properties below are optional
//   			AdjustmentType: jsii.String("adjustmentType"),
//   			CoolDown: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		Market: jsii.String("market"),
//   	},
//   	Name: jsii.String("name"),
//   	Trigger: &ScalingTriggerProperty{
//   		CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			MetricName: jsii.String("metricName"),
//   			Period: jsii.Number(123),
//   			Threshold: jsii.Number(123),
//
//   			// the properties below are optional
//   			Dimensions: []interface{}{
//   				&MetricDimensionProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			EvaluationPeriods: jsii.Number(123),
//   			Namespace: jsii.String("namespace"),
//   			Statistic: jsii.String("statistic"),
//   			Unit: jsii.String("unit"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-scalingrule.html
//
type CfnInstanceGroupConfig_ScalingRuleProperty struct {
	// The conditions that trigger an automatic scaling activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-scalingrule.html#cfn-emr-instancegroupconfig-scalingrule-action
	//
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// The name used to identify an automatic scaling rule.
	//
	// Rule names must be unique within a scaling policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-scalingrule.html#cfn-emr-instancegroupconfig-scalingrule-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The CloudWatch alarm definition that determines when automatic scaling activity is triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-scalingrule.html#cfn-emr-instancegroupconfig-scalingrule-trigger
	//
	Trigger interface{} `field:"required" json:"trigger" yaml:"trigger"`
	// A friendly, more verbose description of the automatic scaling rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-scalingrule.html#cfn-emr-instancegroupconfig-scalingrule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

