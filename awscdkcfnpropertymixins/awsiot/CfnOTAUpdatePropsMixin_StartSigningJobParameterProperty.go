package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   startSigningJobParameterProperty := &StartSigningJobParameterProperty{
//   	Destination: &DestinationProperty{
//   		S3Destination: &S3DestinationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	SigningProfileName: jsii.String("signingProfileName"),
//   	SigningProfileParameter: &SigningProfileParameterProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		CertificatePathOnDevice: jsii.String("certificatePathOnDevice"),
//   		Platform: jsii.String("platform"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-startsigningjobparameter.html
//
type CfnOTAUpdatePropsMixin_StartSigningJobParameterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-startsigningjobparameter.html#cfn-iot-otaupdate-startsigningjobparameter-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The code-signing profile name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-startsigningjobparameter.html#cfn-iot-otaupdate-startsigningjobparameter-signingprofilename
	//
	SigningProfileName *string `field:"optional" json:"signingProfileName" yaml:"signingProfileName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-startsigningjobparameter.html#cfn-iot-otaupdate-startsigningjobparameter-signingprofileparameter
	//
	SigningProfileParameter interface{} `field:"optional" json:"signingProfileParameter" yaml:"signingProfileParameter"`
}

