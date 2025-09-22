package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityProfile.
// Experimental.
type ISecurityProfileRef interface {
	constructs.IConstruct
	// A reference to a SecurityProfile resource.
	// Experimental.
	SecurityProfileRef() *SecurityProfileReference
}

// The jsii proxy for ISecurityProfileRef
type jsiiProxy_ISecurityProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISecurityProfileRef) SecurityProfileRef() *SecurityProfileReference {
	var returns *SecurityProfileReference
	_jsii_.Get(
		j,
		"securityProfileRef",
		&returns,
	)
	return returns
}

