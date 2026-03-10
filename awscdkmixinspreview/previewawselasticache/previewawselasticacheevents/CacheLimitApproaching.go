package previewawselasticacheevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.elasticache@CacheLimitApproaching.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cacheLimitApproaching := awscdkmixinspreview.Events.NewCacheLimitApproaching()
//
// Experimental.
type CacheLimitApproaching interface {
}

// The jsii proxy struct for CacheLimitApproaching
type jsiiProxy_CacheLimitApproaching struct {
	_ byte // padding
}

// Experimental.
func NewCacheLimitApproaching() CacheLimitApproaching {
	_init_.Initialize()

	j := jsiiProxy_CacheLimitApproaching{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheLimitApproaching",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCacheLimitApproaching_Override(c CacheLimitApproaching) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheLimitApproaching",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for Cache Limit Approaching.
// Experimental.
func CacheLimitApproaching_CacheLimitApproachingPattern(options *CacheLimitApproaching_CacheLimitApproachingProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCacheLimitApproaching_CacheLimitApproachingPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.events.CacheLimitApproaching",
		"cacheLimitApproachingPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

