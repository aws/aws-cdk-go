package awscomprehend


// The location of the training documents.
//
// This parameter is required in a request to create a semi-structured document classification model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentClassifierDocumentsProperty := &DocumentClassifierDocumentsProperty{
//   	S3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	TestS3Uri: jsii.String("testS3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifierdocuments.html
//
type CfnDocumentClassifier_DocumentClassifierDocumentsProperty struct {
	// The S3 URI location of the training documents specified in the S3Uri CSV file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifierdocuments.html#cfn-comprehend-documentclassifier-documentclassifierdocuments-s3uri
	//
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The S3 URI location of the test documents included in the TestS3Uri CSV file.
	//
	// This field is not required if you do not specify a test CSV file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifierdocuments.html#cfn-comprehend-documentclassifier-documentclassifierdocuments-tests3uri
	//
	TestS3Uri *string `field:"optional" json:"testS3Uri" yaml:"testS3Uri"`
}

