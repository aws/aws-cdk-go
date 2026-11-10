package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda"
)

// Collection of grant methods for a WebFunction endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var webFunctionEndpointRef IWebFunctionEndpointRef
//
//   webFunctionEndpointGrants := awscdk.Aws_lambda.WebFunctionEndpointGrants_FromWebFunctionEndpoint(webFunctionEndpointRef)
//
type WebFunctionEndpointGrants interface {
	// Grants permission to invoke the endpoint over HTTP.
	//
	// Only meaningful for an endpoint with `WebFunctionAuthType.IAM_AUTH`; requests to it must be
	// signed by the grantee.
	//
	// A role or user in this account gets the permission in its own policy. Any other grantee, such
	// as a service principal or another account, gets a statement in the resource-based policy of
	// the WebFunction the endpoint serves, which requires that function to be defined in this app.
	Invoke(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for WebFunctionEndpointGrants
type jsiiProxy_WebFunctionEndpointGrants struct {
	_ byte // padding
}

// Creates grants for a WebFunction endpoint.
func WebFunctionEndpointGrants_FromWebFunctionEndpoint(endpoint interfacesawslambda.IWebFunctionEndpointRef) WebFunctionEndpointGrants {
	_init_.Initialize()

	if err := validateWebFunctionEndpointGrants_FromWebFunctionEndpointParameters(endpoint); err != nil {
		panic(err)
	}
	var returns WebFunctionEndpointGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.WebFunctionEndpointGrants",
		"fromWebFunctionEndpoint",
		[]interface{}{endpoint},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebFunctionEndpointGrants) Invoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := w.validateInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		w,
		"invoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

