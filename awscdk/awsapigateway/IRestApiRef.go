package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RestApi.
// Experimental.
type IRestApiRef interface {
	constructs.IConstruct
	// A reference to a RestApi resource.
	// Experimental.
	RestApiRef() *RestApiReference
}

// The jsii proxy for IRestApiRef
type jsiiProxy_IRestApiRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRestApiRef) RestApiRef() *RestApiReference {
	var returns *RestApiReference
	_jsii_.Get(
		j,
		"restApiRef",
		&returns,
	)
	return returns
}

