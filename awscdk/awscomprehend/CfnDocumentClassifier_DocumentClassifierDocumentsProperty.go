package awscomprehend


// The location of the training documents.
//
// This parameter is required in a request to create a native document model.
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
type CfnDocumentClassifier_DocumentClassifierDocumentsProperty struct {
	// The S3 URI location of the training documents specified in the S3Uri CSV file.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The S3 URI location of the test documents included in the TestS3Uri CSV file.
	//
	// This field is not required if you do not specify a test CSV file.
	TestS3Uri *string `field:"optional" json:"testS3Uri" yaml:"testS3Uri"`
}

