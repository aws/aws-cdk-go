package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The mode of evaluation for the rule.
//
// Example:
//   var fn Function
//   var samplePolicyText string
//
//
//   config.NewManagedRule(this, jsii.String("ManagedRule"), &ManagedRuleProps{
//   	Identifier: config.ManagedRuleIdentifiers_API_GW_XRAY_ENABLED(),
//   	EvaluationModes: config.EvaluationMode_DETECTIVE_AND_PROACTIVE(),
//   })
//
//   config.NewCustomRule(this, jsii.String("CustomRule"), &CustomRuleProps{
//   	LambdaFunction: fn,
//   	EvaluationModes: config.EvaluationMode_PROACTIVE(),
//   })
//
//   config.NewCustomPolicy(this, jsii.String("CustomPolicy"), &CustomPolicyProps{
//   	PolicyText: samplePolicyText,
//   	EvaluationModes: config.EvaluationMode_DETECTIVE(),
//   })
//
type EvaluationMode interface {
	// The modes of evaluation for the rule.
	Modes() *[]*string
}

// The jsii proxy struct for EvaluationMode
type jsiiProxy_EvaluationMode struct {
	_ byte // padding
}

func (j *jsiiProxy_EvaluationMode) Modes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}


func NewEvaluationMode(modes *[]*string) EvaluationMode {
	_init_.Initialize()

	if err := validateNewEvaluationModeParameters(modes); err != nil {
		panic(err)
	}
	j := jsiiProxy_EvaluationMode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_config.EvaluationMode",
		[]interface{}{modes},
		&j,
	)

	return &j
}

func NewEvaluationMode_Override(e EvaluationMode, modes *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_config.EvaluationMode",
		[]interface{}{modes},
		e,
	)
}

func EvaluationMode_DETECTIVE() EvaluationMode {
	_init_.Initialize()
	var returns EvaluationMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.EvaluationMode",
		"DETECTIVE",
		&returns,
	)
	return returns
}

func EvaluationMode_DETECTIVE_AND_PROACTIVE() EvaluationMode {
	_init_.Initialize()
	var returns EvaluationMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.EvaluationMode",
		"DETECTIVE_AND_PROACTIVE",
		&returns,
	)
	return returns
}

func EvaluationMode_PROACTIVE() EvaluationMode {
	_init_.Initialize()
	var returns EvaluationMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.EvaluationMode",
		"PROACTIVE",
		&returns,
	)
	return returns
}

