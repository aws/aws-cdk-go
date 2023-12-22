package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capabilityConfigurationProperty := &CapabilityConfigurationProperty{
//   	Edi: &EdiConfigurationProperty{
//   		InputLocation: &S3LocationProperty{
//   			BucketName: jsii.String("bucketName"),
//   			Key: jsii.String("key"),
//   		},
//   		OutputLocation: &S3LocationProperty{
//   			BucketName: jsii.String("bucketName"),
//   			Key: jsii.String("key"),
//   		},
//   		TransformerId: jsii.String("transformerId"),
//   		Type: &EdiTypeProperty{
//   			X12Details: &X12DetailsProperty{
//   				TransactionSet: jsii.String("transactionSet"),
//   				Version: jsii.String("version"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-capabilityconfiguration.html
//
type CfnCapability_CapabilityConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-capabilityconfiguration.html#cfn-b2bi-capability-capabilityconfiguration-edi
	//
	Edi interface{} `field:"required" json:"edi" yaml:"edi"`
}

