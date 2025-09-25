package awsshield

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsshield/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProtectionGroup.
// Experimental.
type IProtectionGroupRef interface {
	constructs.IConstruct
	// A reference to a ProtectionGroup resource.
	// Experimental.
	ProtectionGroupRef() *ProtectionGroupReference
}

// The jsii proxy for IProtectionGroupRef
type jsiiProxy_IProtectionGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProtectionGroupRef) ProtectionGroupRef() *ProtectionGroupReference {
	var returns *ProtectionGroupReference
	_jsii_.Get(
		j,
		"protectionGroupRef",
		&returns,
	)
	return returns
}

