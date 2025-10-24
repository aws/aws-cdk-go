package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for defining a CodePipeline Rule.
//
// Example:
//   var sourceAction CodeStarConnectionsSourceAction
//   var buildAction CodeBuildAction
//
//
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
//   	PipelineType: codepipeline.PipelineType_V2,
//   	Stages: []StageProps{
//   		&StageProps{
//   			StageName: jsii.String("Source"),
//   			Actions: []IAction{
//   				sourceAction,
//   			},
//   		},
//   		&StageProps{
//   			StageName: jsii.String("Build"),
//   			Actions: []IAction{
//   				buildAction,
//   			},
//   			// BeforeEntry condition - checks before entering the stage
//   			BeforeEntry: &Conditions{
//   				Conditions: []Condition{
//   					&Condition{
//   						Rules: []Rule{
//   							codepipeline.NewRule(&RuleProps{
//   								Name: jsii.String("LambdaCheck"),
//   								Provider: jsii.String("LambdaInvoke"),
//   								Version: jsii.String("1"),
//   								Configuration: map[string]*string{
//   									"FunctionName": jsii.String("LambdaFunctionName"),
//   								},
//   							}),
//   						},
//   						Result: codepipeline.Result_FAIL,
//   					},
//   				},
//   			},
//   			// OnSuccess condition - checks after successful stage completion
//   			OnSuccess: &Conditions{
//   				Conditions: []Condition{
//   					&Condition{
//   						Result: codepipeline.Result_FAIL,
//   						Rules: []Rule{
//   							codepipeline.NewRule(&RuleProps{
//   								Name: jsii.String("CloudWatchCheck"),
//   								Provider: jsii.String("LambdaInvoke"),
//   								Version: jsii.String("1"),
//   								Configuration: map[string]*string{
//   									"AlarmName": jsii.String("AlarmName1"),
//   									"WaitTime": jsii.String("300"),
//   									 // 5 minutes
//   									"FunctionName": jsii.String("funcName2"),
//   								},
//   							}),
//   						},
//   					},
//   				},
//   			},
//   			// OnFailure condition - handles stage failure
//   			OnFailure: &FailureConditions{
//   				Conditions: []Condition{
//   					&Condition{
//   						Result: codepipeline.Result_ROLLBACK,
//   						Rules: []Rule{
//   							codepipeline.NewRule(&RuleProps{
//   								Name: jsii.String("RollBackOnFailure"),
//   								Provider: jsii.String("LambdaInvoke"),
//   								Version: jsii.String("1"),
//   								Configuration: map[string]*string{
//   									"AlarmName": jsii.String("Alarm"),
//   									"WaitTime": jsii.String("300"),
//   									 // 5 minutes
//   									"FunctionName": jsii.String("funcName1"),
//   								},
//   							}),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   })
//
type RuleProps struct {
	// The shell commands to run with your commands rule in CodePipeline.
	//
	// All commands are supported except multi-line formats. While CodeBuild logs and permissions are used,
	// you do not need to create any resources in CodeBuild.
	// Default: - No commands.
	//
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// The action configuration fields for the rule.
	//
	// This can include custom parameters specific to the rule type.
	// Default: - No configuration.
	//
	Configuration *map[string]interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The input artifacts fields for the rule, such as specifying an input file for the rule.
	//
	// Each string in the array represents an artifact name that this rule will use as input.
	// Default: - No input artifacts.
	//
	InputArtifacts *[]*string `field:"optional" json:"inputArtifacts" yaml:"inputArtifacts"`
	// The name of the rule that is created for the condition.
	//
	// Must be unique within the pipeline.
	//
	// Example:
	//   "VariableCheck"
	//
	// Default: - A unique name will be generated.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The rule provider that implements the rule's functionality.
	//
	// Example:
	//   "DeploymentWindow"
	//
	// See: AWS CodePipeline rule reference for available providers.
	//
	// Default: - No provider, must be specified if rule is used.
	//
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// The AWS Region for the condition associated with the rule.
	//
	// If not specified, uses the pipeline's region.
	// Default: - Pipeline's region.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The IAM role that the rule will use to execute its actions.
	//
	// The role must have sufficient permissions to perform the rule's tasks.
	// Default: - A new role will be created.
	//
	Role awsiam.Role `field:"optional" json:"role" yaml:"role"`
	// The version of the rule to use.
	//
	// Different versions may have different features or behaviors.
	// Default: '1'.
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

