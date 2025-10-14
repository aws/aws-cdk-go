package awselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityGroup.
// Experimental.
type ISecurityGroupRef interface {
	constructs.IConstruct
	// A reference to a SecurityGroup resource.
	// Experimental.
	SecurityGroupRef() *SecurityGroupReference
}

// The jsii proxy for ISecurityGroupRef
type jsiiProxy_ISecurityGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISecurityGroupRef) SecurityGroupRef() *SecurityGroupReference {
	var returns *SecurityGroupReference
	_jsii_.Get(
		j,
		"securityGroupRef",
		&returns,
	)
	return returns
}

