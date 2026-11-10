package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   codeSigningProperty := &CodeSigningProperty{
//   	AwsSignerJobId: jsii.String("awsSignerJobId"),
//   	CustomCodeSigning: &CustomCodeSigningProperty{
//   		CertificateChain: &CodeSigningCertificateChainProperty{
//   			CertificateName: jsii.String("certificateName"),
//   			InlineDocument: jsii.String("inlineDocument"),
//   		},
//   		HashAlgorithm: jsii.String("hashAlgorithm"),
//   		Signature: &CodeSigningSignatureProperty{
//   			InlineDocument: jsii.String("inlineDocument"),
//   		},
//   		SignatureAlgorithm: jsii.String("signatureAlgorithm"),
//   	},
//   	StartSigningJobParameter: &StartSigningJobParameterProperty{
//   		Destination: &DestinationProperty{
//   			S3Destination: &S3DestinationProperty{
//   				Bucket: jsii.String("bucket"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   		SigningProfileName: jsii.String("signingProfileName"),
//   		SigningProfileParameter: &SigningProfileParameterProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			CertificatePathOnDevice: jsii.String("certificatePathOnDevice"),
//   			Platform: jsii.String("platform"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-codesigning.html
//
type CfnOTAUpdatePropsMixin_CodeSigningProperty struct {
	// The ID of the AWSSignerJob which was created to sign the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-codesigning.html#cfn-iot-otaupdate-codesigning-awssignerjobid
	//
	AwsSignerJobId *string `field:"optional" json:"awsSignerJobId" yaml:"awsSignerJobId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-codesigning.html#cfn-iot-otaupdate-codesigning-customcodesigning
	//
	CustomCodeSigning interface{} `field:"optional" json:"customCodeSigning" yaml:"customCodeSigning"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-codesigning.html#cfn-iot-otaupdate-codesigning-startsigningjobparameter
	//
	StartSigningJobParameter interface{} `field:"optional" json:"startSigningJobParameter" yaml:"startSigningJobParameter"`
}

