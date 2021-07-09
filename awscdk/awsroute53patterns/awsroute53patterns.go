package awsroute53patterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53patterns/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Allows creating a domainA -> domainB redirect using CloudFront and S3.
//
// You can specify multiple domains to be redirected.
// Experimental.
type HttpsRedirect interface {
	awscdk.Construct
	Node() awscdk.ConstructNode
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (h *jsiiProxy_HttpsRedirect) OnPrepare() {
	_jsii_.InvokeVoid(
		h,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (h *jsiiProxy_HttpsRedirect) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (h *jsiiProxy_HttpsRedirect) Prepare() {
	_jsii_.InvokeVoid(
		h,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (h *jsiiProxy_HttpsRedirect) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Properties to configure an HTTPS Redirect.
// Experimental.
type HttpsRedirectProps struct {
	// The redirect target fully qualified domain name (FQDN).
	//
	// An alias record
	// will be created that points to your CloudFront distribution. Root domain
	// or sub-domain can be supplied.
	// Experimental.
	TargetDomain *string `json:"targetDomain"`
	// Hosted zone of the domain which will be used to create alias record(s) from domain names in the hosted zone to the target domain.
	//
	// The hosted zone must
	// contain entries for the domain name(s) supplied through `recordNames` that
	// will redirect to the target domain.
	//
	// Domain names in the hosted zone can include a specific domain (example.com)
	// and its subdomains (acme.example.com, zenith.example.com).
	// Experimental.
	Zone awsroute53.IHostedZone `json:"zone"`
	// The AWS Certificate Manager (ACM) certificate that will be associated with the CloudFront distribution that will be created.
	//
	// If provided, the certificate must be
	// stored in us-east-1 (N. Virginia)
	// Experimental.
	Certificate awscertificatemanager.ICertificate `json:"certificate"`
	// The domain names that will redirect to `targetDomain`.
	// Experimental.
	RecordNames *[]*string `json:"recordNames"`
}

