package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Converts between Mixins and Aspects.
//
// Since Mixins and Aspects are both implementations of the visitor pattern,
// they can be converted from each other. Mixins are applied immediately (imperative),
// while Aspects are applied during synthesis (declarative).
type Shims interface {
}

// The jsii proxy struct for Shims
type jsiiProxy_Shims struct {
	_ byte // padding
}

// Wraps a Mixin as an Aspect.
//
// The resulting Aspect defers the Mixin's application to the synthesis phase.
// The Mixin's `supports()` method is used to filter which constructs are visited.
func Shims_AsAspect(mixin constructs.IMixin) IAspect {
	_init_.Initialize()

	if err := validateShims_AsAspectParameters(mixin); err != nil {
		panic(err)
	}
	var returns IAspect

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Shims",
		"asAspect",
		[]interface{}{mixin},
		&returns,
	)

	return returns
}

// Wraps an Aspect as a Mixin.
//
// The resulting Mixin applies the Aspect's `visit()` immediately to every node.
// The Aspect is applied to all constructs since Aspects don't have a `supports()` filter.
func Shims_AsMixin(aspect IAspect) constructs.IMixin {
	_init_.Initialize()

	if err := validateShims_AsMixinParameters(aspect); err != nil {
		panic(err)
	}
	var returns constructs.IMixin

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Shims",
		"asMixin",
		[]interface{}{aspect},
		&returns,
	)

	return returns
}

