package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Includes API for attaching annotations such as warning messages to constructs.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/constructs-go/constructs"
//
//   type myAspect struct {
//   }
//
//   func (this *myAspect) visit(node iConstruct) {
//   	if *node instanceof cdk.CfnResource && *node.CfnResourceType == "Foo::Bar" {
//   		this.error(*node, jsii.String("we do not want a Foo::Bar resource"))
//   	}
//   }
//
//   func (this *myAspect) error(node iConstruct, message *string) {
//   	cdk.Annotations_Of(*node).AddError(*message)
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
//   	cdk.NewCfnResource(stack, jsii.String("Foo"), &CfnResourceProps{
//   		Type: jsii.String("Foo::Bar"),
//   		Properties: map[string]interface{}{
//   			"Fred": jsii.String("Thud"),
//   		},
//   	})
//   	cdk.Aspects_Of(stack).Add(NewMyAspect())
//   	return this
//   }
//
type Annotations interface {
	// Acknowledge a warning. When a warning is acknowledged for a scope all warnings that match the id will be ignored.
	//
	// The acknowledgement will apply to all child scopes.
	//
	// Example:
	//   var myConstruct construct
	//
	//   awscdk.Annotations_Of(myConstruct).AcknowledgeWarning(jsii.String("SomeWarningId"), jsii.String("This warning can be ignored because..."))
	//
	AcknowledgeWarning(id *string, message *string)
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
	// Adds a warning metadata entry to this construct. Prefer using `addWarningV2`.
	//
	// The CLI will display the warning when an app is synthesized, or fail if run
	// in `--strict` mode.
	//
	// Warnings added by this call cannot be acknowledged. This will block users from
	// running in `--strict` mode until the deal with the warning, which makes it
	// effectively not very different from `addError`. Prefer using `addWarningV2` instead.
	AddWarning(message *string)
	// Adds an acknowledgeable warning metadata entry to this construct.
	//
	// The CLI will display the warning when an app is synthesized, or fail if run
	// in `--strict` mode.
	//
	// If the warning is acknowledged using `acknowledgeWarning()`, it will not be shown by
	// the CLI, and will not cause `--strict` mode to fail synthesis.
	//
	// Example:
	//   var myConstruct construct
	//
	//   awscdk.Annotations_Of(myConstruct).AddWarningV2(jsii.String("my-library:Construct.someWarning"), jsii.String("Some message explaining the warning"))
	//
	AddWarningV2(id *string, message *string)
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

func (a *jsiiProxy_Annotations) AcknowledgeWarning(id *string, message *string) {
	if err := a.validateAcknowledgeWarningParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"acknowledgeWarning",
		[]interface{}{id, message},
	)
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

func (a *jsiiProxy_Annotations) AddWarningV2(id *string, message *string) {
	if err := a.validateAddWarningV2Parameters(id, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addWarningV2",
		[]interface{}{id, message},
	)
}

