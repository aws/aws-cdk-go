package mixinsawsresiliencehub


// Properties for CfnResiliencyPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResiliencyPolicyMixinProps := &CfnResiliencyPolicyMixinProps{
//   	DataLocationConstraint: jsii.String("dataLocationConstraint"),
//   	Policy: map[string]interface{}{
//   		"policyKey": &FailurePolicyProperty{
//   			"rpoInSecs": jsii.Number(123),
//   			"rtoInSecs": jsii.Number(123),
//   		},
//   	},
//   	PolicyDescription: jsii.String("policyDescription"),
//   	PolicyName: jsii.String("policyName"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Tier: jsii.String("tier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html
//
type CfnResiliencyPolicyMixinProps struct {
	// Specifies a high-level geographical location constraint for where your resilience policy data can be stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html#cfn-resiliencehub-resiliencypolicy-datalocationconstraint
	//
	DataLocationConstraint *string `field:"optional" json:"dataLocationConstraint" yaml:"dataLocationConstraint"`
	// The resiliency policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html#cfn-resiliencehub-resiliencypolicy-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// Description of the resiliency policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html#cfn-resiliencehub-resiliencypolicy-policydescription
	//
	PolicyDescription *string `field:"optional" json:"policyDescription" yaml:"policyDescription"`
	// The name of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html#cfn-resiliencehub-resiliencypolicy-policyname
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Tags assigned to the resource.
	//
	// A tag is a label that you assign to an AWS resource. Each tag consists of a key/value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html#cfn-resiliencehub-resiliencypolicy-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The tier for this resiliency policy, ranging from the highest severity ( `MissionCritical` ) to lowest ( `NonCritical` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html#cfn-resiliencehub-resiliencypolicy-tier
	//
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

