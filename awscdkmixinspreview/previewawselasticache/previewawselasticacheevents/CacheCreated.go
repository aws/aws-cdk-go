package previewawselasticacheevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.elasticache@CacheCreated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cacheCreated := awscdkmixinspreview.Events.NewCacheCreated()
//
// Experimental.
type CacheCreated interface {
}

// The jsii proxy struct for CacheCreated
type jsiiProxy_CacheCreated struct {
	_ byte // padding
}

// Experimental.
func NewCacheCreated() CacheCreated {
	_init_.Initialize()

	j := jsiiProxy_CacheCreated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheCreated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCacheCreated_Override(c CacheCreated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheCreated",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for Cache Created.
// Experimental.
func CacheCreated_EventPattern(options *CacheCreated_CacheCreatedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCacheCreated_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheCreated",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

