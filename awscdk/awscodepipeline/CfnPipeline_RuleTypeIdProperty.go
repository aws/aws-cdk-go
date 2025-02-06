package awscodepipeline


// The ID for the rule type, which is made up of the combined values for category, owner, provider, and version.
//
// For more information about conditions, see [Stage conditions](https://docs.aws.amazon.com/codepipeline/latest/userguide/stage-conditions.html) . For more information about rules, see the [AWS CodePipeline rule reference](https://docs.aws.amazon.com/codepipeline/latest/userguide/rule-reference.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleTypeIdProperty := &RuleTypeIdProperty{
//   	Category: jsii.String("category"),
//   	Owner: jsii.String("owner"),
//   	Provider: jsii.String("provider"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-ruletypeid.html
//
type CfnPipeline_RuleTypeIdProperty struct {
	// A category defines what kind of rule can be run in the stage, and constrains the provider type for the rule.
	//
	// The valid category is `Rule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-ruletypeid.html#cfn-codepipeline-pipeline-ruletypeid-category
	//
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The creator of the rule being called.
	//
	// The valid value for the `Owner` field in the rule category is `AWS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-ruletypeid.html#cfn-codepipeline-pipeline-ruletypeid-owner
	//
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// The rule provider, such as the `DeploymentWindow` rule.
	//
	// For a list of rule provider names, see the rules listed in the [AWS CodePipeline rule reference](https://docs.aws.amazon.com/codepipeline/latest/userguide/rule-reference.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-ruletypeid.html#cfn-codepipeline-pipeline-ruletypeid-provider
	//
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// A string that describes the rule version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-ruletypeid.html#cfn-codepipeline-pipeline-ruletypeid-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

