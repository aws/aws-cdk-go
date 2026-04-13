package previewawselasticacheevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.elasticache@CacheUpdated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cacheUpdated := awscdkmixinspreview.Events.NewCacheUpdated()
//
// Experimental.
type CacheUpdated interface {
}

// The jsii proxy struct for CacheUpdated
type jsiiProxy_CacheUpdated struct {
	_ byte // padding
}

// Experimental.
func NewCacheUpdated() CacheUpdated {
	_init_.Initialize()

	j := jsiiProxy_CacheUpdated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheUpdated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCacheUpdated_Override(c CacheUpdated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheUpdated",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for Cache Updated.
// Experimental.
func CacheUpdated_EventPattern(options *CacheUpdated_CacheUpdatedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCacheUpdated_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheUpdated",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

