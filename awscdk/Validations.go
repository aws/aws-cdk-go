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
	// Register one or more validation plugins that will be executed during synthesis.
	//
	// Plugins can only be registered within a Stage or App scope.
	// If any plugin reports a violation, synthesis will be interrupted and the
	// report displayed to the user.
	AddPlugins(plugins ...IPolicyValidationPlugin)
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

