package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
)

// Allows the pipeline to be used as an EventBridge rule target.
//
// Example:
//   // A pipeline being used as a target for a CloudWatch event rule.
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline pipeline
//
//
//   // kick off the pipeline every day
//   rule := events.NewRule(this, jsii.String("Daily"), &RuleProps{
//   	Schedule: events.Schedule_Rate(awscdk.Duration_Days(jsii.Number(1))),
//   })
//   rule.AddTarget(targets.NewCodePipeline(pipeline))
//
type CodePipeline interface {
	awsevents.IRuleTarget
	// Returns the rule target specification.
	//
	// NOTE: Do not use the various `inputXxx` options. They can be set in a call to `addTarget`.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for CodePipeline
type jsiiProxy_CodePipeline struct {
	internal.Type__awseventsIRuleTarget
}

func NewCodePipeline(pipeline awscodepipeline.IPipeline, options *CodePipelineTargetOptions) CodePipeline {
	_init_.Initialize()

	if err := validateNewCodePipelineParameters(pipeline, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodePipeline{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.CodePipeline",
		[]interface{}{pipeline, options},
		&j,
	)

	return &j
}

func NewCodePipeline_Override(c CodePipeline, pipeline awscodepipeline.IPipeline, options *CodePipelineTargetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.CodePipeline",
		[]interface{}{pipeline, options},
		c,
	)
}

func (c *jsiiProxy_CodePipeline) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	if err := c.validateBindParameters(_rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

