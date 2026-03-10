package previewawsglueevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.glue@GlueJobRunStatus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueJobRunStatus := awscdkmixinspreview.Events.NewGlueJobRunStatus()
//
// Experimental.
type GlueJobRunStatus interface {
}

// The jsii proxy struct for GlueJobRunStatus
type jsiiProxy_GlueJobRunStatus struct {
	_ byte // padding
}

// Experimental.
func NewGlueJobRunStatus() GlueJobRunStatus {
	_init_.Initialize()

	j := jsiiProxy_GlueJobRunStatus{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueJobRunStatus",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewGlueJobRunStatus_Override(g GlueJobRunStatus) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueJobRunStatus",
		nil, // no parameters
		g,
	)
}

// EventBridge event pattern for Glue Job Run Status.
// Experimental.
func GlueJobRunStatus_GlueJobRunStatusPattern(options *GlueJobRunStatus_GlueJobRunStatusProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateGlueJobRunStatus_GlueJobRunStatusPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueJobRunStatus",
		"glueJobRunStatusPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

