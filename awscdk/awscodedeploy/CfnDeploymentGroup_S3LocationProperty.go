package awscodedeploy


// `S3Location` is a property of the [CodeDeploy DeploymentGroup Revision](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html) property that specifies the location of an application revision that is stored in Amazon Simple Storage Service ( Amazon S3 ).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	BundleType: jsii.String("bundleType"),
//   	ETag: jsii.String("eTag"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-s3location.html
//
type CfnDeploymentGroup_S3LocationProperty struct {
	// The name of the Amazon S3 bucket where the application revision is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-s3location.html#cfn-codedeploy-deploymentgroup-s3location-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the Amazon S3 object that represents the bundled artifacts for the application revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-s3location.html#cfn-codedeploy-deploymentgroup-s3location-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The file type of the application revision. Must be one of the following:.
	//
	// - JSON
	// - tar: A tar archive file.
	// - tgz: A compressed tar archive file.
	// - YAML
	// - zip: A zip archive file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-s3location.html#cfn-codedeploy-deploymentgroup-s3location-bundletype
	//
	BundleType *string `field:"optional" json:"bundleType" yaml:"bundleType"`
	// The ETag of the Amazon S3 object that represents the bundled artifacts for the application revision.
	//
	// If the ETag is not specified as an input parameter, ETag validation of the object is skipped.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-s3location.html#cfn-codedeploy-deploymentgroup-s3location-etag
	//
	ETag *string `field:"optional" json:"eTag" yaml:"eTag"`
	// A specific version of the Amazon S3 object that represents the bundled artifacts for the application revision.
	//
	// If the version is not specified, the system uses the most recent version by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-s3location.html#cfn-codedeploy-deploymentgroup-s3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

