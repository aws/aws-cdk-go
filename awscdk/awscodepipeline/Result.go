package awscodepipeline


// The action to be done when the condition is met.
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
type Result string

const (
	// Rollback.
	Result_ROLLBACK Result = "ROLLBACK"
	// Failure.
	Result_FAIL Result = "FAIL"
	// Retry.
	Result_RETRY Result = "RETRY"
	// Skip.
	Result_SKIP Result = "SKIP"
)

