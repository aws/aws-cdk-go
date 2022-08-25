package awsmediastore


// A setting that enables metrics at the object level.
//
// Each rule contains an object group and an object group name. If the policy includes the MetricPolicyRules parameter, you must include at least one rule. Each metric policy can include up to five rules by default. You can also [request a quota increase](https://docs.aws.amazon.com/servicequotas/home?region=us-east-1#!/services/mediastore/quotas) to allow up to 300 rules per policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricPolicyRuleProperty := &metricPolicyRuleProperty{
//   	objectGroup: jsii.String("objectGroup"),
//   	objectGroupName: jsii.String("objectGroupName"),
//   }
//
type CfnContainer_MetricPolicyRuleProperty struct {
	// A path or file name that defines which objects to include in the group.
	//
	// Wildcards (*) are acceptable.
	ObjectGroup *string `field:"required" json:"objectGroup" yaml:"objectGroup"`
	// A name that allows you to refer to the object group.
	ObjectGroupName *string `field:"required" json:"objectGroupName" yaml:"objectGroupName"`
}

