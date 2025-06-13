package awsbedrock


// Contains configurations for the metadata fields to include or exclude when considering reranking.
//
// If you include the `fieldsToExclude` field, the reranker ignores all the metadata fields that you specify. If you include the `fieldsToInclude` field, the reranker uses only the metadata fields that you specify and ignores all others. You can include only one of these fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rerankingMetadataSelectiveModeConfigurationProperty := &RerankingMetadataSelectiveModeConfigurationProperty{
//   	FieldsToExclude: []interface{}{
//   		&FieldForRerankingProperty{
//   			FieldName: jsii.String("fieldName"),
//   		},
//   	},
//   	FieldsToInclude: []interface{}{
//   		&FieldForRerankingProperty{
//   			FieldName: jsii.String("fieldName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-rerankingmetadataselectivemodeconfiguration.html
//
type CfnFlow_RerankingMetadataSelectiveModeConfigurationProperty struct {
	// An array of objects, each of which specifies a metadata field to exclude from consideration when reranking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-rerankingmetadataselectivemodeconfiguration.html#cfn-bedrock-flow-rerankingmetadataselectivemodeconfiguration-fieldstoexclude
	//
	FieldsToExclude interface{} `field:"optional" json:"fieldsToExclude" yaml:"fieldsToExclude"`
	// An array of objects, each of which specifies a metadata field to include in consideration when reranking.
	//
	// The remaining metadata fields are ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-rerankingmetadataselectivemodeconfiguration.html#cfn-bedrock-flow-rerankingmetadataselectivemodeconfiguration-fieldstoinclude
	//
	FieldsToInclude interface{} `field:"optional" json:"fieldsToInclude" yaml:"fieldsToInclude"`
}

