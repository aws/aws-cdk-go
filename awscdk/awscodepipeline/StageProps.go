package awscodepipeline


// Construction properties of a Pipeline Stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var action action
//
//   stageProps := &StageProps{
//   	StageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	Actions: []iAction{
//   		action,
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

