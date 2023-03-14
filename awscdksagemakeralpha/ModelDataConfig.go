// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha


// The configuration needed to reference model artifacts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import sagemaker_alpha "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   modelDataConfig := &ModelDataConfig{
//   	Uri: jsii.String("uri"),
//   }
//
// Experimental.
type ModelDataConfig struct {
	// The S3 path where the model artifacts, which result from model training, are stored.
	//
	// This path
	// must point to a single gzip compressed tar archive (.tar.gz suffix).
	// Experimental.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

