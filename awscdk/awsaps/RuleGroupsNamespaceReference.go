package awsaps


// A reference to a RuleGroupsNamespace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleGroupsNamespaceReference := &RuleGroupsNamespaceReference{
//   	RuleGroupsNamespaceArn: jsii.String("ruleGroupsNamespaceArn"),
//   }
//
type RuleGroupsNamespaceReference struct {
	// The Arn of the RuleGroupsNamespace resource.
	RuleGroupsNamespaceArn *string `field:"required" json:"ruleGroupsNamespaceArn" yaml:"ruleGroupsNamespaceArn"`
}

