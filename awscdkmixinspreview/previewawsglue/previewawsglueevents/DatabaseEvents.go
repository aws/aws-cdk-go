package previewawsglueevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
)

// EventBridge event patterns for Database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var databaseRef IDatabaseRef
//
//   databaseEvents := awscdkmixinspreview.Events.DatabaseEvents_FromDatabase(databaseRef)
//
// Experimental.
type DatabaseEvents interface {
	// EventBridge event pattern for Database Glue Data Catalog Database State Change.
	// Experimental.
	GlueDataCatalogDatabaseStateChangePattern(options *DatabaseEvents_GlueDataCatalogDatabaseStateChange_GlueDataCatalogDatabaseStateChangeProps) *awsevents.EventPattern
	// EventBridge event pattern for Database Glue Data Catalog Table State Change.
	// Experimental.
	GlueDataCatalogTableStateChangePattern(options *DatabaseEvents_GlueDataCatalogTableStateChange_GlueDataCatalogTableStateChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for DatabaseEvents
type jsiiProxy_DatabaseEvents struct {
	_ byte // padding
}

// Create DatabaseEvents from a Database reference.
// Experimental.
func DatabaseEvents_FromDatabase(databaseRef interfacesawsglue.IDatabaseRef) DatabaseEvents {
	_init_.Initialize()

	if err := validateDatabaseEvents_FromDatabaseParameters(databaseRef); err != nil {
		panic(err)
	}
	var returns DatabaseEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.events.DatabaseEvents",
		"fromDatabase",
		[]interface{}{databaseRef},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseEvents) GlueDataCatalogDatabaseStateChangePattern(options *DatabaseEvents_GlueDataCatalogDatabaseStateChange_GlueDataCatalogDatabaseStateChangeProps) *awsevents.EventPattern {
	if err := d.validateGlueDataCatalogDatabaseStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"glueDataCatalogDatabaseStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseEvents) GlueDataCatalogTableStateChangePattern(options *DatabaseEvents_GlueDataCatalogTableStateChange_GlueDataCatalogTableStateChangeProps) *awsevents.EventPattern {
	if err := d.validateGlueDataCatalogTableStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"glueDataCatalogTableStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

