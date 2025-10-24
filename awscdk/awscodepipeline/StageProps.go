package awscodepipeline


// Construction properties of a Pipeline Stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var action Action
//   var rule Rule
//
//   stageProps := &StageProps{
//   	StageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	Actions: []IAction{
//   		action,
//   	},
//   	BeforeEntry: &Conditions{
//   		Conditions: []Condition{
//   			&Condition{
//   				Result: awscdk.Aws_codepipeline.Result_ROLLBACK,
//   				Rules: []Rule{
//   					rule,
//   				},
//   			},
//   		},
//   	},
//   	OnFailure: &FailureConditions{
//   		Conditions: []Condition{
//   			&Condition{
//   				Result: awscdk.*Aws_codepipeline.Result_ROLLBACK,
//   				Rules: []Rule{
//   					rule,
//   				},
//   			},
//   		},
//   		Result: awscdk.*Aws_codepipeline.Result_ROLLBACK,
//   		RetryMode: awscdk.*Aws_codepipeline.RetryMode_ALL_ACTIONS,
//   	},
//   	OnSuccess: &Conditions{
//   		Conditions: []Condition{
//   			&Condition{
//   				Result: awscdk.*Aws_codepipeline.Result_ROLLBACK,
//   				Rules: []Rule{
//   					rule,
//   				},
//   			},
//   		},
//   	},
//   	TransitionDisabledReason: jsii.String("transitionDisabledReason"),
//   	TransitionToEnabled: jsii.Boolean(false),
//   }
//
type StageProps struct {
	// The physical, human-readable name to assign to this Pipeline Stage.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// The list of Actions to create this Stage with.
	//
	// You can always add more Actions later by calling `IStage#addAction`.
	Actions *[]IAction `field:"optional" json:"actions" yaml:"actions"`
	// The method to use when a stage allows entry.
	// Default: - No conditions are applied before stage entry.
	//
	BeforeEntry *Conditions `field:"optional" json:"beforeEntry" yaml:"beforeEntry"`
	// The method to use when a stage has not completed successfully.
	// Default: - No failure conditions are applied.
	//
	OnFailure *FailureConditions `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The method to use when a stage has succeeded.
	// Default: - No success conditions are applied.
	//
	OnSuccess *Conditions `field:"optional" json:"onSuccess" yaml:"onSuccess"`
	// The reason for disabling transition to this stage.
	//
	// Only applicable
	// if `transitionToEnabled` is set to `false`.
	// Default: 'Transition disabled'.
	//
	TransitionDisabledReason *string `field:"optional" json:"transitionDisabledReason" yaml:"transitionDisabledReason"`
	// Whether to enable transition to this stage.
	// Default: true.
	//
	TransitionToEnabled *bool `field:"optional" json:"transitionToEnabled" yaml:"transitionToEnabled"`
}

