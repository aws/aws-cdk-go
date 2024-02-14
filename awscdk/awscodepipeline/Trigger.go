package awscodepipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var action action
//
//   trigger := awscdk.Aws_codepipeline.NewTrigger(&TriggerProps{
//   	ProviderType: awscdk.*Aws_codepipeline.ProviderType_CODE_STAR_SOURCE_CONNECTION,
//
//   	// the properties below are optional
//   	GitConfiguration: &GitConfiguration{
//   		SourceAction: action,
//
//   		// the properties below are optional
//   		PushFilter: []gitPushFilter{
//   			&gitPushFilter{
//   				TagsExcludes: []*string{
//   					jsii.String("tagsExcludes"),
//   				},
//   				TagsIncludes: []*string{
//   					jsii.String("tagsIncludes"),
//   				},
//   			},
//   		},
//   	},
//   })
//
type Trigger interface {
	// The pipeline source action where the trigger configuration.
	SourceAction() IAction
}

// The jsii proxy struct for Trigger
type jsiiProxy_Trigger struct {
	_ byte // padding
}

func (j *jsiiProxy_Trigger) SourceAction() IAction {
	var returns IAction
	_jsii_.Get(
		j,
		"sourceAction",
		&returns,
	)
	return returns
}


func NewTrigger(props *TriggerProps) Trigger {
	_init_.Initialize()

	if err := validateNewTriggerParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Trigger{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Trigger",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewTrigger_Override(t Trigger, props *TriggerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Trigger",
		[]interface{}{props},
		t,
	)
}

