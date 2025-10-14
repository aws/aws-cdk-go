package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiCache.
// Experimental.
type IApiCacheRef interface {
	constructs.IConstruct
	// A reference to a ApiCache resource.
	// Experimental.
	ApiCacheRef() *ApiCacheReference
}

// The jsii proxy for IApiCacheRef
type jsiiProxy_IApiCacheRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApiCacheRef) ApiCacheRef() *ApiCacheReference {
	var returns *ApiCacheReference
	_jsii_.Get(
		j,
		"apiCacheRef",
		&returns,
	)
	return returns
}

