package awsresiliencehub


// Properties for defining a `CfnResiliencyPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResiliencyPolicyProps := &CfnResiliencyPolicyProps{
//   	Policy: map[string]interface{}{
//   		"policyKey": &FailurePolicyProperty{
//   			"rpoInSecs": jsii.Number(123),
//   			"rtoInSecs": jsii.Number(123),
//   		},
//   	},
//   	PolicyName: jsii.String("policyName"),
//   	Tier: jsii.String("tier"),
//
//   	// the properties below are optional
//   	DataLocationConstraint: jsii.String("dataLocationConstraint"),
//   	PolicyDescription: jsii.String("policyDescription"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html
//
type CfnResiliencyPolicyProps struct {
	// The resiliency policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html#cfn-resiliencehub-resiliencypolicy-policy
	//
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// The name of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html#cfn-resiliencehub-resiliencypolicy-policyname
	//
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The tier for this resiliency policy, ranging from the highest severity ( `MissionCritical` ) to lowest ( `NonCritical` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html#cfn-resiliencehub-resiliencypolicy-tier
	//
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Specifies a high-level geographical location constraint for where your resilience policy data can be stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html#cfn-resiliencehub-resiliencypolicy-datalocationconstraint
	//
	DataLocationConstraint *string `field:"optional" json:"dataLocationConstraint" yaml:"dataLocationConstraint"`
	// The description for the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html#cfn-resiliencehub-resiliencypolicy-policydescription
	//
	PolicyDescription *string `field:"optional" json:"policyDescription" yaml:"policyDescription"`
	// Tags assigned to the resource.
	//
	// A tag is a label that you assign to an AWS resource. Each tag consists of a key/value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-resiliencypolicy.html#cfn-resiliencehub-resiliencypolicy-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

