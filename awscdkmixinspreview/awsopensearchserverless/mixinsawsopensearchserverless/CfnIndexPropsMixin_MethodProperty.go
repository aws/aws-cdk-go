package mixinsawsopensearchserverless


// Configuration for k-NN search method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   methodProperty := &MethodProperty{
//   	Engine: jsii.String("engine"),
//   	Name: jsii.String("name"),
//   	Parameters: &ParametersProperty{
//   		EfConstruction: jsii.Number(123),
//   		M: jsii.Number(123),
//   	},
//   	SpaceType: jsii.String("spaceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-method.html
//
type CfnIndexPropsMixin_MethodProperty struct {
	// The k-NN search engine to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-method.html#cfn-opensearchserverless-index-method-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The algorithm name for k-NN search.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-method.html#cfn-opensearchserverless-index-method-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Additional parameters for the k-NN algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-method.html#cfn-opensearchserverless-index-method-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The distance function used for k-NN search.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-method.html#cfn-opensearchserverless-index-method-spacetype
	//
	SpaceType *string `field:"optional" json:"spaceType" yaml:"spaceType"`
}

