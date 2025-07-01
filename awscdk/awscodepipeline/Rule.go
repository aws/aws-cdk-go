package awscodepipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a rule in AWS CodePipeline that can be used to add conditions and controls to pipeline execution.
//
// Example:
//   var sourceAction codeStarConnectionsSourceAction
//   var buildAction codeBuildAction
//
//
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
//   	PipelineType: codepipeline.PipelineType_V2,
//   	Stages: []stageProps{
//   		&stageProps{
//   			StageName: jsii.String("Source"),
//   			Actions: []iAction{
//   				sourceAction,
//   			},
//   		},
//   		&stageProps{
//   			StageName: jsii.String("Build"),
//   			Actions: []*iAction{
//   				buildAction,
//   			},
//   			// BeforeEntry condition - checks before entering the stage
//   			BeforeEntry: &Conditions{
//   				Conditions: []condition{
//   					&condition{
//   						Rules: []rule{
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
//   				Conditions: []*condition{
//   					&condition{
//   						Result: codepipeline.Result_FAIL,
//   						Rules: []*rule{
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
//   				Conditions: []*condition{
//   					&condition{
//   						Result: codepipeline.Result_ROLLBACK,
//   						Rules: []*rule{
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
type Rule interface {
	// The name of the rule, if specified in the properties.
	RuleName() *string
	// Returns a reference to the rule that can be used in pipeline stage conditions.
	//
	// Returns: A string in the format "#{rule.ruleName}" that can be used to reference this rule
	Reference() *string
}

// The jsii proxy struct for Rule
type jsiiProxy_Rule struct {
	_ byte // padding
}

func (j *jsiiProxy_Rule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}


// Creates a new Rule instance.
func NewRule(props *RuleProps) Rule {
	_init_.Initialize()

	if err := validateNewRuleParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Rule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Rule",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new Rule instance.
func NewRule_Override(r Rule, props *RuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Rule",
		[]interface{}{props},
		r,
	)
}

func (r *jsiiProxy_Rule) Reference() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"reference",
		nil, // no parameters
		&returns,
	)

	return returns
}

