package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Identity.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var user User
//
//
//   identity := ses.NewEmailIdentity(this, jsii.String("Identity"), &EmailIdentityProps{
//   	Identity: ses.Identity_Domain(jsii.String("cdk.dev")),
//   })
//
//   identity.grantSendEmail(user)
//
type Identity interface {
	// The hosted zone associated with this identity.
	// Default: - no hosted zone is associated and no records are created.
	//
	HostedZone() awsroute53.IPublicHostedZone
	// The value of the identity.
	Value() *string
}

// The jsii proxy struct for Identity
type jsiiProxy_Identity struct {
	_ byte // padding
}

func (j *jsiiProxy_Identity) HostedZone() awsroute53.IPublicHostedZone {
	var returns awsroute53.IPublicHostedZone
	_jsii_.Get(
		j,
		"hostedZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Identity) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewIdentity_Override(i Identity) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.Identity",
		nil, // no parameters
		i,
	)
}

// Verify a domain name.
//
// DKIM records will have to be added manually to complete the verification
// process.
func Identity_Domain(domain *string) Identity {
	_init_.Initialize()

	if err := validateIdentity_DomainParameters(domain); err != nil {
		panic(err)
	}
	var returns Identity

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.Identity",
		"domain",
		[]interface{}{domain},
		&returns,
	)

	return returns
}

// Verify an email address.
//
// To complete the verification process look for an email from
// no-reply-aws@amazon.com, open it and click the link.
func Identity_Email(email *string) Identity {
	_init_.Initialize()

	if err := validateIdentity_EmailParameters(email); err != nil {
		panic(err)
	}
	var returns Identity

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.Identity",
		"email",
		[]interface{}{email},
		&returns,
	)

	return returns
}

// Verify a public hosted zone.
//
// DKIM and MAIL FROM records will be added automatically to the hosted
// zone.
func Identity_PublicHostedZone(hostedZone awsroute53.IPublicHostedZone) Identity {
	_init_.Initialize()

	if err := validateIdentity_PublicHostedZoneParameters(hostedZone); err != nil {
		panic(err)
	}
	var returns Identity

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.Identity",
		"publicHostedZone",
		[]interface{}{hostedZone},
		&returns,
	)

	return returns
}

