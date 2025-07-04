package awsbedrock


// Contains configurations for the metadata to use in reranking.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataConfigurationForRerankingProperty := &MetadataConfigurationForRerankingProperty{
//   	SelectionMode: jsii.String("selectionMode"),
//
//   	// the properties below are optional
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
type CfnFlowVersion_MetadataConfigurationForRerankingProperty struct {
	// Specifies whether to consider all metadata when reranking, or only the metadata that you select.
	//
	// If you specify `SELECTIVE` , include the `selectiveModeConfiguration` field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-metadataconfigurationforreranking.html#cfn-bedrock-flowversion-metadataconfigurationforreranking-selectionmode
	//
	SelectionMode *string `field:"required" json:"selectionMode" yaml:"selectionMode"`
	// Contains configurations for the metadata fields to include or exclude when considering reranking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-metadataconfigurationforreranking.html#cfn-bedrock-flowversion-metadataconfigurationforreranking-selectivemodeconfiguration
	//
	SelectiveModeConfiguration interface{} `field:"optional" json:"selectiveModeConfiguration" yaml:"selectiveModeConfiguration"`
}

