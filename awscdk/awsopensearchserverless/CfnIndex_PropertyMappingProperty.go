package awsopensearchserverless


// Property mappings for the OpenSearch Serverless index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var propertyMappingProperty_ PropertyMappingProperty
//
//   propertyMappingProperty := &PropertyMappingProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Analyzer: jsii.String("analyzer"),
//   	CompressionLevel: jsii.String("compressionLevel"),
//   	Dimension: jsii.Number(123),
//   	Index: jsii.Boolean(false),
//   	Method: &MethodProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Engine: jsii.String("engine"),
//   		Parameters: &ParametersProperty{
//   			EfConstruction: jsii.Number(123),
//   			M: jsii.Number(123),
//   		},
//   		SpaceType: jsii.String("spaceType"),
//   	},
//   	Properties: map[string]interface{}{
//   		"propertiesKey": propertyMappingProperty_,
//   	},
//   	SpaceType: jsii.String("spaceType"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html
//
type CfnIndex_PropertyMappingProperty struct {
	// The field data type.
	//
	// Must be a valid OpenSearch field type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html#cfn-opensearchserverless-index-propertymapping-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The analyzer to use for this field (for text and keyword fields).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html#cfn-opensearchserverless-index-propertymapping-analyzer
	//
	Analyzer *string `field:"optional" json:"analyzer" yaml:"analyzer"`
	// The compression level for knn_vector fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html#cfn-opensearchserverless-index-propertymapping-compressionlevel
	//
	CompressionLevel *string `field:"optional" json:"compressionLevel" yaml:"compressionLevel"`
	// Dimension size for vector fields, defines the number of dimensions in the vector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html#cfn-opensearchserverless-index-propertymapping-dimension
	//
	Dimension *float64 `field:"optional" json:"dimension" yaml:"dimension"`
	// Whether a field should be indexed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html#cfn-opensearchserverless-index-propertymapping-index
	//
	Index interface{} `field:"optional" json:"index" yaml:"index"`
	// Configuration for k-NN search method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html#cfn-opensearchserverless-index-propertymapping-method
	//
	Method interface{} `field:"optional" json:"method" yaml:"method"`
	// Defines the fields within the mapping, including their types and configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html#cfn-opensearchserverless-index-propertymapping-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// The distance function used for k-NN search (field-level, outside Method).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html#cfn-opensearchserverless-index-propertymapping-spacetype
	//
	SpaceType *string `field:"optional" json:"spaceType" yaml:"spaceType"`
	// Default value for the field when not specified in a document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html#cfn-opensearchserverless-index-propertymapping-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

