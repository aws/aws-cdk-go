package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ediConfigurationProperty := &EdiConfigurationProperty{
//   	InputLocation: &S3LocationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		Key: jsii.String("key"),
//   	},
//   	OutputLocation: &S3LocationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		Key: jsii.String("key"),
//   	},
//   	TransformerId: jsii.String("transformerId"),
//   	Type: &EdiTypeProperty{
//   		X12Details: &X12DetailsProperty{
//   			TransactionSet: jsii.String("transactionSet"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html
//
type CfnCapability_EdiConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-inputlocation
	//
	InputLocation interface{} `field:"required" json:"inputLocation" yaml:"inputLocation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-outputlocation
	//
	OutputLocation interface{} `field:"required" json:"outputLocation" yaml:"outputLocation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-transformerid
	//
	TransformerId *string `field:"required" json:"transformerId" yaml:"transformerId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-type
	//
	Type interface{} `field:"required" json:"type" yaml:"type"`
}

