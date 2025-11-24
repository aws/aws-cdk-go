package mixinsawsentityresolution

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSchemaMappingPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSchemaMappingMixinProps := &CfnSchemaMappingMixinProps{
//   	Description: jsii.String("description"),
//   	MappedInputFields: []interface{}{
//   		&SchemaInputAttributeProperty{
//   			FieldName: jsii.String("fieldName"),
//   			GroupName: jsii.String("groupName"),
//   			Hashed: jsii.Boolean(false),
//   			MatchKey: jsii.String("matchKey"),
//   			SubType: jsii.String("subType"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SchemaName: jsii.String("schemaName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-schemamapping.html
//
type CfnSchemaMappingMixinProps struct {
	// A description of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-schemamapping.html#cfn-entityresolution-schemamapping-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of `MappedInputFields` .
	//
	// Each `MappedInputField` corresponds to a column the source data table, and contains column name plus additional information that AWS Entity Resolution uses for matching.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-schemamapping.html#cfn-entityresolution-schemamapping-mappedinputfields
	//
	MappedInputFields interface{} `field:"optional" json:"mappedInputFields" yaml:"mappedInputFields"`
	// The name of the schema.
	//
	// There can't be multiple `SchemaMappings` with the same name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-schemamapping.html#cfn-entityresolution-schemamapping-schemaname
	//
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-schemamapping.html#cfn-entityresolution-schemamapping-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

