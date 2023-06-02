package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Aspects can be applied to CDK tree scopes and can operate on the tree before synthesis.
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
type Aspects interface {
	// The list of aspects which were directly applied on this scope.
	All() *[]IAspect
	// Adds an aspect to apply this scope before synthesis.
	Add(aspect IAspect)
}

// The jsii proxy struct for Aspects
type jsiiProxy_Aspects struct {
	_ byte // padding
}

func (j *jsiiProxy_Aspects) All() *[]IAspect {
	var returns *[]IAspect
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}


// Returns the `Aspects` object associated with a construct scope.
func Aspects_Of(scope constructs.IConstruct) Aspects {
	_init_.Initialize()

	if err := validateAspects_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns Aspects

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Aspects",
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

