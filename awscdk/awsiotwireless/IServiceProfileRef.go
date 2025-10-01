package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceProfile.
// Experimental.
type IServiceProfileRef interface {
	constructs.IConstruct
	// A reference to a ServiceProfile resource.
	// Experimental.
	ServiceProfileRef() *ServiceProfileReference
}

// The jsii proxy for IServiceProfileRef
type jsiiProxy_IServiceProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServiceProfileRef) ServiceProfileRef() *ServiceProfileReference {
	var returns *ServiceProfileReference
	_jsii_.Get(
		j,
		"serviceProfileRef",
		&returns,
	)
	return returns
}

