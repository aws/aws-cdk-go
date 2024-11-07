package awscodepipeline


// The configuration that specifies the result, such as rollback, to occur upon stage failure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   failureConditionsProperty := &FailureConditionsProperty{
//   	Conditions: []interface{}{
//   		&ConditionProperty{
//   			Result: jsii.String("result"),
//   			Rules: []interface{}{
//   				&RuleDeclarationProperty{
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
type CfnPipeline_FailureConditionsProperty struct {
	// The conditions that are configured as failure conditions.
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

