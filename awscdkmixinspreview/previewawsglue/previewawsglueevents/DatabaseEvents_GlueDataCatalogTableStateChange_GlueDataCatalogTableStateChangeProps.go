package previewawsglueevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Database aws.glue@GlueDataCatalogTableStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueDataCatalogTableStateChangeProps := &GlueDataCatalogTableStateChangeProps{
//   	ChangedPartitions: []*string{
//   		jsii.String("changedPartitions"),
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
//   	TableName: []*string{
//   		jsii.String("tableName"),
//   	},
//   	TypeOfChange: []*string{
//   		jsii.String("typeOfChange"),
//   	},
//   }
//
// Experimental.
type DatabaseEvents_GlueDataCatalogTableStateChange_GlueDataCatalogTableStateChangeProps struct {
	// changedPartitions property.
	//
	// Specify an array of string values to match this event if the actual value of changedPartitions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChangedPartitions *[]*string `field:"optional" json:"changedPartitions" yaml:"changedPartitions"`
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
	// tableName property.
	//
	// Specify an array of string values to match this event if the actual value of tableName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TableName *[]*string `field:"optional" json:"tableName" yaml:"tableName"`
	// typeOfChange property.
	//
	// Specify an array of string values to match this event if the actual value of typeOfChange is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TypeOfChange *[]*string `field:"optional" json:"typeOfChange" yaml:"typeOfChange"`
}

