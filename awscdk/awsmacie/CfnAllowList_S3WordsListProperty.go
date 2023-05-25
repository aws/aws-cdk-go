package awsmacie


// Specifies the location and name of an Amazon Simple Storage Service ( Amazon S3 ) object that lists specific, predefined text to ignore when inspecting data sources for sensitive data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3WordsListProperty := &S3WordsListProperty{
//   	BucketName: jsii.String("bucketName"),
//   	ObjectKey: jsii.String("objectKey"),
//   }
//
type CfnAllowList_S3WordsListProperty struct {
	// The full name of the S3 bucket that contains the object.
	//
	// This value correlates to the `Name` field of a bucket's properties in Amazon S3 .
	//
	// This value is case sensitive. In addition, don't use wildcard characters or specify partial values for the name.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The full name of the S3 object.
	//
	// This value correlates to the `Key` field of an object's properties in Amazon S3 . If the name includes a path, include the complete path. For example, `AllowLists/Macie/MyList.txt` .
	//
	// This value is case sensitive. In addition, don't use wildcard characters or specify partial values for the name.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
}

