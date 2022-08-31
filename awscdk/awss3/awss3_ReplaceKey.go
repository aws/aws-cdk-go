package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &bucketProps{
//   	websiteRoutingRules: []routingRule{
//   		&routingRule{
//   			hostName: jsii.String("www.example.com"),
//   			httpRedirectCode: jsii.String("302"),
//   			protocol: s3.redirectProtocol_HTTPS,
//   			replaceKey: s3.replaceKey.prefixWith(jsii.String("test/")),
//   			condition: &routingRuleCondition{
//   				httpErrorCodeReturnedEquals: jsii.String("200"),
//   				keyPrefixEquals: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type ReplaceKey interface {
	// Experimental.
	PrefixWithKey() *string
	// Experimental.
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
// Experimental.
func ReplaceKey_PrefixWith(keyReplacement *string) ReplaceKey {
	_init_.Initialize()

	if err := validateReplaceKey_PrefixWithParameters(keyReplacement); err != nil {
		panic(err)
	}
	var returns ReplaceKey

	_jsii_.StaticInvoke(
		"monocdk.aws_s3.ReplaceKey",
		"prefixWith",
		[]interface{}{keyReplacement},
		&returns,
	)

	return returns
}

// The specific object key to use in the redirect request.
// Experimental.
func ReplaceKey_With(keyReplacement *string) ReplaceKey {
	_init_.Initialize()

	if err := validateReplaceKey_WithParameters(keyReplacement); err != nil {
		panic(err)
	}
	var returns ReplaceKey

	_jsii_.StaticInvoke(
		"monocdk.aws_s3.ReplaceKey",
		"with",
		[]interface{}{keyReplacement},
		&returns,
	)

	return returns
}

