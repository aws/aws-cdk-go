package awsquicksight


// The parameters for S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ParametersProperty := &s3ParametersProperty{
//   	manifestFileLocation: &manifestFileLocationProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnDataSource_S3ParametersProperty struct {
	// Location of the Amazon S3 manifest file.
	//
	// This is NULL if the manifest file was uploaded into Amazon QuickSight.
	ManifestFileLocation interface{} `field:"required" json:"manifestFileLocation" yaml:"manifestFileLocation"`
}

