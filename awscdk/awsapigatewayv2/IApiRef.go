package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Api.
// Experimental.
type IApiRef interface {
	constructs.IConstruct
	// A reference to a Api resource.
	// Experimental.
	ApiRef() *ApiReference
}

// The jsii proxy for IApiRef
type jsiiProxy_IApiRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApiRef) ApiRef() *ApiReference {
	var returns *ApiReference
	_jsii_.Get(
		j,
		"apiRef",
		&returns,
	)
	return returns
}

