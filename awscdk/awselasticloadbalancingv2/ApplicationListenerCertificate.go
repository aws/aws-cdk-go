package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Add certificates to a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationListener applicationListener
//   var listenerCertificate listenerCertificate
//
//   applicationListenerCertificate := awscdk.Aws_elasticloadbalancingv2.NewApplicationListenerCertificate(this, jsii.String("MyApplicationListenerCertificate"), &ApplicationListenerCertificateProps{
//   	Listener: applicationListener,
//
//   	// the properties below are optional
//   	Certificates: []iListenerCertificate{
//   		listenerCertificate,
//   	},
//   })
//
type ApplicationListenerCertificate interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ApplicationListenerCertificate
type jsiiProxy_ApplicationListenerCertificate struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApplicationListenerCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewApplicationListenerCertificate(scope constructs.Construct, id *string, props *ApplicationListenerCertificateProps) ApplicationListenerCertificate {
	_init_.Initialize()

	if err := validateNewApplicationListenerCertificateParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationListenerCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApplicationListenerCertificate_Override(a ApplicationListenerCertificate, scope constructs.Construct, id *string, props *ApplicationListenerCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerCertificate",
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
func ApplicationListenerCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationListenerCertificate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationListenerCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

