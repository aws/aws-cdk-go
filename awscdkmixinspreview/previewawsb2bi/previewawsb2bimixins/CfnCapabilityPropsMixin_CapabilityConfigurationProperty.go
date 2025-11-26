package previewawsb2bimixins


// A capability object.
//
// Currently, only EDI (electronic data interchange) capabilities are supported. A trading capability contains the information required to transform incoming EDI documents into JSON or XML outputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   capabilityConfigurationProperty := &CapabilityConfigurationProperty{
//   	Edi: &EdiConfigurationProperty{
//   		CapabilityDirection: jsii.String("capabilityDirection"),
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
type CfnCapabilityPropsMixin_CapabilityConfigurationProperty struct {
	// An EDI (electronic data interchange) configuration object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-capabilityconfiguration.html#cfn-b2bi-capability-capabilityconfiguration-edi
	//
	Edi interface{} `field:"optional" json:"edi" yaml:"edi"`
}

