package awscodepipeline


// The condition for the stage.
//
// A condition is made up of the rules and the result for the condition. For more information about conditions, see [Stage conditions](https://docs.aws.amazon.com/codepipeline/latest/userguide/stage-conditions.html) . For more information about rules, see the [AWS CodePipeline rule reference](https://docs.aws.amazon.com/codepipeline/latest/userguide/rule-reference.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   conditionProperty := &ConditionProperty{
//   	Result: jsii.String("result"),
//   	Rules: []interface{}{
//   		&RuleDeclarationProperty{
//   			Commands: []*string{
//   				jsii.String("commands"),
//   			},
//   			Configuration: configuration,
//   			InputArtifacts: []interface{}{
//   				&InputArtifactProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			Region: jsii.String("region"),
//   			RoleArn: jsii.String("roleArn"),
//   			RuleTypeId: &RuleTypeIdProperty{
//   				Category: jsii.String("category"),
//   				Owner: jsii.String("owner"),
//   				Provider: jsii.String("provider"),
//   				Version: jsii.String("version"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-condition.html
//
type CfnPipeline_ConditionProperty struct {
	// The action to be done when the condition is met.
	//
	// For example, rolling back an execution for a failure condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-condition.html#cfn-codepipeline-pipeline-condition-result
	//
	Result *string `field:"optional" json:"result" yaml:"result"`
	// The rules that make up the condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-condition.html#cfn-codepipeline-pipeline-condition-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

