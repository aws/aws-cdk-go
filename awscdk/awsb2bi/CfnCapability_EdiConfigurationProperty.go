package awsb2bi


// Specifies the details for the EDI (electronic data interchange) transformation.
//
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
//
//   	// the properties below are optional
//   	CapabilityDirection: jsii.String("capabilityDirection"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html
//
type CfnCapability_EdiConfigurationProperty struct {
	// Contains the Amazon S3 bucket and prefix for the location of the input file, which is contained in an `S3Location` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-inputlocation
	//
	InputLocation interface{} `field:"required" json:"inputLocation" yaml:"inputLocation"`
	// Contains the Amazon S3 bucket and prefix for the location of the output file, which is contained in an `S3Location` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-outputlocation
	//
	OutputLocation interface{} `field:"required" json:"outputLocation" yaml:"outputLocation"`
	// Returns the system-assigned unique identifier for the transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-transformerid
	//
	TransformerId *string `field:"required" json:"transformerId" yaml:"transformerId"`
	// Returns the type of the capability.
	//
	// Currently, only `edi` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-type
	//
	Type interface{} `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-ediconfiguration.html#cfn-b2bi-capability-ediconfiguration-capabilitydirection
	//
	CapabilityDirection *string `field:"optional" json:"capabilityDirection" yaml:"capabilityDirection"`
}

