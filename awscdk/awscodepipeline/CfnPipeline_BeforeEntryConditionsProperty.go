package awscodepipeline


// The conditions for making checks for entry to a stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   beforeEntryConditionsProperty := &BeforeEntryConditionsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-beforeentryconditions.html
//
type CfnPipeline_BeforeEntryConditionsProperty struct {
	// The conditions that are configured as entry conditions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-beforeentryconditions.html#cfn-codepipeline-pipeline-beforeentryconditions-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
}

