package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServerlessCache.
// Experimental.
type IServerlessCacheRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServerlessCacheRef
type jsiiProxy_IServerlessCacheRef struct {
	internal.Type__constructsIConstruct
}

