package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// This is a collection of ProjectInjectors assigned to this scope.
//
// It is keyed by constructUniqueId.  There can be only one ProjectInjector for a constructUniqueId.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   propertyInjectors := cdk.PropertyInjectors_Of(this)
//
type PropertyInjectors interface {
	// The scope attached to Injectors.
	Scope() constructs.IConstruct
	// Add a list of  IPropertyInjectors to this collection of PropertyInjectors.
	Add(propsInjectors ...IPropertyInjector)
	// Get the PropertyInjector that is registered to the Construct's uniqueId.
	//
	// Returns: - the IPropertyInjector for that construct uniqueId.
	For(uniqueId *string) IPropertyInjector
	// This returns a list of the Constructs that are supporting by this PropertyInjectors.
	//
	// Returns: a list of string showing the supported Constructs.
	SupportedClasses() *[]*string
}

// The jsii proxy struct for PropertyInjectors
type jsiiProxy_PropertyInjectors struct {
	_ byte // padding
}

func (j *jsiiProxy_PropertyInjectors) Scope() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}


// Return whether the given object has a PropertyInjectors property.
//
// We do attribute detection since we can't reliably use 'instanceof'.
func PropertyInjectors_HasPropertyInjectors(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePropertyInjectors_HasPropertyInjectorsParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.PropertyInjectors",
		"hasPropertyInjectors",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns the `PropertyInjectors` object associated with a construct scope.
//
// If `PropertyInjectors` object doesn't exist on this scope, then it creates one and attaches it to scope.
func PropertyInjectors_Of(scope constructs.IConstruct) PropertyInjectors {
	_init_.Initialize()

	if err := validatePropertyInjectors_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns PropertyInjectors

	_jsii_.StaticInvoke(
		"aws-cdk-lib.PropertyInjectors",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PropertyInjectors) Add(propsInjectors ...IPropertyInjector) {
	args := []interface{}{}
	for _, a := range propsInjectors {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"add",
		args,
	)
}

func (p *jsiiProxy_PropertyInjectors) For(uniqueId *string) IPropertyInjector {
	if err := p.validateForParameters(uniqueId); err != nil {
		panic(err)
	}
	var returns IPropertyInjector

	_jsii_.Invoke(
		p,
		"for",
		[]interface{}{uniqueId},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PropertyInjectors) SupportedClasses() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"supportedClasses",
		nil, // no parameters
		&returns,
	)

	return returns
}

