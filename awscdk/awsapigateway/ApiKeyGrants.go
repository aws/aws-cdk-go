package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway"
)

// Collection of grant methods for a IApiKeyRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var apiKeyRef IApiKeyRef
//
//   apiKeyGrants := awscdk.Aws_apigateway.ApiKeyGrants_FromApiKey(apiKeyRef)
//
type ApiKeyGrants interface {
	Resource() interfacesawsapigateway.IApiKeyRef
	// Permits the IAM principal all read operations through this key.
	Read(grantee awsiam.IGrantable) awsiam.Grant
	// Permits the IAM principal all read and write operations through this key.
	ReadWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Permits the IAM principal all write operations through this key.
	Write(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for ApiKeyGrants
type jsiiProxy_ApiKeyGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_ApiKeyGrants) Resource() interfacesawsapigateway.IApiKeyRef {
	var returns interfacesawsapigateway.IApiKeyRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for ApiKeyGrants.
func ApiKeyGrants_FromApiKey(resource interfacesawsapigateway.IApiKeyRef) ApiKeyGrants {
	_init_.Initialize()

	if err := validateApiKeyGrants_FromApiKeyParameters(resource); err != nil {
		panic(err)
	}
	var returns ApiKeyGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ApiKeyGrants",
		"fromApiKey",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiKeyGrants) Read(grantee awsiam.IGrantable) awsiam.Grant {
	if err := a.validateReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"read",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiKeyGrants) ReadWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := a.validateReadWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"readWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiKeyGrants) Write(grantee awsiam.IGrantable) awsiam.Grant {
	if err := a.validateWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"write",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

