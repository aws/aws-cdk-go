package awsopensearchserverless


// Properties for defining a `CfnIndex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var propertyMappingProperty_ propertyMappingProperty
//
//   cfnIndexProps := &CfnIndexProps{
//   	CollectionEndpoint: jsii.String("collectionEndpoint"),
//   	IndexName: jsii.String("indexName"),
//
//   	// the properties below are optional
//   	Mappings: &MappingsProperty{
//   		Properties: map[string]interface{}{
//   			"propertiesKey": &propertyMappingProperty{
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"dimension": jsii.Number(123),
//   				"index": jsii.Boolean(false),
//   				"method": &MethodProperty{
//   					"engine": jsii.String("engine"),
//   					"name": jsii.String("name"),
//
//   					// the properties below are optional
//   					"parameters": &ParametersProperty{
//   						"efConstruction": jsii.Number(123),
//   						"m": jsii.Number(123),
//   					},
//   					"spaceType": jsii.String("spaceType"),
//   				},
//   				"properties": map[string]interface{}{
//   					"propertiesKey": propertyMappingProperty_,
//   				},
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
type CfnIndexProps struct {
	// The endpoint for the collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-index.html#cfn-opensearchserverless-index-collectionendpoint
	//
	CollectionEndpoint *string `field:"required" json:"collectionEndpoint" yaml:"collectionEndpoint"`
	// The name of the OpenSearch Serverless index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-index.html#cfn-opensearchserverless-index-indexname
	//
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Index Mappings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-index.html#cfn-opensearchserverless-index-mappings
	//
	Mappings interface{} `field:"optional" json:"mappings" yaml:"mappings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-index.html#cfn-opensearchserverless-index-settings
	//
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
}

