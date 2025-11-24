package mixinsawseventschemas


// Properties for CfnSchemaPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSchemaMixinProps := &CfnSchemaMixinProps{
//   	Content: jsii.String("content"),
//   	Description: jsii.String("description"),
//   	RegistryName: jsii.String("registryName"),
//   	SchemaName: jsii.String("schemaName"),
//   	Tags: []TagsEntryProperty{
//   		&TagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html
//
type CfnSchemaMixinProps struct {
	// The source of the schema definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
	// A description of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the schema registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-registryname
	//
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// The name of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-schemaname
	//
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Tags associated with the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-tags
	//
	Tags *[]*CfnSchemaPropsMixin_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
	// The type of schema.
	//
	// Valid types include `OpenApi3` and `JSONSchemaDraft4` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

