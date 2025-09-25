package awsopensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityConfig.
// Experimental.
type ISecurityConfigRef interface {
	constructs.IConstruct
	// A reference to a SecurityConfig resource.
	// Experimental.
	SecurityConfigRef() *SecurityConfigReference
}

// The jsii proxy for ISecurityConfigRef
type jsiiProxy_ISecurityConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISecurityConfigRef) SecurityConfigRef() *SecurityConfigReference {
	var returns *SecurityConfigReference
	_jsii_.Get(
		j,
		"securityConfigRef",
		&returns,
	)
	return returns
}

