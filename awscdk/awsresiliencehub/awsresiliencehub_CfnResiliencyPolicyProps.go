package awsresiliencehub


// Properties for defining a `CfnResiliencyPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResiliencyPolicyProps := &cfnResiliencyPolicyProps{
//   	policy: map[string]interface{}{
//   		"policyKey": &FailurePolicyProperty{
//   			"rpoInSecs": jsii.Number(123),
//   			"rtoInSecs": jsii.Number(123),
//   		},
//   	},
//   	policyName: jsii.String("policyName"),
//   	tier: jsii.String("tier"),
//
//   	// the properties below are optional
//   	dataLocationConstraint: jsii.String("dataLocationConstraint"),
//   	policyDescription: jsii.String("policyDescription"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnResiliencyPolicyProps struct {
	// The resiliency policy.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// The name of the policy.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The tier for this resiliency policy, ranging from the highest severity ( `MissionCritical` ) to lowest ( `NonCritical` ).
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Specifies a high-level geographical location constraint for where your resilience policy data can be stored.
	DataLocationConstraint *string `field:"optional" json:"dataLocationConstraint" yaml:"dataLocationConstraint"`
	// The description for the policy.
	PolicyDescription *string `field:"optional" json:"policyDescription" yaml:"policyDescription"`
	// The tags assigned to the resource.
	//
	// A tag is a label that you assign to an AWS resource. Each tag consists of a key/value pair.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

