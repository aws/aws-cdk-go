package mixinsawsbedrock


// Configuration for how metadata should be used during the reranking process in Knowledge Base vector searches.
//
// This determines which metadata fields are included or excluded when reordering search results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metadataConfigurationForRerankingProperty := &MetadataConfigurationForRerankingProperty{
//   	SelectionMode: jsii.String("selectionMode"),
//   	SelectiveModeConfiguration: &RerankingMetadataSelectiveModeConfigurationProperty{
//   		FieldsToExclude: []interface{}{
//   			&FieldForRerankingProperty{
//   				FieldName: jsii.String("fieldName"),
//   			},
//   		},
//   		FieldsToInclude: []interface{}{
//   			&FieldForRerankingProperty{
//   				FieldName: jsii.String("fieldName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-metadataconfigurationforreranking.html
//
type CfnFlowVersionPropsMixin_MetadataConfigurationForRerankingProperty struct {
	// The mode for selecting which metadata fields to include in the reranking process.
	//
	// Valid values are ALL (use all available metadata fields) or SELECTIVE (use only specified fields).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-metadataconfigurationforreranking.html#cfn-bedrock-flowversion-metadataconfigurationforreranking-selectionmode
	//
	SelectionMode *string `field:"optional" json:"selectionMode" yaml:"selectionMode"`
	// Configuration for selective mode, which allows you to explicitly include or exclude specific metadata fields during reranking.
	//
	// This is only used when selectionMode is set to SELECTIVE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-metadataconfigurationforreranking.html#cfn-bedrock-flowversion-metadataconfigurationforreranking-selectivemodeconfiguration
	//
	SelectiveModeConfiguration interface{} `field:"optional" json:"selectiveModeConfiguration" yaml:"selectiveModeConfiguration"`
}

