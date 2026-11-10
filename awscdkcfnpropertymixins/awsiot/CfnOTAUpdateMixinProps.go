package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnOTAUpdatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnOTAUpdateMixinProps := &CfnOTAUpdateMixinProps{
//   	AdditionalParameters: map[string]*string{
//   		"additionalParametersKey": jsii.String("additionalParameters"),
//   	},
//   	AwsJobAbortConfig: &AwsJobAbortConfigProperty{
//   		AbortCriteriaList: []interface{}{
//   			&AwsJobAbortCriteriaProperty{
//   				Action: jsii.String("action"),
//   				FailureType: jsii.String("failureType"),
//   				MinNumberOfExecutedThings: jsii.Number(123),
//   				ThresholdPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	AwsJobExecutionsRolloutConfig: &AwsJobExecutionsRolloutConfigProperty{
//   		ExponentialRate: &AwsJobExponentialRolloutRateProperty{
//   			BaseRatePerMinute: jsii.Number(123),
//   			IncrementFactor: jsii.Number(123),
//   			RateIncreaseCriteria: &AwsJobRateIncreaseCriteriaProperty{
//   				NumberOfNotifiedThings: jsii.Number(123),
//   				NumberOfSucceededThings: jsii.Number(123),
//   			},
//   		},
//   		MaximumPerMinute: jsii.Number(123),
//   	},
//   	AwsJobPresignedUrlConfig: &AwsJobPresignedUrlConfigProperty{
//   		ExpiresInSec: jsii.Number(123),
//   	},
//   	AwsJobTimeoutConfig: &AwsJobTimeoutConfigProperty{
//   		InProgressTimeoutInMinutes: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	Files: []interface{}{
//   		&OTAUpdateFileProperty{
//   			Attributes: map[string]*string{
//   				"attributesKey": jsii.String("attributes"),
//   			},
//   			CodeSigning: &CodeSigningProperty{
//   				AwsSignerJobId: jsii.String("awsSignerJobId"),
//   				CustomCodeSigning: &CustomCodeSigningProperty{
//   					CertificateChain: &CodeSigningCertificateChainProperty{
//   						CertificateName: jsii.String("certificateName"),
//   						InlineDocument: jsii.String("inlineDocument"),
//   					},
//   					HashAlgorithm: jsii.String("hashAlgorithm"),
//   					Signature: &CodeSigningSignatureProperty{
//   						InlineDocument: jsii.String("inlineDocument"),
//   					},
//   					SignatureAlgorithm: jsii.String("signatureAlgorithm"),
//   				},
//   				StartSigningJobParameter: &StartSigningJobParameterProperty{
//   					Destination: &DestinationProperty{
//   						S3Destination: &S3DestinationProperty{
//   							Bucket: jsii.String("bucket"),
//   							Prefix: jsii.String("prefix"),
//   						},
//   					},
//   					SigningProfileName: jsii.String("signingProfileName"),
//   					SigningProfileParameter: &SigningProfileParameterProperty{
//   						CertificateArn: jsii.String("certificateArn"),
//   						CertificatePathOnDevice: jsii.String("certificatePathOnDevice"),
//   						Platform: jsii.String("platform"),
//   					},
//   				},
//   			},
//   			FileLocation: &FileLocationProperty{
//   				S3Location: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Key: jsii.String("key"),
//   					Version: jsii.String("version"),
//   				},
//   				Stream: &StreamProperty{
//   					FileId: jsii.Number(123),
//   					StreamId: jsii.String("streamId"),
//   				},
//   			},
//   			FileName: jsii.String("fileName"),
//   			FileType: jsii.Number(123),
//   			FileVersion: jsii.String("fileVersion"),
//   		},
//   	},
//   	OtaUpdateId: jsii.String("otaUpdateId"),
//   	Protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Targets: []*string{
//   		jsii.String("targets"),
//   	},
//   	TargetSelection: jsii.String("targetSelection"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html
//
type CfnOTAUpdateMixinProps struct {
	// A list of additional OTA update parameters, which are name-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-additionalparameters
	//
	AdditionalParameters interface{} `field:"optional" json:"additionalParameters" yaml:"additionalParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-awsjobabortconfig
	//
	AwsJobAbortConfig interface{} `field:"optional" json:"awsJobAbortConfig" yaml:"awsJobAbortConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-awsjobexecutionsrolloutconfig
	//
	AwsJobExecutionsRolloutConfig interface{} `field:"optional" json:"awsJobExecutionsRolloutConfig" yaml:"awsJobExecutionsRolloutConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-awsjobpresignedurlconfig
	//
	AwsJobPresignedUrlConfig interface{} `field:"optional" json:"awsJobPresignedUrlConfig" yaml:"awsJobPresignedUrlConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-awsjobtimeoutconfig
	//
	AwsJobTimeoutConfig interface{} `field:"optional" json:"awsJobTimeoutConfig" yaml:"awsJobTimeoutConfig"`
	// The description of the OTA update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The files to be streamed by the OTA update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-files
	//
	Files interface{} `field:"optional" json:"files" yaml:"files"`
	// The ID of the OTA update to be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-otaupdateid
	//
	OtaUpdateId *string `field:"optional" json:"otaUpdateId" yaml:"otaUpdateId"`
	// The protocol used to transfer the OTA update image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-protocols
	//
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// The IAM role that grants AWS IoT Core access to the Amazon S3, IoT jobs and AWS Code Signing resources to create an OTA update job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Metadata which can be used to manage updates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The devices targeted to receive OTA updates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-targets
	//
	Targets *[]*string `field:"optional" json:"targets" yaml:"targets"`
	// Specifies whether the update will continue to run (CONTINUOUS), or will be complete after all things specified as targets have completed the update (SNAPSHOT).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html#cfn-iot-otaupdate-targetselection
	//
	TargetSelection *string `field:"optional" json:"targetSelection" yaml:"targetSelection"`
}

