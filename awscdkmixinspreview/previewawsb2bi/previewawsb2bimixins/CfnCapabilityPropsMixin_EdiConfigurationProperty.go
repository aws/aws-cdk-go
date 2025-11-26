package previewawsb2bimixins


// Specifies the details for the EDI (electronic data interchange) transformation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ediConfigurationProperty := &EdiConfigurationProperty{
//   	CapabilityDirection: jsii.String("capabilityDirection"),
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
type CfnCapabilityPropsMixin_EdiConfigurationProperty struct {
	// Specifies whether this is capability is for inbound or outbound transformations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-capabilitydirection
	//
	CapabilityDirection *string `field:"optional" json:"capabilityDirection" yaml:"capabilityDirection"`
	// Contains the Amazon S3 bucket and prefix for the location of the input file, which is contained in an `S3Location` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-inputlocation
	//
	InputLocation interface{} `field:"optional" json:"inputLocation" yaml:"inputLocation"`
	// Contains the Amazon S3 bucket and prefix for the location of the output file, which is contained in an `S3Location` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-outputlocation
	//
	OutputLocation interface{} `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	// Returns the system-assigned unique identifier for the transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-transformerid
	//
	TransformerId *string `field:"optional" json:"transformerId" yaml:"transformerId"`
	// Returns the type of the capability.
	//
	// Currently, only `edi` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-type
	//
	Type interface{} `field:"optional" json:"type" yaml:"type"`
}

