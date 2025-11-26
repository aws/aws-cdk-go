package previewawsopensearchserverlessmixins


// Property mappings for the OpenSearch Serverless index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var propertyMappingProperty_ PropertyMappingProperty
//
//   propertyMappingProperty := &PropertyMappingProperty{
//   	Dimension: jsii.Number(123),
//   	Index: jsii.Boolean(false),
//   	Method: &MethodProperty{
//   		Engine: jsii.String("engine"),
//   		Name: jsii.String("name"),
//   		Parameters: &ParametersProperty{
//   			EfConstruction: jsii.Number(123),
//   			M: jsii.Number(123),
//   		},
//   		SpaceType: jsii.String("spaceType"),
//   	},
//   	Properties: map[string]interface{}{
//   		"propertiesKey": propertyMappingProperty_,
//   	},
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html
//
type CfnIndexPropsMixin_PropertyMappingProperty struct {
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
	// The field data type.
	//
	// Must be a valid OpenSearch field type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html#cfn-opensearchserverless-index-propertymapping-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Default value for the field when not specified in a document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-propertymapping.html#cfn-opensearchserverless-index-propertymapping-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

