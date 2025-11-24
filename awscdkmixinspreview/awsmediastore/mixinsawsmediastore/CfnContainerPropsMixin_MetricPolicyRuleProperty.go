package mixinsawsmediastore


// A setting that enables metrics at the object level.
//
// Each rule contains an object group and an object group name. If the policy includes the MetricPolicyRules parameter, you must include at least one rule. Each metric policy can include up to five rules by default. You can also [request a quota increase](https://docs.aws.amazon.com/servicequotas/home?region=us-east-1#!/services/mediastore/quotas) to allow up to 300 rules per policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricPolicyRuleProperty := &MetricPolicyRuleProperty{
//   	ObjectGroup: jsii.String("objectGroup"),
//   	ObjectGroupName: jsii.String("objectGroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-metricpolicyrule.html
//
type CfnContainerPropsMixin_MetricPolicyRuleProperty struct {
	// A path or file name that defines which objects to include in the group.
	//
	// Wildcards (*) are acceptable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-metricpolicyrule.html#cfn-mediastore-container-metricpolicyrule-objectgroup
	//
	ObjectGroup *string `field:"optional" json:"objectGroup" yaml:"objectGroup"`
	// A name that allows you to refer to the object group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-metricpolicyrule.html#cfn-mediastore-container-metricpolicyrule-objectgroupname
	//
	ObjectGroupName *string `field:"optional" json:"objectGroupName" yaml:"objectGroupName"`
}

