package awscodepipelineactions


// The CodePipeline variables emitted by the S3 source Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SourceVariables := &s3SourceVariables{
//   	eTag: jsii.String("eTag"),
//   	versionId: jsii.String("versionId"),
//   }
//
type S3SourceVariables struct {
	// The e-tag of the S3 version of the object that triggered the build.
	ETag *string `field:"required" json:"eTag" yaml:"eTag"`
	// The identifier of the S3 version of the object that triggered the build.
	VersionId *string `field:"required" json:"versionId" yaml:"versionId"`
}

