package mixinsawsopensearchserverless


// Additional parameters for the k-NN algorithm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parametersProperty := &ParametersProperty{
//   	EfConstruction: jsii.Number(123),
//   	M: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-parameters.html
//
type CfnIndexPropsMixin_ParametersProperty struct {
	// The size of the dynamic list used during k-NN graph creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-parameters.html#cfn-opensearchserverless-index-parameters-efconstruction
	//
	EfConstruction *float64 `field:"optional" json:"efConstruction" yaml:"efConstruction"`
	// Number of neighbors to consider during k-NN search.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-parameters.html#cfn-opensearchserverless-index-parameters-m
	//
	M *float64 `field:"optional" json:"m" yaml:"m"`
}

