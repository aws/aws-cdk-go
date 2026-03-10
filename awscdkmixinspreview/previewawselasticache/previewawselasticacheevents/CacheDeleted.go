package previewawselasticacheevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.elasticache@CacheDeleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cacheDeleted := awscdkmixinspreview.Events.NewCacheDeleted()
//
// Experimental.
type CacheDeleted interface {
}

// The jsii proxy struct for CacheDeleted
type jsiiProxy_CacheDeleted struct {
	_ byte // padding
}

// Experimental.
func NewCacheDeleted() CacheDeleted {
	_init_.Initialize()

	j := jsiiProxy_CacheDeleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheDeleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCacheDeleted_Override(c CacheDeleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheDeleted",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for Cache Deleted.
// Experimental.
func CacheDeleted_CacheDeletedPattern(options *CacheDeleted_CacheDeletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCacheDeleted_CacheDeletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheDeleted",
		"cacheDeletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

