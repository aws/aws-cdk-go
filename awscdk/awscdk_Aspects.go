// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Aspects can be applied to CDK tree scopes and can operate on the tree before synthesis.
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
// Experimental.
type Aspects interface {
	// The list of aspects which were directly applied on this scope.
	// Experimental.
	Aspects() *[]IAspect
	// Adds an aspect to apply this scope before synthesis.
	// Experimental.
	Add(aspect IAspect)
}

// The jsii proxy struct for Aspects
type jsiiProxy_Aspects struct {
	_ byte // padding
}

func (j *jsiiProxy_Aspects) Aspects() *[]IAspect {
	var returns *[]IAspect
	_jsii_.Get(
		j,
		"aspects",
		&returns,
	)
	return returns
}


// Returns the `Aspects` object associated with a construct scope.
// Experimental.
func Aspects_Of(scope IConstruct) Aspects {
	_init_.Initialize()

	if err := validateAspects_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns Aspects

	_jsii_.StaticInvoke(
		"monocdk.Aspects",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Aspects) Add(aspect IAspect) {
	if err := a.validateAddParameters(aspect); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"add",
		[]interface{}{aspect},
	)
}

