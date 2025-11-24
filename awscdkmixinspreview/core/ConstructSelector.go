package core

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Selects constructs from a construct tree based on various criteria.
//
// Example:
//   // Apply to all constructs in a scope
//   awscdkmixinspreview.Mixins_Of(scope).Apply(NewEncryptionAtRest())
//
//   // Apply to specific resource types
//   awscdkmixinspreview.Mixins_Of(scope, awscdkmixinspreview.ConstructSelector_ResourcesOfType(s3.CfnBucket)).Apply(NewEncryptionAtRest())
//
//   // Apply to constructs matching a pattern
//   awscdkmixinspreview.Mixins_Of(scope, awscdkmixinspreview.ConstructSelector_ById(/.*-prod-.*/)).Apply(NewProductionSecurityMixin())
//
// Experimental.
type ConstructSelector interface {
	// Selects constructs from the given scope based on the selector's criteria.
	// Experimental.
	Select(scope constructs.IConstruct) *[]constructs.IConstruct
}

// The jsii proxy struct for ConstructSelector
type jsiiProxy_ConstructSelector struct {
	_ byte // padding
}

// Experimental.
func NewConstructSelector_Override(c ConstructSelector) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		nil, // no parameters
		c,
	)
}

// Selects all constructs in the tree.
// Experimental.
func ConstructSelector_All() ConstructSelector {
	_init_.Initialize()

	var returns ConstructSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Selects constructs whose IDs match a pattern.
// Experimental.
func ConstructSelector_ById(pattern interface{}) ConstructSelector {
	_init_.Initialize()

	if err := validateConstructSelector_ByIdParameters(pattern); err != nil {
		panic(err)
	}
	var returns ConstructSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		"byId",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Selects CfnResource constructs or the default CfnResource child.
// Experimental.
func ConstructSelector_CfnResource() ConstructSelector {
	_init_.Initialize()

	var returns ConstructSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		"cfnResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Selects constructs of a specific type.
// Experimental.
func ConstructSelector_ResourcesOfType(type_ interface{}) ConstructSelector {
	_init_.Initialize()

	if err := validateConstructSelector_ResourcesOfTypeParameters(type_); err != nil {
		panic(err)
	}
	var returns ConstructSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		"resourcesOfType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructSelector) Select(scope constructs.IConstruct) *[]constructs.IConstruct {
	if err := c.validateSelectParameters(scope); err != nil {
		panic(err)
	}
	var returns *[]constructs.IConstruct

	_jsii_.Invoke(
		c,
		"select",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

