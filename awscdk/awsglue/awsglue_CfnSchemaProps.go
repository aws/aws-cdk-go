package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnSchema`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSchemaProps := &cfnSchemaProps{
//   	compatibility: jsii.String("compatibility"),
//   	dataFormat: jsii.String("dataFormat"),
//   	name: jsii.String("name"),
//   	schemaDefinition: jsii.String("schemaDefinition"),
//
//   	// the properties below are optional
//   	checkpointVersion: &schemaVersionProperty{
//   		isLatest: jsii.Boolean(false),
//   		versionNumber: jsii.Number(123),
//   	},
//   	description: jsii.String("description"),
//   	registry: &registryProperty{
//   		arn: jsii.String("arn"),
//   		name: jsii.String("name"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSchemaProps struct {
	// The compatibility mode of the schema.
	Compatibility *string `field:"required" json:"compatibility" yaml:"compatibility"`
	// The data format of the schema definition.
	//
	// Currently only `AVRO` is supported.
	DataFormat *string `field:"required" json:"dataFormat" yaml:"dataFormat"`
	// Name of the schema to be created of max length of 255, and may only contain letters, numbers, hyphen, underscore, dollar sign, or hash mark.
	//
	// No whitespace.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema definition using the `DataFormat` setting for `SchemaName` .
	SchemaDefinition *string `field:"required" json:"schemaDefinition" yaml:"schemaDefinition"`
	// Specify the `VersionNumber` or the `IsLatest` for setting the checkpoint for the schema.
	//
	// This is only required for updating a checkpoint.
	CheckpointVersion interface{} `field:"optional" json:"checkpointVersion" yaml:"checkpointVersion"`
	// A description of the schema if specified when created.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The registry where a schema is stored.
	Registry interface{} `field:"optional" json:"registry" yaml:"registry"`
	// AWS tags that contain a key value pair and may be searched by console, command line, or API.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

