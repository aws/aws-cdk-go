package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceAlgorithmProperty := &SourceAlgorithmProperty{
//   	AlgorithmName: jsii.String("algorithmName"),
//
//   	// the properties below are optional
//   	ModelDataUrl: jsii.String("modelDataUrl"),
//   }
//
type CfnModelCard_SourceAlgorithmProperty struct {
	// `CfnModelCard.SourceAlgorithmProperty.AlgorithmName`.
	AlgorithmName *string `field:"required" json:"algorithmName" yaml:"algorithmName"`
	// `CfnModelCard.SourceAlgorithmProperty.ModelDataUrl`.
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
}

