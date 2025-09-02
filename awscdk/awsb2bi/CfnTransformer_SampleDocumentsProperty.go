package awsb2bi


// Describes a structure that contains the Amazon S3 bucket and an array of the corresponding keys used to identify the location for your sample documents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sampleDocumentsProperty := &SampleDocumentsProperty{
//   	BucketName: jsii.String("bucketName"),
//   	Keys: []interface{}{
//   		&SampleDocumentKeysProperty{
//   			Input: jsii.String("input"),
//   			Output: jsii.String("output"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-sampledocuments.html
//
type CfnTransformer_SampleDocumentsProperty struct {
	// Contains the Amazon S3 bucket that is used to hold your sample documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-sampledocuments.html#cfn-b2bi-transformer-sampledocuments-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Contains an array of the Amazon S3 keys used to identify the location for your sample documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-sampledocuments.html#cfn-b2bi-transformer-sampledocuments-keys
	//
	Keys interface{} `field:"required" json:"keys" yaml:"keys"`
}

