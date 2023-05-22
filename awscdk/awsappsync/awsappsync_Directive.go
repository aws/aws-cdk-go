package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Directives for types.
//
// i.e. @aws_iam or @aws_subscribe
//
// Example:
//   var api graphqlApi
//   var film interfaceType
//
//
//   api.AddSubscription(jsii.String("addedFilm"), appsync.NewField(&FieldOptions{
//   	ReturnType: film.Attribute(),
//   	Args: map[string]graphqlType{
//   		"id": appsync.*graphqlType_id(&BaseTypeOptions{
//   			"isRequired": jsii.Boolean(true),
//   		}),
//   	},
//   	Directives: []directive{
//   		appsync.*directive_Subscribe(jsii.String("addFilm")),
//   	},
//   }))
//
// Experimental.
type Directive interface {
	// The authorization type of this directive.
	// Experimental.
	Mode() AuthorizationType
	// the authorization modes for this intermediate type.
	// Experimental.
	Modes() *[]AuthorizationType
	// Experimental.
	SetModes(val *[]AuthorizationType)
	// Mutation fields for a subscription directive.
	// Experimental.
	MutationFields() *[]*string
	// Generate the directive statement.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Directive
type jsiiProxy_Directive struct {
	_ byte // padding
}

func (j *jsiiProxy_Directive) Mode() AuthorizationType {
	var returns AuthorizationType
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Directive) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Directive) MutationFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mutationFields",
		&returns,
	)
	return returns
}


func (j *jsiiProxy_Directive)SetModes(val *[]AuthorizationType) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

// Add the @aws_api_key directive.
// Experimental.
func Directive_ApiKey() Directive {
	_init_.Initialize()

	var returns Directive

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Directive",
		"apiKey",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Add the @aws_auth or @aws_cognito_user_pools directive.
// Experimental.
func Directive_Cognito(groups ...*string) Directive {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range groups {
		args = append(args, a)
	}

	var returns Directive

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Directive",
		"cognito",
		args,
		&returns,
	)

	return returns
}

// Add a custom directive.
// Experimental.
func Directive_Custom(statement *string) Directive {
	_init_.Initialize()

	if err := validateDirective_CustomParameters(statement); err != nil {
		panic(err)
	}
	var returns Directive

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Directive",
		"custom",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

// Add the @aws_iam directive.
// Experimental.
func Directive_Iam() Directive {
	_init_.Initialize()

	var returns Directive

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Directive",
		"iam",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Add the @aws_oidc directive.
// Experimental.
func Directive_Oidc() Directive {
	_init_.Initialize()

	var returns Directive

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Directive",
		"oidc",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Add the @aws_subscribe directive.
//
// Only use for top level Subscription type.
// Experimental.
func Directive_Subscribe(mutations ...*string) Directive {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range mutations {
		args = append(args, a)
	}

	var returns Directive

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Directive",
		"subscribe",
		args,
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Directive) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

