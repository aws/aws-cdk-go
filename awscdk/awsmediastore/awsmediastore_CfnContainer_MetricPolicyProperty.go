package awsmediastore


// The metric policy that is associated with the container.
//
// A metric policy allows AWS Elemental MediaStore to send metrics to Amazon CloudWatch. In the policy, you must indicate whether you want MediaStore to send container-level metrics. You can also include rules to define groups of objects that you want MediaStore to send object-level metrics for.
//
// To view examples of how to construct a metric policy for your use case, see [Example Metric Policies](https://docs.aws.amazon.com/mediastore/latest/ug/policies-metric-examples.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricPolicyProperty := &metricPolicyProperty{
//   	containerLevelMetrics: jsii.String("containerLevelMetrics"),
//
//   	// the properties below are optional
//   	metricPolicyRules: []interface{}{
//   		&metricPolicyRuleProperty{
//   			objectGroup: jsii.String("objectGroup"),
//   			objectGroupName: jsii.String("objectGroupName"),
//   		},
//   	},
//   }
//
type CfnContainer_MetricPolicyProperty struct {
	// A setting to enable or disable metrics at the container level.
	ContainerLevelMetrics *string `field:"required" json:"containerLevelMetrics" yaml:"containerLevelMetrics"`
	// A parameter that holds an array of rules that enable metrics at the object level.
	//
	// This parameter is optional, but if you choose to include it, you must also include at least one rule. By default, you can include up to five rules. You can also [request a quota increase](https://docs.aws.amazon.com/servicequotas/home?region=us-east-1#!/services/mediastore/quotas) to allow up to 300 rules per policy.
	MetricPolicyRules interface{} `field:"optional" json:"metricPolicyRules" yaml:"metricPolicyRules"`
}

