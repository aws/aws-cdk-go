package awsglue


// The algorithm-specific parameters that are associated with the machine learning transform.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformParametersProperty := &TransformParametersProperty{
//   	TransformType: jsii.String("transformType"),
//
//   	// the properties below are optional
//   	FindMatchesParameters: &FindMatchesParametersProperty{
//   		PrimaryKeyColumnName: jsii.String("primaryKeyColumnName"),
//
//   		// the properties below are optional
//   		AccuracyCostTradeoff: jsii.Number(123),
//   		EnforceProvidedLabels: jsii.Boolean(false),
//   		PrecisionRecallTradeoff: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-transformparameters.html
//
type CfnMLTransform_TransformParametersProperty struct {
	// The type of machine learning transform. `FIND_MATCHES` is the only option.
	//
	// For information about the types of machine learning transforms, see [Creating Machine Learning Transforms](https://docs.aws.amazon.com/glue/latest/dg/add-job-machine-learning-transform.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-transformparameters.html#cfn-glue-mltransform-transformparameters-transformtype
	//
	TransformType *string `field:"required" json:"transformType" yaml:"transformType"`
	// The parameters for the find matches algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-transformparameters.html#cfn-glue-mltransform-transformparameters-findmatchesparameters
	//
	FindMatchesParameters interface{} `field:"optional" json:"findMatchesParameters" yaml:"findMatchesParameters"`
}

