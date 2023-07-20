package awscomprehend


// Provide the location for output data from a custom classifier job.
//
// This field is mandatory if you are training a native document model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentClassifierOutputDataConfigProperty := &DocumentClassifierOutputDataConfigProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifieroutputdataconfig.html
//
type CfnDocumentClassifier_DocumentClassifierOutputDataConfigProperty struct {
	// ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses to encrypt the output results from an analysis job.
	//
	// The KmsKeyId can be one of the following formats:
	//
	// - KMS Key ID: `"1234abcd-12ab-34cd-56ef-1234567890ab"`
	// - Amazon Resource Name (ARN) of a KMS Key: `"arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`
	// - KMS Key Alias: `"alias/ExampleAlias"`
	// - ARN of a KMS Key Alias: `"arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifieroutputdataconfig.html#cfn-comprehend-documentclassifier-documentclassifieroutputdataconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// When you use the `OutputDataConfig` object while creating a custom classifier, you specify the Amazon S3 location where you want to write the confusion matrix and other output files.
	//
	// The URI must be in the same Region as the API endpoint that you are calling. The location is used as the prefix for the actual location of this output file.
	//
	// When the custom classifier job is finished, the service creates the output file in a directory specific to the job. The `S3Uri` field contains the location of the output file, called `output.tar.gz` . It is a compressed archive that contains the confusion matrix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifieroutputdataconfig.html#cfn-comprehend-documentclassifier-documentclassifieroutputdataconfig-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

