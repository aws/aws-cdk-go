package mixinsawsmediastore


// The metric policy that is associated with the container.
//
// A metric policy allows AWS Elemental MediaStore to send metrics to Amazon CloudWatch. In the policy, you must indicate whether you want MediaStore to send container-level metrics. You can also include rules to define groups of objects that you want MediaStore to send object-level metrics for.
//
// To view examples of how to construct a metric policy for your use case, see [Example Metric Policies](https://docs.aws.amazon.com/mediastore/latest/ug/policies-metric-examples.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricPolicyProperty := &MetricPolicyProperty{
//   	ContainerLevelMetrics: jsii.String("containerLevelMetrics"),
//   	MetricPolicyRules: []interface{}{
//   		&MetricPolicyRuleProperty{
//   			ObjectGroup: jsii.String("objectGroup"),
//   			ObjectGroupName: jsii.String("objectGroupName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-metricpolicy.html
//
type CfnContainerPropsMixin_MetricPolicyProperty struct {
	// A setting to enable or disable metrics at the container level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-metricpolicy.html#cfn-mediastore-container-metricpolicy-containerlevelmetrics
	//
	ContainerLevelMetrics *string `field:"optional" json:"containerLevelMetrics" yaml:"containerLevelMetrics"`
	// A parameter that holds an array of rules that enable metrics at the object level.
	//
	// This parameter is optional, but if you choose to include it, you must also include at least one rule. By default, you can include up to five rules. You can also [request a quota increase](https://docs.aws.amazon.com/servicequotas/home?region=us-east-1#!/services/mediastore/quotas) to allow up to 300 rules per policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-metricpolicy.html#cfn-mediastore-container-metricpolicy-metricpolicyrules
	//
	MetricPolicyRules interface{} `field:"optional" json:"metricPolicyRules" yaml:"metricPolicyRules"`
}

