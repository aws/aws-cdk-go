package awselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityGroupIngress.
// Experimental.
type ISecurityGroupIngressRef interface {
	constructs.IConstruct
	// A reference to a SecurityGroupIngress resource.
	// Experimental.
	SecurityGroupIngressRef() *SecurityGroupIngressReference
}

// The jsii proxy for ISecurityGroupIngressRef
type jsiiProxy_ISecurityGroupIngressRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISecurityGroupIngressRef) SecurityGroupIngressRef() *SecurityGroupIngressReference {
	var returns *SecurityGroupIngressReference
	_jsii_.Get(
		j,
		"securityGroupIngressRef",
		&returns,
	)
	return returns
}

