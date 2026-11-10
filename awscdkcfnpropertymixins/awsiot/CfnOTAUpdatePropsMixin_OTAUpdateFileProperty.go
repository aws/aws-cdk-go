package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   oTAUpdateFileProperty := &OTAUpdateFileProperty{
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	CodeSigning: &CodeSigningProperty{
//   		AwsSignerJobId: jsii.String("awsSignerJobId"),
//   		CustomCodeSigning: &CustomCodeSigningProperty{
//   			CertificateChain: &CodeSigningCertificateChainProperty{
//   				CertificateName: jsii.String("certificateName"),
//   				InlineDocument: jsii.String("inlineDocument"),
//   			},
//   			HashAlgorithm: jsii.String("hashAlgorithm"),
//   			Signature: &CodeSigningSignatureProperty{
//   				InlineDocument: jsii.String("inlineDocument"),
//   			},
//   			SignatureAlgorithm: jsii.String("signatureAlgorithm"),
//   		},
//   		StartSigningJobParameter: &StartSigningJobParameterProperty{
//   			Destination: &DestinationProperty{
//   				S3Destination: &S3DestinationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Prefix: jsii.String("prefix"),
//   				},
//   			},
//   			SigningProfileName: jsii.String("signingProfileName"),
//   			SigningProfileParameter: &SigningProfileParameterProperty{
//   				CertificateArn: jsii.String("certificateArn"),
//   				CertificatePathOnDevice: jsii.String("certificatePathOnDevice"),
//   				Platform: jsii.String("platform"),
//   			},
//   		},
//   	},
//   	FileLocation: &FileLocationProperty{
//   		S3Location: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   			Version: jsii.String("version"),
//   		},
//   		Stream: &StreamProperty{
//   			FileId: jsii.Number(123),
//   			StreamId: jsii.String("streamId"),
//   		},
//   	},
//   	FileName: jsii.String("fileName"),
//   	FileType: jsii.Number(123),
//   	FileVersion: jsii.String("fileVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-otaupdatefile.html
//
type CfnOTAUpdatePropsMixin_OTAUpdateFileProperty struct {
	// A list of name-attribute pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-otaupdatefile.html#cfn-iot-otaupdate-otaupdatefile-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-otaupdatefile.html#cfn-iot-otaupdate-otaupdatefile-codesigning
	//
	CodeSigning interface{} `field:"optional" json:"codeSigning" yaml:"codeSigning"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-otaupdatefile.html#cfn-iot-otaupdate-otaupdatefile-filelocation
	//
	FileLocation interface{} `field:"optional" json:"fileLocation" yaml:"fileLocation"`
	// The name of the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-otaupdatefile.html#cfn-iot-otaupdate-otaupdatefile-filename
	//
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// An integer value you can include in the job document to allow your devices to identify the type of file received from the cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-otaupdatefile.html#cfn-iot-otaupdate-otaupdatefile-filetype
	//
	FileType *float64 `field:"optional" json:"fileType" yaml:"fileType"`
	// The file version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-otaupdatefile.html#cfn-iot-otaupdate-otaupdatefile-fileversion
	//
	FileVersion *string `field:"optional" json:"fileVersion" yaml:"fileVersion"`
}

