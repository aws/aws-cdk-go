package mixinsawscodepipeline


// The configuration that specifies the result, such as rollback, to occur upon stage failure.
//
// For more information about conditions, see [Stage conditions](https://docs.aws.amazon.com/codepipeline/latest/userguide/stage-conditions.html) and [How do stage conditions work?](https://docs.aws.amazon.com/codepipeline/latest/userguide/concepts-how-it-works-conditions.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var configuration interface{}
//
//   failureConditionsProperty := &FailureConditionsProperty{
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
//   	Result: jsii.String("result"),
//   	RetryConfiguration: &RetryConfigurationProperty{
//   		RetryMode: jsii.String("retryMode"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-failureconditions.html
//
type CfnPipelinePropsMixin_FailureConditionsProperty struct {
	// The conditions that are configured as failure conditions.
	//
	// For more information about conditions, see [Stage conditions](https://docs.aws.amazon.com/codepipeline/latest/userguide/stage-conditions.html) and [How do stage conditions work?](https://docs.aws.amazon.com/codepipeline/latest/userguide/concepts-how-it-works-conditions.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-failureconditions.html#cfn-codepipeline-pipeline-failureconditions-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The specified result for when the failure conditions are met, such as rolling back the stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-failureconditions.html#cfn-codepipeline-pipeline-failureconditions-result
	//
	Result *string `field:"optional" json:"result" yaml:"result"`
	// The retry configuration specifies automatic retry for a failed stage, along with the configured retry mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-failureconditions.html#cfn-codepipeline-pipeline-failureconditions-retryconfiguration
	//
	RetryConfiguration interface{} `field:"optional" json:"retryConfiguration" yaml:"retryConfiguration"`
}

