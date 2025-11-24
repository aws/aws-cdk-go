package mixinsawsopensearchserverless


// Properties for CfnIndexPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var propertyMappingProperty_ PropertyMappingProperty
//
//   cfnIndexMixinProps := &CfnIndexMixinProps{
//   	CollectionEndpoint: jsii.String("collectionEndpoint"),
//   	IndexName: jsii.String("indexName"),
//   	Mappings: &MappingsProperty{
//   		Properties: map[string]interface{}{
//   			"propertiesKey": &PropertyMappingProperty{
//   				"dimension": jsii.Number(123),
//   				"index": jsii.Boolean(false),
//   				"method": &MethodProperty{
//   					"engine": jsii.String("engine"),
//   					"name": jsii.String("name"),
//   					"parameters": &ParametersProperty{
//   						"efConstruction": jsii.Number(123),
//   						"m": jsii.Number(123),
//   					},
//   					"spaceType": jsii.String("spaceType"),
//   				},
//   				"properties": map[string]interface{}{
//   					"propertiesKey": propertyMappingProperty_,
//   				},
//   				"type": jsii.String("type"),
//   				"value": jsii.String("value"),
//   			},
//   		},
//   	},
//   	Settings: &IndexSettingsProperty{
//   		Index: &IndexProperty{
//   			Knn: jsii.Boolean(false),
//   			KnnAlgoParamEfSearch: jsii.Number(123),
//   			RefreshInterval: jsii.String("refreshInterval"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-index.html
//
type CfnIndexMixinProps struct {
	// The endpoint for the collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-index.html#cfn-opensearchserverless-index-collectionendpoint
	//
	CollectionEndpoint *string `field:"optional" json:"collectionEndpoint" yaml:"collectionEndpoint"`
	// The name of the OpenSearch Serverless index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-index.html#cfn-opensearchserverless-index-indexname
	//
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Index mappings for the OpenSearch Serverless index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-index.html#cfn-opensearchserverless-index-mappings
	//
	Mappings interface{} `field:"optional" json:"mappings" yaml:"mappings"`
	// Index settings for the OpenSearch Serverless index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-index.html#cfn-opensearchserverless-index-settings
	//
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
}

