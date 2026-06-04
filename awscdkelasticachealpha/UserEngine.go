package awscdkelasticachealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkelasticachealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Engine type for ElastiCache users and user groups.
//
// Use the named static members for the engines currently supported by ElastiCache
// user/user-group resources. To target an engine not yet represented by a named
// instance, use `UserEngine.of(engineType)`.
//
// Example:
//   user := elasticache.NewIamUser(this, jsii.String("User"), &IamUserProps{
//   	// set user engine
//   	Engine: elasticache.UserEngine_REDIS(),
//
//   	// set user id
//   	UserId: jsii.String("my-user"),
//
//   	// set username
//   	UserName: jsii.String("my-user"),
//
//   	// set access string
//   	AccessControl: elasticache.AccessControl_FromAccessString(jsii.String("on ~* +@all")),
//   })
//
// Experimental.
type UserEngine interface {
	// The engine type, for example `'valkey'` or `'redis'`.
	//
	// Maps directly to the `Engine` property of `AWS::ElastiCache::User` and
	// `AWS::ElastiCache::UserGroup`.
	// Experimental.
	EngineType() *string
	// Returns the engine type as a string (for example, `'valkey'`).
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for UserEngine
type jsiiProxy_UserEngine struct {
	_ byte // padding
}

func (j *jsiiProxy_UserEngine) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}


// Create a new `UserEngine` with an arbitrary engine type.
// Experimental.
func UserEngine_Of(engineType *string) UserEngine {
	_init_.Initialize()

	if err := validateUserEngine_OfParameters(engineType); err != nil {
		panic(err)
	}
	var returns UserEngine

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-elasticache-alpha.UserEngine",
		"of",
		[]interface{}{engineType},
		&returns,
	)

	return returns
}

func UserEngine_REDIS() UserEngine {
	_init_.Initialize()
	var returns UserEngine
	_jsii_.StaticGet(
		"@aws-cdk/aws-elasticache-alpha.UserEngine",
		"REDIS",
		&returns,
	)
	return returns
}

func UserEngine_VALKEY() UserEngine {
	_init_.Initialize()
	var returns UserEngine
	_jsii_.StaticGet(
		"@aws-cdk/aws-elasticache-alpha.UserEngine",
		"VALKEY",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_UserEngine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

