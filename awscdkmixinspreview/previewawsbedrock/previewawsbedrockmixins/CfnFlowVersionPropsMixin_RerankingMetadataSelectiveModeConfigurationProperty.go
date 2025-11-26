package previewawsbedrockmixins


// Configuration for selectively including or excluding metadata fields during the reranking process.
//
// This allows you to control which metadata attributes are considered when reordering search results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-rerankingmetadataselectivemodeconfiguration.html
//
type CfnFlowVersionPropsMixin_RerankingMetadataSelectiveModeConfigurationProperty struct {
	// A list of metadata field names to explicitly exclude from the reranking process.
	//
	// All metadata fields except these will be considered when reordering search results. This parameter cannot be used together with fieldsToInclude.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-rerankingmetadataselectivemodeconfiguration.html#cfn-bedrock-flowversion-rerankingmetadataselectivemodeconfiguration-fieldstoexclude
	//
	FieldsToExclude interface{} `field:"optional" json:"fieldsToExclude" yaml:"fieldsToExclude"`
	// A list of metadata field names to explicitly include in the reranking process.
	//
	// Only these fields will be considered when reordering search results. This parameter cannot be used together with fieldsToExclude.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-rerankingmetadataselectivemodeconfiguration.html#cfn-bedrock-flowversion-rerankingmetadataselectivemodeconfiguration-fieldstoinclude
	//
	FieldsToInclude interface{} `field:"optional" json:"fieldsToInclude" yaml:"fieldsToInclude"`
}

