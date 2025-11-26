package previewawsglueevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.glue@GlueDataCatalogDatabaseStateChange event types for Database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueDataCatalogDatabaseStateChange := #error#.NewGlueDataCatalogDatabaseStateChange()
//
// Experimental.
type DatabaseEvents_GlueDataCatalogDatabaseStateChange interface {
}

// The jsii proxy struct for DatabaseEvents_GlueDataCatalogDatabaseStateChange
type jsiiProxy_DatabaseEvents_GlueDataCatalogDatabaseStateChange struct {
	_ byte // padding
}

// Experimental.
func NewDatabaseEvents_GlueDataCatalogDatabaseStateChange() DatabaseEvents_GlueDataCatalogDatabaseStateChange {
	_init_.Initialize()

	j := jsiiProxy_DatabaseEvents_GlueDataCatalogDatabaseStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.DatabaseEvents.GlueDataCatalogDatabaseStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatabaseEvents_GlueDataCatalogDatabaseStateChange_Override(d DatabaseEvents_GlueDataCatalogDatabaseStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.DatabaseEvents.GlueDataCatalogDatabaseStateChange",
		nil, // no parameters
		d,
	)
}

