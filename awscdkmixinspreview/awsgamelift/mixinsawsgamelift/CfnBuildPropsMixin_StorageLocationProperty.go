package mixinsawsgamelift


// The location in Amazon S3 where build or script files are stored for access by Amazon GameLift.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   storageLocationProperty := &StorageLocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   	ObjectVersion: jsii.String("objectVersion"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html
//
type CfnBuildPropsMixin_StorageLocationProperty struct {
	// An Amazon S3 bucket identifier. The name of the S3 bucket.
	//
	// > Amazon GameLift doesn't support uploading from Amazon S3 buckets with names that contain a dot (.).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storagelocation-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The name of the zip file that contains the build files or script files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storagelocation-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A version of a stored file to retrieve, if the object versioning feature is turned on for the S3 bucket.
	//
	// Use this parameter to specify a specific version. If this parameter isn't set, Amazon GameLift Servers retrieves the latest version of the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storagelocation-objectversion
	//
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
	// The ARNfor an IAM role that allows Amazon GameLift to access the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storagelocation-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

