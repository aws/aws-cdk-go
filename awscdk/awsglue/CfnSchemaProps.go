package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSchema`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSchemaProps := &CfnSchemaProps{
//   	Compatibility: jsii.String("compatibility"),
//   	DataFormat: jsii.String("dataFormat"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CheckpointVersion: &SchemaVersionProperty{
//   		IsLatest: jsii.Boolean(false),
//   		VersionNumber: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	Registry: &RegistryProperty{
//   		Arn: jsii.String("arn"),
//   		Name: jsii.String("name"),
//   	},
//   	SchemaDefinition: jsii.String("schemaDefinition"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html
//
type CfnSchemaProps struct {
	// The compatibility mode of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-compatibility
	//
	Compatibility *string `field:"required" json:"compatibility" yaml:"compatibility"`
	// The data format of the schema definition.
	//
	// Currently only `AVRO` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-dataformat
	//
	DataFormat *string `field:"required" json:"dataFormat" yaml:"dataFormat"`
	// Name of the schema to be created of max length of 255, and may only contain letters, numbers, hyphen, underscore, dollar sign, or hash mark.
	//
	// No whitespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify the `VersionNumber` or the `IsLatest` for setting the checkpoint for the schema.
	//
	// This is only required for updating a checkpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-checkpointversion
	//
	CheckpointVersion interface{} `field:"optional" json:"checkpointVersion" yaml:"checkpointVersion"`
	// A description of the schema if specified when created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The registry where a schema is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-registry
	//
	Registry interface{} `field:"optional" json:"registry" yaml:"registry"`
	// The schema definition using the `DataFormat` setting for `SchemaName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-schemadefinition
	//
	SchemaDefinition *string `field:"optional" json:"schemaDefinition" yaml:"schemaDefinition"`
	// AWS tags that contain a key value pair and may be searched by console, command line, or API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

