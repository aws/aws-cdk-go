package awseventschemas


// Properties for defining a `CfnSchema`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSchemaProps := &CfnSchemaProps{
//   	Content: jsii.String("content"),
//   	RegistryName: jsii.String("registryName"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	SchemaName: jsii.String("schemaName"),
//   	Tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html
//
type CfnSchemaProps struct {
	// The source of the schema definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-content
	//
	Content *string `field:"required" json:"content" yaml:"content"`
	// The name of the schema registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-registryname
	//
	RegistryName *string `field:"required" json:"registryName" yaml:"registryName"`
	// The type of schema.
	//
	// Valid types include `OpenApi3` and `JSONSchemaDraft4` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-schemaname
	//
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Tags associated with the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-tags
	//
	Tags *[]*CfnSchema_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}

