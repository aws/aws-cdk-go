package previewawsglueevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.glue@GlueDataCatalogTableStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueDataCatalogTableStateChange := awscdkmixinspreview.Events.NewGlueDataCatalogTableStateChange()
//
// Experimental.
type GlueDataCatalogTableStateChange interface {
}

// The jsii proxy struct for GlueDataCatalogTableStateChange
type jsiiProxy_GlueDataCatalogTableStateChange struct {
	_ byte // padding
}

// Experimental.
func NewGlueDataCatalogTableStateChange() GlueDataCatalogTableStateChange {
	_init_.Initialize()

	j := jsiiProxy_GlueDataCatalogTableStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueDataCatalogTableStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewGlueDataCatalogTableStateChange_Override(g GlueDataCatalogTableStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueDataCatalogTableStateChange",
		nil, // no parameters
		g,
	)
}

// EventBridge event pattern for Glue Data Catalog Table State Change.
// Experimental.
func GlueDataCatalogTableStateChange_GlueDataCatalogTableStateChangePattern(options *GlueDataCatalogTableStateChange_GlueDataCatalogTableStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateGlueDataCatalogTableStateChange_GlueDataCatalogTableStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueDataCatalogTableStateChange",
		"glueDataCatalogTableStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

