package previewawsglueevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.glue@GlueJobStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueJobStateChange := awscdkmixinspreview.Events.NewGlueJobStateChange()
//
// Experimental.
type GlueJobStateChange interface {
}

// The jsii proxy struct for GlueJobStateChange
type jsiiProxy_GlueJobStateChange struct {
	_ byte // padding
}

// Experimental.
func NewGlueJobStateChange() GlueJobStateChange {
	_init_.Initialize()

	j := jsiiProxy_GlueJobStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueJobStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewGlueJobStateChange_Override(g GlueJobStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueJobStateChange",
		nil, // no parameters
		g,
	)
}

// EventBridge event pattern for Glue Job State Change.
// Experimental.
func GlueJobStateChange_GlueJobStateChangePattern(options *GlueJobStateChange_GlueJobStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateGlueJobStateChange_GlueJobStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueJobStateChange",
		"glueJobStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

