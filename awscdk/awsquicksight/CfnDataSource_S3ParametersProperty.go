package awsquicksight


// The parameters for S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ParametersProperty := &S3ParametersProperty{
//   	ManifestFileLocation: &ManifestFileLocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//
//   	// the properties below are optional
//   	RoleArn: jsii.String("roleArn"),
//   }
//
type CfnDataSource_S3ParametersProperty struct {
	// Location of the Amazon S3 manifest file.
	//
	// This is NULL if the manifest file was uploaded into Amazon QuickSight.
	ManifestFileLocation interface{} `field:"required" json:"manifestFileLocation" yaml:"manifestFileLocation"`
	// Use the `RoleArn` structure to override an account-wide role for a specific S3 data source.
	//
	// For example, say an account administrator has turned off all S3 access with an account-wide role. The administrator can then use `RoleArn` to bypass the account-wide role and allow S3 access for the single S3 data source that is specified in the structure, even if the account-wide role forbidding S3 access is still active.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

