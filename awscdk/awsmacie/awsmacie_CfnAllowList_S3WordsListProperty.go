package awsmacie


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3WordsListProperty := &s3WordsListProperty{
//   	bucketName: jsii.String("bucketName"),
//   	objectKey: jsii.String("objectKey"),
//   }
//
type CfnAllowList_S3WordsListProperty struct {
	// `CfnAllowList.S3WordsListProperty.BucketName`.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// `CfnAllowList.S3WordsListProperty.ObjectKey`.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
}

