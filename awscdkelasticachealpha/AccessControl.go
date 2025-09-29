package awscdkelasticachealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkelasticachealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Access control configuration for ElastiCache users.
//
// Example:
//   user := elasticache.NewIamUser(this, jsii.String("User"), &IamUserProps{
//   	// set user engine
//   	Engine: elasticache.UserEngine_REDIS,
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
type AccessControl interface {
	// The access string that defines user's permissions.
	// Experimental.
	AccessString() *string
}

// The jsii proxy struct for AccessControl
type jsiiProxy_AccessControl struct {
	_ byte // padding
}

func (j *jsiiProxy_AccessControl) AccessString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessString",
		&returns,
	)
	return returns
}


// Experimental.
func NewAccessControl_Override(a AccessControl) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-elasticache-alpha.AccessControl",
		nil, // no parameters
		a,
	)
}

// Create access control from an access string.
// Experimental.
func AccessControl_FromAccessString(accessString *string) AccessControl {
	_init_.Initialize()

	if err := validateAccessControl_FromAccessStringParameters(accessString); err != nil {
		panic(err)
	}
	var returns AccessControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-elasticache-alpha.AccessControl",
		"fromAccessString",
		[]interface{}{accessString},
		&returns,
	)

	return returns
}

