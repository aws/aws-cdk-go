package awsglue


// The algorithm-specific parameters that are associated with the machine learning transform.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformParametersProperty := &transformParametersProperty{
//   	transformType: jsii.String("transformType"),
//
//   	// the properties below are optional
//   	findMatchesParameters: &findMatchesParametersProperty{
//   		primaryKeyColumnName: jsii.String("primaryKeyColumnName"),
//
//   		// the properties below are optional
//   		accuracyCostTradeoff: jsii.Number(123),
//   		enforceProvidedLabels: jsii.Boolean(false),
//   		precisionRecallTradeoff: jsii.Number(123),
//   	},
//   }
//
type CfnMLTransform_TransformParametersProperty struct {
	// The type of machine learning transform. `FIND_MATCHES` is the only option.
	//
	// For information about the types of machine learning transforms, see [Creating Machine Learning Transforms](https://docs.aws.amazon.com/glue/latest/dg/add-job-machine-learning-transform.html) .
	TransformType *string `field:"required" json:"transformType" yaml:"transformType"`
	// The parameters for the find matches algorithm.
	FindMatchesParameters interface{} `field:"optional" json:"findMatchesParameters" yaml:"findMatchesParameters"`
}

