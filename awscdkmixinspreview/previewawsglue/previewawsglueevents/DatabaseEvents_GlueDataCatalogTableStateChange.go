package previewawsglueevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.glue@GlueDataCatalogTableStateChange event types for Database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueDataCatalogTableStateChange := #error#.NewGlueDataCatalogTableStateChange()
//
// Experimental.
type DatabaseEvents_GlueDataCatalogTableStateChange interface {
}

// The jsii proxy struct for DatabaseEvents_GlueDataCatalogTableStateChange
type jsiiProxy_DatabaseEvents_GlueDataCatalogTableStateChange struct {
	_ byte // padding
}

// Experimental.
func NewDatabaseEvents_GlueDataCatalogTableStateChange() DatabaseEvents_GlueDataCatalogTableStateChange {
	_init_.Initialize()

	j := jsiiProxy_DatabaseEvents_GlueDataCatalogTableStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.DatabaseEvents.GlueDataCatalogTableStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatabaseEvents_GlueDataCatalogTableStateChange_Override(d DatabaseEvents_GlueDataCatalogTableStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.DatabaseEvents.GlueDataCatalogTableStateChange",
		nil, // no parameters
		d,
	)
}

