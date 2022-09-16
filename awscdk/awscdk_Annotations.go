// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Includes API for attaching annotations such as warning messages to constructs.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/constructs-go/constructs"
//
//   type myAspect struct {
//   }
//
//   func (this *myAspect) visit(node iConstruct) {
//   	if *node instanceof cdk.cfnResource && *node.cfnResourceType == "Foo::Bar" {
//   		this.error(*node, jsii.String("we do not want a Foo::Bar resource"))
//   	}
//   }
//
//   func (this *myAspect) error(node iConstruct, message *string) {
//   	cdk.annotations.of(*node).addError(*message)
//   }
//
//   type myStack struct {
//   	stack
//   }
//
//   func newMyStack(scope construct, id *string) *myStack {
//   	this := &myStack{}
//   	cdk.NewStack_Override(this, scope, id)
//
//   	stack := cdk.NewStack()
//   	cdk.NewCfnResource(stack, jsii.String("Foo"), &cfnResourceProps{
//   		type: jsii.String("Foo::Bar"),
//   		properties: map[string]interface{}{
//   			"Fred": jsii.String("Thud"),
//   		},
//   	})
//   	cdk.aspects.of(stack).add(NewMyAspect())
//   	return this
//   }
//
type Annotations interface {
	// Adds a deprecation warning for a specific API.
	//
	// Deprecations will be added only once per construct as a warning and will be
	// deduplicated based on the `api`.
	//
	// If the environment variable `CDK_BLOCK_DEPRECATIONS` is set, this method
	// will throw an error instead with the deprecation message.
	AddDeprecation(api *string, message *string)
	// Adds an { "error": <message> } metadata entry to this construct.
	//
	// The toolkit will fail deployment of any stack that has errors reported against it.
	AddError(message *string)
	// Adds an info metadata entry to this construct.
	//
	// The CLI will display the info message when apps are synthesized.
	AddInfo(message *string)
	// Adds a warning metadata entry to this construct.
	//
	// The CLI will display the warning when an app is synthesized, or fail if run
	// in --strict mode.
	AddWarning(message *string)
}

// The jsii proxy struct for Annotations
type jsiiProxy_Annotations struct {
	_ byte // padding
}

// Returns the annotations API for a construct scope.
func Annotations_Of(scope constructs.IConstruct) Annotations {
	_init_.Initialize()

	if err := validateAnnotations_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns Annotations

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Annotations",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Annotations) AddDeprecation(api *string, message *string) {
	if err := a.validateAddDeprecationParameters(api, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addDeprecation",
		[]interface{}{api, message},
	)
}

func (a *jsiiProxy_Annotations) AddError(message *string) {
	if err := a.validateAddErrorParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addError",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_Annotations) AddInfo(message *string) {
	if err := a.validateAddInfoParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addInfo",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_Annotations) AddWarning(message *string) {
	if err := a.validateAddWarningParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addWarning",
		[]interface{}{message},
	)
}

