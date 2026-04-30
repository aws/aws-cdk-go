package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Manages validations for CDK constructs.
//
// Example:
//   var myApp App
//   var plugin IPolicyValidationPlugin
//
//   awscdk.Validations_Of(myApp).AddPlugins(plugin)
//
type Validations interface {
	// Acknowledge one or more rules, suppressing them from validation output.
	//
	// Acknowledgments are recorded to construct metadata so that downstream
	// plugins (e.g. CDK Nag) can read them for audit trails.
	//
	// Currently only annotation warnings can be suppressed. Annotation errors
	// are not yet acknowledgeable.
	//
	// If an ID has no well-known prefix, it is assumed to be an annotation rule
	// for backwards compatibility.
	Acknowledge(rules ...*Acknowledgment)
	// Adds an error metadata entry to this construct.
	//
	// Synthesis will be interrupted when errors are reported.
	//
	// Note: Annotation errors are not currently acknowledgeable. The ID is
	// recorded for identification purposes but `acknowledge()` will not
	// suppress errors added via this method.
	AddError(id *string, message *string)
	// Register one or more validation plugins that will be executed during synthesis.
	//
	// Plugins can only be registered within a Stage or App scope.
	// If any plugin reports a violation, synthesis will be interrupted and the
	// report displayed to the user.
	AddPlugins(plugins ...IPolicyValidationPlugin)
	// Adds a warning metadata entry to this construct that can be acknowledged.
	//
	// The CLI will display the warning when an app is synthesized, or fail if run
	// in `--strict` mode.
	//
	// The ID will be stored with the `annotation` prefix (e.g. `annotation::MyWarning`).
	// Use this prefixed ID when calling `acknowledge()` to suppress the warning.
	AddWarning(id *string, message *string)
}

// The jsii proxy struct for Validations
type jsiiProxy_Validations struct {
	_ byte // padding
}

// Returns the Validations for the given construct scope.
func Validations_Of(scope constructs.IConstruct) Validations {
	_init_.Initialize()

	if err := validateValidations_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns Validations

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Validations",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func Validations_ACKNOWLEDGED_RULES_METADATA_KEY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.Validations",
		"ACKNOWLEDGED_RULES_METADATA_KEY",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_Validations) Acknowledge(rules ...*Acknowledgment) {
	if err := v.validateAcknowledgeParameters(&rules); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range rules {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"acknowledge",
		args,
	)
}

func (v *jsiiProxy_Validations) AddError(id *string, message *string) {
	if err := v.validateAddErrorParameters(id, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addError",
		[]interface{}{id, message},
	)
}

func (v *jsiiProxy_Validations) AddPlugins(plugins ...IPolicyValidationPlugin) {
	args := []interface{}{}
	for _, a := range plugins {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"addPlugins",
		args,
	)
}

func (v *jsiiProxy_Validations) AddWarning(id *string, message *string) {
	if err := v.validateAddWarningParameters(id, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addWarning",
		[]interface{}{id, message},
	)
}

