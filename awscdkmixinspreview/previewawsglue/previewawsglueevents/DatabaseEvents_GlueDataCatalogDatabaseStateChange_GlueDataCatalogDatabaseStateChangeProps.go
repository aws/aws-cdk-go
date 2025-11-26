package previewawsglueevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Database aws.glue@GlueDataCatalogDatabaseStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueDataCatalogDatabaseStateChangeProps := &GlueDataCatalogDatabaseStateChangeProps{
//   	ChangedTables: []*string{
//   		jsii.String("changedTables"),
//   	},
//   	DatabaseName: []*string{
//   		jsii.String("databaseName"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	TypeOfChange: []*string{
//   		jsii.String("typeOfChange"),
//   	},
//   }
//
// Experimental.
type DatabaseEvents_GlueDataCatalogDatabaseStateChange_GlueDataCatalogDatabaseStateChangeProps struct {
	// changedTables property.
	//
	// Specify an array of string values to match this event if the actual value of changedTables is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChangedTables *[]*string `field:"optional" json:"changedTables" yaml:"changedTables"`
	// databaseName property.
	//
	// Specify an array of string values to match this event if the actual value of databaseName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Database reference.
	//
	// Experimental.
	DatabaseName *[]*string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// typeOfChange property.
	//
	// Specify an array of string values to match this event if the actual value of typeOfChange is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TypeOfChange *[]*string `field:"optional" json:"typeOfChange" yaml:"typeOfChange"`
}

