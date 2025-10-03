package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserGroup.
// Experimental.
type IUserGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserGroupRef
type jsiiProxy_IUserGroupRef struct {
	internal.Type__constructsIConstruct
}

