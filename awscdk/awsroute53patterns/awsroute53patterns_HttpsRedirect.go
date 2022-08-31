package awsroute53patterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53patterns/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Allows creating a domainA -> domainB redirect using CloudFront and S3.
//
// You can specify multiple domains to be redirected.
//
// Example:
//   patterns.NewHttpsRedirect(this, jsii.String("Redirect"), &httpsRedirectProps{
//   	recordNames: []*string{
//   		jsii.String("foo.example.com"),
//   	},
//   	targetDomain: jsii.String("bar.example.com"),
//   	zone: route53.hostedZone.fromHostedZoneAttributes(this, jsii.String("HostedZone"), &hostedZoneAttributes{
//   		hostedZoneId: jsii.String("ID"),
//   		zoneName: jsii.String("example.com"),
//   	}),
//   })
//
// Experimental.
type HttpsRedirect interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for HttpsRedirect
type jsiiProxy_HttpsRedirect struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_HttpsRedirect) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpsRedirect(scope constructs.Construct, id *string, props *HttpsRedirectProps) HttpsRedirect {
	_init_.Initialize()

	j := jsiiProxy_HttpsRedirect{}

	_jsii_.Create(
		"monocdk.aws_route53_patterns.HttpsRedirect",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpsRedirect_Override(h HttpsRedirect, scope constructs.Construct, id *string, props *HttpsRedirectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_patterns.HttpsRedirect",
		[]interface{}{scope, id, props},
		h,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func HttpsRedirect_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53_patterns.HttpsRedirect",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpsRedirect) OnPrepare() {
	_jsii_.InvokeVoid(
		h,
		"onPrepare",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HttpsRedirect) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (h *jsiiProxy_HttpsRedirect) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpsRedirect) Prepare() {
	_jsii_.InvokeVoid(
		h,
		"prepare",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HttpsRedirect) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
		"synthesize",
		[]interface{}{session},
	)
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

func (h *jsiiProxy_HttpsRedirect) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

