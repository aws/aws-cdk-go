package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &BucketProps{
//   	WebsiteRoutingRules: []routingRule{
//   		&routingRule{
//   			HostName: jsii.String("www.example.com"),
//   			HttpRedirectCode: jsii.String("302"),
//   			Protocol: s3.RedirectProtocol_HTTPS,
//   			ReplaceKey: s3.ReplaceKey_PrefixWith(jsii.String("test/")),
//   			Condition: &RoutingRuleCondition{
//   				HttpErrorCodeReturnedEquals: jsii.String("200"),
//   				KeyPrefixEquals: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   })
//
type ReplaceKey interface {
	PrefixWithKey() *string
	WithKey() *string
}

// The jsii proxy struct for ReplaceKey
type jsiiProxy_ReplaceKey struct {
	_ byte // padding
}

func (j *jsiiProxy_ReplaceKey) PrefixWithKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixWithKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplaceKey) WithKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"withKey",
		&returns,
	)
	return returns
}


// The object key prefix to use in the redirect request.
func ReplaceKey_PrefixWith(keyReplacement *string) ReplaceKey {
	_init_.Initialize()

	if err := validateReplaceKey_PrefixWithParameters(keyReplacement); err != nil {
		panic(err)
	}
	var returns ReplaceKey

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.ReplaceKey",
		"prefixWith",
		[]interface{}{keyReplacement},
		&returns,
	)

	return returns
}

// The specific object key to use in the redirect request.
func ReplaceKey_With(keyReplacement *string) ReplaceKey {
	_init_.Initialize()

	if err := validateReplaceKey_WithParameters(keyReplacement); err != nil {
		panic(err)
	}
	var returns ReplaceKey

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.ReplaceKey",
		"with",
		[]interface{}{keyReplacement},
		&returns,
	)

	return returns
}

