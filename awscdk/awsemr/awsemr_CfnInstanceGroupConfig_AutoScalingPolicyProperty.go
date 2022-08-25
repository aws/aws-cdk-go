package awsemr


// `AutoScalingPolicy` defines how an instance group dynamically adds and terminates EC2 instances in response to the value of a CloudWatch metric.
//
// For more information, see [Using Automatic Scaling in Amazon EMR](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-automatic-scaling.html) in the *Amazon EMR Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingPolicyProperty := &autoScalingPolicyProperty{
//   	constraints: &scalingConstraintsProperty{
//   		maxCapacity: jsii.Number(123),
//   		minCapacity: jsii.Number(123),
//   	},
//   	rules: []interface{}{
//   		&scalingRuleProperty{
//   			action: &scalingActionProperty{
//   				simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   					scalingAdjustment: jsii.Number(123),
//
//   					// the properties below are optional
//   					adjustmentType: jsii.String("adjustmentType"),
//   					coolDown: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				market: jsii.String("market"),
//   			},
//   			name: jsii.String("name"),
//   			trigger: &scalingTriggerProperty{
//   				cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					metricName: jsii.String("metricName"),
//   					period: jsii.Number(123),
//   					threshold: jsii.Number(123),
//
//   					// the properties below are optional
//   					dimensions: []interface{}{
//   						&metricDimensionProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					evaluationPeriods: jsii.Number(123),
//   					namespace: jsii.String("namespace"),
//   					statistic: jsii.String("statistic"),
//   					unit: jsii.String("unit"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   		},
//   	},
//   }
//
type CfnInstanceGroupConfig_AutoScalingPolicyProperty struct {
	// The upper and lower EC2 instance limits for an automatic scaling policy.
	//
	// Automatic scaling activity will not cause an instance group to grow above or below these limits.
	Constraints interface{} `field:"required" json:"constraints" yaml:"constraints"`
	// The scale-in and scale-out rules that comprise the automatic scaling policy.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

