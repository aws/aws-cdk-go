package awsshield

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsshield/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Protection.
// Experimental.
type IProtectionRef interface {
	constructs.IConstruct
	// A reference to a Protection resource.
	// Experimental.
	ProtectionRef() *ProtectionReference
}

// The jsii proxy for IProtectionRef
type jsiiProxy_IProtectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProtectionRef) ProtectionRef() *ProtectionReference {
	var returns *ProtectionReference
	_jsii_.Get(
		j,
		"protectionRef",
		&returns,
	)
	return returns
}

