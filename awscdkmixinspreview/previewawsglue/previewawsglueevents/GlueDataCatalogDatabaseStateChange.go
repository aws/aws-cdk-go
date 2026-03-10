package previewawsglueevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.glue@GlueDataCatalogDatabaseStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueDataCatalogDatabaseStateChange := awscdkmixinspreview.Events.NewGlueDataCatalogDatabaseStateChange()
//
// Experimental.
type GlueDataCatalogDatabaseStateChange interface {
}

// The jsii proxy struct for GlueDataCatalogDatabaseStateChange
type jsiiProxy_GlueDataCatalogDatabaseStateChange struct {
	_ byte // padding
}

// Experimental.
func NewGlueDataCatalogDatabaseStateChange() GlueDataCatalogDatabaseStateChange {
	_init_.Initialize()

	j := jsiiProxy_GlueDataCatalogDatabaseStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueDataCatalogDatabaseStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewGlueDataCatalogDatabaseStateChange_Override(g GlueDataCatalogDatabaseStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueDataCatalogDatabaseStateChange",
		nil, // no parameters
		g,
	)
}

// EventBridge event pattern for Glue Data Catalog Database State Change.
// Experimental.
func GlueDataCatalogDatabaseStateChange_GlueDataCatalogDatabaseStateChangePattern(options *GlueDataCatalogDatabaseStateChange_GlueDataCatalogDatabaseStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateGlueDataCatalogDatabaseStateChange_GlueDataCatalogDatabaseStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueDataCatalogDatabaseStateChange",
		"glueDataCatalogDatabaseStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

