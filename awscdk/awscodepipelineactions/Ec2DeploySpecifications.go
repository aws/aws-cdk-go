package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A deploy specifications for EC2 deploy action.
//
// Example:
//   sourceOutput := codepipeline.NewArtifact()
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &PipelineProps{
//   	PipelineType: codepipeline.PipelineType_V2,
//   })
//   deployAction := codepipeline_actions.NewEc2DeployAction(&Ec2DeployActionProps{
//   	ActionName: jsii.String("Ec2Deploy"),
//   	Input: sourceOutput,
//   	InstanceType: codepipeline_actions.Ec2InstanceType_EC2,
//   	InstanceTagKey: jsii.String("Name"),
//   	InstanceTagValue: jsii.String("MyInstance"),
//   	DeploySpecifications: codepipeline_actions.Ec2DeploySpecifications_Inline(&Ec2DeploySpecificationsInlineProps{
//   		TargetDirectory: jsii.String("/home/ec2-user/deploy"),
//   		PreScript: jsii.String("scripts/pre-deploy.sh"),
//   		PostScript: jsii.String("scripts/post-deploy.sh"),
//   	}),
//   })
//   deployStage := pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Deploy"),
//   	Actions: []IAction{
//   		deployAction,
//   	},
//   })
//
type Ec2DeploySpecifications interface {
	// The callback invoked when this deploy specifications is bound to an action.
	//
	// Returns: the action configurations.
	Bind(scope constructs.Construct) interface{}
}

// The jsii proxy struct for Ec2DeploySpecifications
type jsiiProxy_Ec2DeploySpecifications struct {
	_ byte // padding
}

func NewEc2DeploySpecifications_Override(e Ec2DeploySpecifications) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.Ec2DeploySpecifications",
		nil, // no parameters
		e,
	)
}

// Store deploy specifications as action configurations.
func Ec2DeploySpecifications_Inline(props *Ec2DeploySpecificationsInlineProps) Ec2DeploySpecifications {
	_init_.Initialize()

	if err := validateEc2DeploySpecifications_InlineParameters(props); err != nil {
		panic(err)
	}
	var returns Ec2DeploySpecifications

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.Ec2DeploySpecifications",
		"inline",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2DeploySpecifications) Bind(scope constructs.Construct) interface{} {
	if err := e.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

