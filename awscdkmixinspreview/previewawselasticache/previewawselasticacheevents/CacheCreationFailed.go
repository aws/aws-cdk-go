package previewawselasticacheevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.elasticache@CacheCreationFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cacheCreationFailed := awscdkmixinspreview.Events.NewCacheCreationFailed()
//
// Experimental.
type CacheCreationFailed interface {
}

// The jsii proxy struct for CacheCreationFailed
type jsiiProxy_CacheCreationFailed struct {
	_ byte // padding
}

// Experimental.
func NewCacheCreationFailed() CacheCreationFailed {
	_init_.Initialize()

	j := jsiiProxy_CacheCreationFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheCreationFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCacheCreationFailed_Override(c CacheCreationFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheCreationFailed",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for Cache Creation Failed.
// Experimental.
func CacheCreationFailed_CacheCreationFailedPattern(options *CacheCreationFailed_CacheCreationFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCacheCreationFailed_CacheCreationFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheCreationFailed",
		"cacheCreationFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

