package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiCache.
// Experimental.
type IApiCacheRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApiCacheRef
type jsiiProxy_IApiCacheRef struct {
	internal.Type__constructsIConstruct
}

