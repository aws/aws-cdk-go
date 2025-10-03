package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityGroup.
// Experimental.
type ISecurityGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISecurityGroupRef
type jsiiProxy_ISecurityGroupRef struct {
	internal.Type__constructsIConstruct
}

