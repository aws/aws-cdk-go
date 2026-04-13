package previewawselasticacheevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.elasticache@CacheUpdateFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cacheUpdateFailed := awscdkmixinspreview.Events.NewCacheUpdateFailed()
//
// Experimental.
type CacheUpdateFailed interface {
}

// The jsii proxy struct for CacheUpdateFailed
type jsiiProxy_CacheUpdateFailed struct {
	_ byte // padding
}

// Experimental.
func NewCacheUpdateFailed() CacheUpdateFailed {
	_init_.Initialize()

	j := jsiiProxy_CacheUpdateFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheUpdateFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCacheUpdateFailed_Override(c CacheUpdateFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheUpdateFailed",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for Cache Update Failed.
// Experimental.
func CacheUpdateFailed_EventPattern(options *CacheUpdateFailed_CacheUpdateFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCacheUpdateFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheUpdateFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

