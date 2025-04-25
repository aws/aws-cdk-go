package awsssm


// Identifying information about a document attachment, including the file name and a key-value pair that identifies the location of an attachment to a document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attachmentsSourceProperty := &AttachmentsSourceProperty{
//   	Key: jsii.String("key"),
//   	Name: jsii.String("name"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html
//
type CfnDocument_AttachmentsSourceProperty struct {
	// The key of a key-value pair that identifies the location of an attachment to a document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The name of the document attachment file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of a key-value pair that identifies the location of an attachment to a document.
	//
	// The format for *Value* depends on the type of key you specify.
	//
	// - For the key *SourceUrl* , the value is an S3 bucket location. For example:
	//
	// `"Values": [ "s3://amzn-s3-demo-bucket/my-prefix" ]`
	// - For the key *S3FileUrl* , the value is a file in an S3 bucket. For example:
	//
	// `"Values": [ "s3://amzn-s3-demo-bucket/my-prefix/my-file.py" ]`
	// - For the key *AttachmentReference* , the value is constructed from the name of another SSM document in your account, a version number of that document, and a file attached to that document version that you want to reuse. For example:
	//
	// `"Values": [ "MyOtherDocument/3/my-other-file.py" ]`
	//
	// However, if the SSM document is shared with you from another account, the full SSM document ARN must be specified instead of the document name only. For example:
	//
	// `"Values": [ "arn:aws:ssm:us-east-2:111122223333:document/OtherAccountDocument/3/their-file.py" ]`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

