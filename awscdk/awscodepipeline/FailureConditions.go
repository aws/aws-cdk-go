package awscodepipeline


// The configuration that specifies the result, such as rollback, to occur upon stage failure.
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
type FailureConditions struct {
	// The conditions that are configured as entry conditions, making check to succeed the stage, or fail the stage.
	// Default: - No conditions are configured.
	//
	Conditions *[]*Condition `field:"optional" json:"conditions" yaml:"conditions"`
	// The specified result for when the failure conditions are met, such as rolling back the stage.
	// Default: FAIL.
	//
	Result Result `field:"optional" json:"result" yaml:"result"`
	// The method that you want to configure for automatic stage retry on stage failure.
	// Default: ALL_ACTIONS.
	//
	RetryMode RetryMode `field:"optional" json:"retryMode" yaml:"retryMode"`
}

