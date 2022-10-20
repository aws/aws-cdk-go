package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipelineactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

// A CodePipeline source action for BitBucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var role role
//
//   bitBucketSourceAction := awscdk.Aws_codepipeline_actions.NewBitBucketSourceAction(&bitBucketSourceActionProps{
//   	actionName: jsii.String("actionName"),
//   	connectionArn: jsii.String("connectionArn"),
//   	output: artifact,
//   	owner: jsii.String("owner"),
//   	repo: jsii.String("repo"),
//
//   	// the properties below are optional
//   	branch: jsii.String("branch"),
//   	codeBuildCloneOutput: jsii.Boolean(false),
//   	role: role,
//   	runOrder: jsii.Number(123),
//   	triggerOnPush: jsii.Boolean(false),
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   })
//
// Deprecated: use CodeStarConnectionsSourceAction instead.
type BitBucketSourceAction interface {
	awscodepipeline.IAction
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	// Deprecated: use CodeStarConnectionsSourceAction instead.
	ActionProperties() *awscodepipeline.ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	// Deprecated: use CodeStarConnectionsSourceAction instead.
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	// Deprecated: use CodeStarConnectionsSourceAction instead.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
}

// The jsii proxy struct for BitBucketSourceAction
type jsiiProxy_BitBucketSourceAction struct {
	internal.Type__awscodepipelineIAction
}

func (j *jsiiProxy_BitBucketSourceAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}


// Deprecated: use CodeStarConnectionsSourceAction instead.
func NewBitBucketSourceAction(props *BitBucketSourceActionProps) BitBucketSourceAction {
	_init_.Initialize()

	if err := validateNewBitBucketSourceActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BitBucketSourceAction{}

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.BitBucketSourceAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: use CodeStarConnectionsSourceAction instead.
func NewBitBucketSourceAction_Override(b BitBucketSourceAction, props *BitBucketSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.BitBucketSourceAction",
		[]interface{}{props},
		b,
	)
}

func (b *jsiiProxy_BitBucketSourceAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := b.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BitBucketSourceAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := b.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		b,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

