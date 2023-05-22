package awsamplify

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Custom rewrite/redirect rule for an Amplify App.
//
// Example:
//   var amplifyApp app
//
//   amplifyApp.AddCustomRule(map[string]interface{}{
//   	"source": jsii.String("/docs/specific-filename.html"),
//   	"target": jsii.String("/documents/different-filename.html"),
//   	"status": amplify.RedirectStatus_TEMPORARY_REDIRECT,
//   })
//
// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
//
// Experimental.
type CustomRule interface {
	// The condition for a URL rewrite or redirect rule, e.g. country code.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Experimental.
	Condition() *string
	// The source pattern for a URL rewrite or redirect rule.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Experimental.
	Source() *string
	// The status code for a URL rewrite or redirect rule.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Experimental.
	Status() RedirectStatus
	// The target pattern for a URL rewrite or redirect rule.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Experimental.
	Target() *string
}

// The jsii proxy struct for CustomRule
type jsiiProxy_CustomRule struct {
	_ byte // padding
}

func (j *jsiiProxy_CustomRule) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) Status() RedirectStatus {
	var returns RedirectStatus
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomRule(options *CustomRuleOptions) CustomRule {
	_init_.Initialize()

	if err := validateNewCustomRuleParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomRule{}

	_jsii_.Create(
		"monocdk.aws_amplify.CustomRule",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomRule_Override(c CustomRule, options *CustomRuleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.CustomRule",
		[]interface{}{options},
		c,
	)
}

func CustomRule_SINGLE_PAGE_APPLICATION_REDIRECT() CustomRule {
	_init_.Initialize()
	var returns CustomRule
	_jsii_.StaticGet(
		"monocdk.aws_amplify.CustomRule",
		"SINGLE_PAGE_APPLICATION_REDIRECT",
		&returns,
	)
	return returns
}

