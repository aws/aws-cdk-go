package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityGroupEgress.
// Experimental.
type ISecurityGroupEgressRef interface {
	constructs.IConstruct
	// A reference to a SecurityGroupEgress resource.
	// Experimental.
	SecurityGroupEgressRef() *SecurityGroupEgressReference
}

// The jsii proxy for ISecurityGroupEgressRef
type jsiiProxy_ISecurityGroupEgressRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISecurityGroupEgressRef) SecurityGroupEgressRef() *SecurityGroupEgressReference {
	var returns *SecurityGroupEgressReference
	_jsii_.Get(
		j,
		"securityGroupEgressRef",
		&returns,
	)
	return returns
}

