package awss3express


// Specifies the default server-side encryption to apply to new objects in the bucket.
//
// If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverSideEncryptionByDefaultProperty := &ServerSideEncryptionByDefaultProperty{
//   	SseAlgorithm: jsii.String("sseAlgorithm"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-serversideencryptionbydefault.html
//
type CfnDirectoryBucket_ServerSideEncryptionByDefaultProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-serversideencryptionbydefault.html#cfn-s3express-directorybucket-serversideencryptionbydefault-ssealgorithm
	//
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
}

