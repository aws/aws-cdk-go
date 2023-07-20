package awsemr


// `AutoScalingPolicy` is a subproperty of `InstanceGroupConfig` .
//
// `AutoScalingPolicy` defines how an instance group dynamically adds and terminates EC2 instances in response to the value of a CloudWatch metric. For more information, see [Using Automatic Scaling in Amazon EMR](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-automatic-scaling.html) in the *Amazon EMR Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingPolicyProperty := &AutoScalingPolicyProperty{
//   	Constraints: &ScalingConstraintsProperty{
//   		MaxCapacity: jsii.Number(123),
//   		MinCapacity: jsii.Number(123),
//   	},
//   	Rules: []interface{}{
//   		&ScalingRuleProperty{
//   			Action: &ScalingActionProperty{
//   				SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   					ScalingAdjustment: jsii.Number(123),
//
//   					// the properties below are optional
//   					AdjustmentType: jsii.String("adjustmentType"),
//   					CoolDown: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				Market: jsii.String("market"),
//   			},
//   			Name: jsii.String("name"),
//   			Trigger: &ScalingTriggerProperty{
//   				CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   					ComparisonOperator: jsii.String("comparisonOperator"),
//   					MetricName: jsii.String("metricName"),
//   					Period: jsii.Number(123),
//   					Threshold: jsii.Number(123),
//
//   					// the properties below are optional
//   					Dimensions: []interface{}{
//   						&MetricDimensionProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					EvaluationPeriods: jsii.Number(123),
//   					Namespace: jsii.String("namespace"),
//   					Statistic: jsii.String("statistic"),
//   					Unit: jsii.String("unit"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-autoscalingpolicy.html
//
type CfnCluster_AutoScalingPolicyProperty struct {
	// The upper and lower Amazon EC2 instance limits for an automatic scaling policy.
	//
	// Automatic scaling activity will not cause an instance group to grow above or below these limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-autoscalingpolicy.html#cfn-emr-cluster-autoscalingpolicy-constraints
	//
	Constraints interface{} `field:"required" json:"constraints" yaml:"constraints"`
	// The scale-in and scale-out rules that comprise the automatic scaling policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-autoscalingpolicy.html#cfn-emr-cluster-autoscalingpolicy-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

