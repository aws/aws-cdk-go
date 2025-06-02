package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An allow list receipt filter.
//
// Example:
//   ses.NewAllowListReceiptFilter(this, jsii.String("AllowList"), &AllowListReceiptFilterProps{
//   	Ips: []*string{
//   		jsii.String("10.0.0.0/16"),
//   		jsii.String("1.2.3.4/16"),
//   	},
//   })
//
type AllowListReceiptFilter interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AllowListReceiptFilter
type jsiiProxy_AllowListReceiptFilter struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AllowListReceiptFilter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewAllowListReceiptFilter(scope constructs.Construct, id *string, props *AllowListReceiptFilterProps) AllowListReceiptFilter {
	_init_.Initialize()

	if err := validateNewAllowListReceiptFilterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AllowListReceiptFilter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.AllowListReceiptFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAllowListReceiptFilter_Override(a AllowListReceiptFilter, scope constructs.Construct, id *string, props *AllowListReceiptFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.AllowListReceiptFilter",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AllowListReceiptFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAllowListReceiptFilter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.AllowListReceiptFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AllowListReceiptFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

