package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an API Gateway authorizer.
type IAuthorizer interface {
	// The authorization type of this authorizer.
	AuthorizationType() AuthorizationType
	// The authorizer ID.
	AuthorizerId() *string
}

// The jsii proxy for IAuthorizer
type jsiiProxy_IAuthorizer struct {
	_ byte // padding
}

func (j *jsiiProxy_IAuthorizer) AuthorizationType() AuthorizationType {
	var returns AuthorizationType
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthorizer) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

