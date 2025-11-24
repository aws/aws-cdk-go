package mixinsawsemr


// `ScalingRule` is a subproperty of the `AutoScalingPolicy` property type.
//
// `ScalingRule` defines the scale-in or scale-out rules for scaling activity, including the CloudWatch metric alarm that triggers activity, how EC2 instances are added or removed, and the periodicity of adjustments. The automatic scaling policy for an instance group can comprise one or more automatic scaling rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scalingRuleProperty := &ScalingRuleProperty{
//   	Action: &ScalingActionProperty{
//   		Market: jsii.String("market"),
//   		SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   			AdjustmentType: jsii.String("adjustmentType"),
//   			CoolDown: jsii.Number(123),
//   			ScalingAdjustment: jsii.Number(123),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Trigger: &ScalingTriggerProperty{
//   		CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			Dimensions: []interface{}{
//   				&MetricDimensionProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			EvaluationPeriods: jsii.Number(123),
//   			MetricName: jsii.String("metricName"),
//   			Namespace: jsii.String("namespace"),
//   			Period: jsii.Number(123),
//   			Statistic: jsii.String("statistic"),
//   			Threshold: jsii.Number(123),
//   			Unit: jsii.String("unit"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scalingrule.html
//
type CfnClusterPropsMixin_ScalingRuleProperty struct {
	// The conditions that trigger an automatic scaling activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scalingrule.html#cfn-emr-cluster-scalingrule-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// A friendly, more verbose description of the automatic scaling rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scalingrule.html#cfn-emr-cluster-scalingrule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name used to identify an automatic scaling rule.
	//
	// Rule names must be unique within a scaling policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scalingrule.html#cfn-emr-cluster-scalingrule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The CloudWatch alarm definition that determines when automatic scaling activity is triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scalingrule.html#cfn-emr-cluster-scalingrule-trigger
	//
	Trigger interface{} `field:"optional" json:"trigger" yaml:"trigger"`
}

