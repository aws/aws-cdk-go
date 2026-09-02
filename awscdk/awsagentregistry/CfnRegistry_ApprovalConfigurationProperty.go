package awsagentregistry


// Configuration for the registry's record approval workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   approvalConfigurationProperty := &ApprovalConfigurationProperty{
//   	AutoApprovalRules: []*string{
//   		jsii.String("autoApprovalRules"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-approvalconfiguration.html
//
type CfnRegistry_ApprovalConfigurationProperty struct {
	// The rules that determine which registry records are automatically approved on submission.
	//
	// When omitted or empty, submitted records require manual review.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-approvalconfiguration.html#cfn-agentregistry-registry-approvalconfiguration-autoapprovalrules
	//
	AutoApprovalRules *[]*string `field:"optional" json:"autoApprovalRules" yaml:"autoApprovalRules"`
}

