package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Guradrail settings for BedrockInvokeModel.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   model := bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_AMAZON_TITAN_TEXT_G1_EXPRESS_V1())
//
//   task := tasks.NewBedrockInvokeModel(this, jsii.String("Prompt Model with guardrail"), &BedrockInvokeModelProps{
//   	Model: Model,
//   	Body: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"inputText": jsii.String("Generate a list of five first names."),
//   		"textGenerationConfig": map[string]*f64{
//   			"maxTokenCount": jsii.Number(100),
//   			"temperature": jsii.Number(1),
//   		},
//   	}),
//   	Guardrail: tasks.Guardrail_Enable(jsii.String("guardrailId"), jsii.Number(1)),
//   	ResultSelector: map[string]interface{}{
//   		"names": sfn.JsonPath_stringAt(jsii.String("$.Body.results[0].outputText")),
//   	},
//   })
//
type Guardrail interface {
	// The identitifier of guardrail.
	GuardrailIdentifier() *string
	// The version of guardrail.
	GuardrailVersion() *string
}

// The jsii proxy struct for Guardrail
type jsiiProxy_Guardrail struct {
	_ byte // padding
}

func (j *jsiiProxy_Guardrail) GuardrailIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Guardrail) GuardrailVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailVersion",
		&returns,
	)
	return returns
}


// Enable guardrail.
func Guardrail_Enable(identifier *string, version *float64) Guardrail {
	_init_.Initialize()

	if err := validateGuardrail_EnableParameters(identifier, version); err != nil {
		panic(err)
	}
	var returns Guardrail

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.Guardrail",
		"enable",
		[]interface{}{identifier, version},
		&returns,
	)

	return returns
}

// Enable guardrail with DRAFT version.
func Guardrail_EnableDraft(identifier *string) Guardrail {
	_init_.Initialize()

	if err := validateGuardrail_EnableDraftParameters(identifier); err != nil {
		panic(err)
	}
	var returns Guardrail

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.Guardrail",
		"enableDraft",
		[]interface{}{identifier},
		&returns,
	)

	return returns
}

