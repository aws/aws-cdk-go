package previewawssagemakermixins


// A list of algorithms that were used to create a model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceAlgorithmSpecificationProperty := &SourceAlgorithmSpecificationProperty{
//   	SourceAlgorithms: []interface{}{
//   		&SourceAlgorithmProperty{
//   			AlgorithmName: jsii.String("algorithmName"),
//   			ModelDataUrl: jsii.String("modelDataUrl"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-sourcealgorithmspecification.html
//
type CfnModelPackagePropsMixin_SourceAlgorithmSpecificationProperty struct {
	// A list of the algorithms that were used to create a model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-sourcealgorithmspecification.html#cfn-sagemaker-modelpackage-sourcealgorithmspecification-sourcealgorithms
	//
	SourceAlgorithms interface{} `field:"optional" json:"sourceAlgorithms" yaml:"sourceAlgorithms"`
}

