package awsroute53patterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53patterns/internal"
	"github.com/aws/constructs-go/constructs/v10"
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
type HttpsRedirect interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func HttpsRedirect_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

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

// Properties to configure an HTTPS Redirect.
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
type HttpsRedirectProps struct {
	// The redirect target fully qualified domain name (FQDN).
	//
	// An alias record
	// will be created that points to your CloudFront distribution. Root domain
	// or sub-domain can be supplied.
	TargetDomain *string `field:"required" json:"targetDomain" yaml:"targetDomain"`
	// Hosted zone of the domain which will be used to create alias record(s) from domain names in the hosted zone to the target domain.
	//
	// The hosted zone must
	// contain entries for the domain name(s) supplied through `recordNames` that
	// will redirect to the target domain.
	//
	// Domain names in the hosted zone can include a specific domain (example.com)
	// and its subdomains (acme.example.com, zenith.example.com).
	Zone awsroute53.IHostedZone `field:"required" json:"zone" yaml:"zone"`
	// The AWS Certificate Manager (ACM) certificate that will be associated with the CloudFront distribution that will be created.
	//
	// If provided, the certificate must be
	// stored in us-east-1 (N. Virginia)
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// The domain names that will redirect to `targetDomain`.
	RecordNames *[]*string `field:"optional" json:"recordNames" yaml:"recordNames"`
}

