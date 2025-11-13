package interfacesawsconfig


// A reference to a ConfigRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configRuleReference := &ConfigRuleReference{
//   	ConfigRuleArn: jsii.String("configRuleArn"),
//   	ConfigRuleName: jsii.String("configRuleName"),
//   }
//
type ConfigRuleReference struct {
	// The ARN of the ConfigRule resource.
	ConfigRuleArn *string `field:"required" json:"configRuleArn" yaml:"configRuleArn"`
	// The ConfigRuleName of the ConfigRule resource.
	ConfigRuleName *string `field:"required" json:"configRuleName" yaml:"configRuleName"`
}

