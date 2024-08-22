package awscodepipeline


// Represents information about the rule to be created for an associated condition.
//
// An example would be creating a new rule for an entry condition, such as a rule that checks for a test result before allowing the run to enter the deployment stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   ruleDeclarationProperty := &RuleDeclarationProperty{
//   	Configuration: configuration,
//   	InputArtifacts: []interface{}{
//   		&InputArtifactProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Region: jsii.String("region"),
//   	RoleArn: jsii.String("roleArn"),
//   	RuleTypeId: &RuleTypeIdProperty{
//   		Category: jsii.String("category"),
//   		Owner: jsii.String("owner"),
//   		Provider: jsii.String("provider"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-ruledeclaration.html
//
type CfnPipeline_RuleDeclarationProperty struct {
	// The action configuration fields for the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-ruledeclaration.html#cfn-codepipeline-pipeline-ruledeclaration-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The input artifacts fields for the rule, such as specifying an input file for the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-ruledeclaration.html#cfn-codepipeline-pipeline-ruledeclaration-inputartifacts
	//
	InputArtifacts interface{} `field:"optional" json:"inputArtifacts" yaml:"inputArtifacts"`
	// The name of the rule that is created for the condition, such as CheckAllResults.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-ruledeclaration.html#cfn-codepipeline-pipeline-ruledeclaration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Region for the condition associated with the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-ruledeclaration.html#cfn-codepipeline-pipeline-ruledeclaration-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The pipeline role ARN associated with the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-ruledeclaration.html#cfn-codepipeline-pipeline-ruledeclaration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The ID for the rule type, which is made up of the combined values for category, owner, provider, and version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-ruledeclaration.html#cfn-codepipeline-pipeline-ruledeclaration-ruletypeid
	//
	RuleTypeId interface{} `field:"optional" json:"ruleTypeId" yaml:"ruleTypeId"`
}

