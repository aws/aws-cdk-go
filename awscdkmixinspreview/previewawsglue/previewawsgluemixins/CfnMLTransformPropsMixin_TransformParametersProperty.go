package previewawsgluemixins


// The algorithm-specific parameters that are associated with the machine learning transform.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transformParametersProperty := &TransformParametersProperty{
//   	FindMatchesParameters: &FindMatchesParametersProperty{
//   		AccuracyCostTradeoff: jsii.Number(123),
//   		EnforceProvidedLabels: jsii.Boolean(false),
//   		PrecisionRecallTradeoff: jsii.Number(123),
//   		PrimaryKeyColumnName: jsii.String("primaryKeyColumnName"),
//   	},
//   	TransformType: jsii.String("transformType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-transformparameters.html
//
type CfnMLTransformPropsMixin_TransformParametersProperty struct {
	// The parameters for the find matches algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-transformparameters.html#cfn-glue-mltransform-transformparameters-findmatchesparameters
	//
	FindMatchesParameters interface{} `field:"optional" json:"findMatchesParameters" yaml:"findMatchesParameters"`
	// The type of machine learning transform. `FIND_MATCHES` is the only option.
	//
	// For information about the types of machine learning transforms, see [Working with machine learning transforms](https://docs.aws.amazon.com/glue/latest/dg/console-machine-learning-transforms.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-transformparameters.html#cfn-glue-mltransform-transformparameters-transformtype
	//
	TransformType *string `field:"optional" json:"transformType" yaml:"transformType"`
}

