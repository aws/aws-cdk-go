package awssagemaker


// Specifies the configuration of your endpoint for model monitor data capture.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCaptureConfigProperty := &DataCaptureConfigProperty{
//   	CaptureOptions: []interface{}{
//   		&CaptureOptionProperty{
//   			CaptureMode: jsii.String("captureMode"),
//   		},
//   	},
//   	DestinationS3Uri: jsii.String("destinationS3Uri"),
//   	InitialSamplingPercentage: jsii.Number(123),
//
//   	// the properties below are optional
//   	CaptureContentTypeHeader: &CaptureContentTypeHeaderProperty{
//   		CsvContentTypes: []*string{
//   			jsii.String("csvContentTypes"),
//   		},
//   		JsonContentTypes: []*string{
//   			jsii.String("jsonContentTypes"),
//   		},
//   	},
//   	EnableCapture: jsii.Boolean(false),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html
//
type CfnEndpointConfig_DataCaptureConfigProperty struct {
	// Specifies whether the endpoint captures input data to your model, output data from your model, or both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig-captureoptions
	//
	CaptureOptions interface{} `field:"required" json:"captureOptions" yaml:"captureOptions"`
	// The S3 bucket where model monitor stores captured data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig-destinations3uri
	//
	DestinationS3Uri *string `field:"required" json:"destinationS3Uri" yaml:"destinationS3Uri"`
	// The percentage of data to capture.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig-initialsamplingpercentage
	//
	InitialSamplingPercentage *float64 `field:"required" json:"initialSamplingPercentage" yaml:"initialSamplingPercentage"`
	// A list of the JSON and CSV content type that the endpoint captures.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig-capturecontenttypeheader
	//
	CaptureContentTypeHeader interface{} `field:"optional" json:"captureContentTypeHeader" yaml:"captureContentTypeHeader"`
	// Set to `True` to enable data capture.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig-enablecapture
	//
	EnableCapture interface{} `field:"optional" json:"enableCapture" yaml:"enableCapture"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt the captured data at rest using Amazon S3 server-side encryption.
	//
	// The KmsKeyId can be any of the following formats: Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab Key ARN: arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab Alias name: alias/ExampleAlias Alias name ARN: arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias If you don't provide a KMS key ID, Amazon SageMaker uses the default KMS key for Amazon S3 for your role's account. For more information, see KMS-Managed Encryption Keys (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) in the Amazon Simple Storage Service Developer Guide. The KMS key policy must grant permission to the IAM role that you specify in your CreateModel (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateModel.html) request. For more information, see Using Key Policies in AWS KMS (http://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html) in the AWS Key Management Service Developer Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

