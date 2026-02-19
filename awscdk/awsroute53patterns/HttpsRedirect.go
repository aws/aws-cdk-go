package awsroute53patterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53patterns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Allows creating a domainA -> domainB redirect using CloudFront and S3.
//
// You can specify multiple domains to be redirected.
//
// Example:
//   patterns.NewHttpsRedirect(this, jsii.String("Redirect"), &HttpsRedirectProps{
//   	RecordNames: []*string{
//   		jsii.String("foo.example.com"),
//   	},
//   	TargetDomain: jsii.String("bar.example.com"),
//   	Zone: route53.HostedZone_FromHostedZoneAttributes(this, jsii.String("HostedZone"), &HostedZoneAttributes{
//   		HostedZoneId: jsii.String("ID"),
//   		ZoneName: jsii.String("example.com"),
//   	}),
//   })
//
type HttpsRedirect interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for HttpsRedirect
type jsiiProxy_HttpsRedirect struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_HttpsRedirect) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewHttpsRedirect(scope constructs.Construct, id *string, props *HttpsRedirectProps) HttpsRedirect {
	_init_.Initialize()

	if err := validateNewHttpsRedirectParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpsRedirect{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_patterns.HttpsRedirect",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewHttpsRedirect_Override(h HttpsRedirect, scope constructs.Construct, id *string, props *HttpsRedirectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_patterns.HttpsRedirect",
		[]interface{}{scope, id, props},
		h,
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
func HttpsRedirect_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHttpsRedirect_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53_patterns.HttpsRedirect",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpsRedirect) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpsRedirect) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		h,
		"with",
		args,
		&returns,
	)

	return returns
}

