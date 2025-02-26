package awssagemaker


// A list of algorithms that were used to create a model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceAlgorithmSpecificationProperty := &SourceAlgorithmSpecificationProperty{
//   	SourceAlgorithms: []interface{}{
//   		&SourceAlgorithmProperty{
//   			AlgorithmName: jsii.String("algorithmName"),
//
//   			// the properties below are optional
//   			ModelDataUrl: jsii.String("modelDataUrl"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-sourcealgorithmspecification.html
//
type CfnModelPackage_SourceAlgorithmSpecificationProperty struct {
	// A list of the algorithms that were used to create a model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-sourcealgorithmspecification.html#cfn-sagemaker-modelpackage-sourcealgorithmspecification-sourcealgorithms
	//
	SourceAlgorithms interface{} `field:"required" json:"sourceAlgorithms" yaml:"sourceAlgorithms"`
}

