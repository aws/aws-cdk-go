package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awseventstargets/internal"
)

// Allows the pipeline to be used as an EventBridge rule target.
//
// Example:
//   // A pipeline being used as a target for a CloudWatch event rule.
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//   import events "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline pipeline
//
//
//   // kick off the pipeline every day
//   rule := events.NewRule(this, jsii.String("Daily"), &ruleProps{
//   	schedule: events.schedule.rate(awscdk.Duration.days(jsii.Number(1))),
//   })
//   rule.addTarget(targets.NewCodePipeline(pipeline))
//
// Experimental.
type CodePipeline interface {
	awsevents.IRuleTarget
	// Returns the rule target specification.
	//
	// NOTE: Do not use the various `inputXxx` options. They can be set in a call to `addTarget`.
	// Experimental.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for CodePipeline
type jsiiProxy_CodePipeline struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewCodePipeline(pipeline awscodepipeline.IPipeline, options *CodePipelineTargetOptions) CodePipeline {
	_init_.Initialize()

	j := jsiiProxy_CodePipeline{}

	_jsii_.Create(
		"monocdk.aws_events_targets.CodePipeline",
		[]interface{}{pipeline, options},
		&j,
	)

	return &j
}

// Experimental.
func NewCodePipeline_Override(c CodePipeline, pipeline awscodepipeline.IPipeline, options *CodePipelineTargetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.CodePipeline",
		[]interface{}{pipeline, options},
		c,
	)
}

func (c *jsiiProxy_CodePipeline) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

