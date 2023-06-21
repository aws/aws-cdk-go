package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// The identity to use for DKIM.
//
// Example:
//   var myHostedZone iPublicHostedZone
//
//
//   ses.NewEmailIdentity(this, jsii.String("Identity"), &EmailIdentityProps{
//   	Identity: ses.Identity_PublicHostedZone(myHostedZone),
//   	DkimIdentity: ses.DkimIdentity_ByoDkim(&ByoDkimOptions{
//   		PrivateKey: awscdk.SecretValue_SecretsManager(jsii.String("dkim-private-key")),
//   		PublicKey: jsii.String("...base64-encoded-public-key..."),
//   		Selector: jsii.String("selector"),
//   	}),
//   })
//
type DkimIdentity interface {
	// Binds this DKIM identity to the email identity.
	Bind(emailIdentity EmailIdentity, hostedZone awsroute53.IPublicHostedZone) *DkimIdentityConfig
}

// The jsii proxy struct for DkimIdentity
type jsiiProxy_DkimIdentity struct {
	_ byte // padding
}

func NewDkimIdentity_Override(d DkimIdentity) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.DkimIdentity",
		nil, // no parameters
		d,
	)
}

// Bring Your Own DKIM.
// See: https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-bring-your-own.html
//
func DkimIdentity_ByoDkim(options *ByoDkimOptions) DkimIdentity {
	_init_.Initialize()

	if err := validateDkimIdentity_ByoDkimParameters(options); err != nil {
		panic(err)
	}
	var returns DkimIdentity

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.DkimIdentity",
		"byoDkim",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Easy DKIM.
// See: https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-easy.html
//
func DkimIdentity_EasyDkim(signingKeyLength EasyDkimSigningKeyLength) DkimIdentity {
	_init_.Initialize()

	var returns DkimIdentity

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.DkimIdentity",
		"easyDkim",
		[]interface{}{signingKeyLength},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DkimIdentity) Bind(emailIdentity EmailIdentity, hostedZone awsroute53.IPublicHostedZone) *DkimIdentityConfig {
	if err := d.validateBindParameters(emailIdentity); err != nil {
		panic(err)
	}
	var returns *DkimIdentityConfig

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{emailIdentity, hostedZone},
		&returns,
	)

	return returns
}

