package awsb2bi


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-sampledocuments.html#cfn-b2bi-transformer-sampledocuments-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-sampledocuments.html#cfn-b2bi-transformer-sampledocuments-keys
	//
	Keys interface{} `field:"required" json:"keys" yaml:"keys"`
}

