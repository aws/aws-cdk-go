package awscodepipeline


// The conditions for making checks that, if met, succeed a stage.
//
// For more information about conditions, see [Stage conditions](https://docs.aws.amazon.com/codepipeline/latest/userguide/stage-conditions.html) and [How do stage conditions work?](https://docs.aws.amazon.com/codepipeline/latest/userguide/concepts-how-it-works-conditions.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   successConditionsProperty := &SuccessConditionsProperty{
//   	Conditions: []interface{}{
//   		&ConditionProperty{
//   			Result: jsii.String("result"),
//   			Rules: []interface{}{
//   				&RuleDeclarationProperty{
//   					Commands: []*string{
//   						jsii.String("commands"),
//   					},
//   					Configuration: configuration,
//   					InputArtifacts: []interface{}{
//   						&InputArtifactProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Name: jsii.String("name"),
//   					Region: jsii.String("region"),
//   					RoleArn: jsii.String("roleArn"),
//   					RuleTypeId: &RuleTypeIdProperty{
//   						Category: jsii.String("category"),
//   						Owner: jsii.String("owner"),
//   						Provider: jsii.String("provider"),
//   						Version: jsii.String("version"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-successconditions.html
//
type CfnPipeline_SuccessConditionsProperty struct {
	// The conditions that are success conditions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-successconditions.html#cfn-codepipeline-pipeline-successconditions-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
}

