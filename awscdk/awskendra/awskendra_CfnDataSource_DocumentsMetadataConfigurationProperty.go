package awskendra


// Document metadata files that contain information such as the document access control information, source URI, document author, and custom attributes.
//
// Each metadata file contains metadata about a single document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentsMetadataConfigurationProperty := &documentsMetadataConfigurationProperty{
//   	s3Prefix: jsii.String("s3Prefix"),
//   }
//
type CfnDataSource_DocumentsMetadataConfigurationProperty struct {
	// A prefix used to filter metadata configuration files in the AWS S3 bucket.
	//
	// The S3 bucket might contain multiple metadata files. Use `S3Prefix` to include only the desired metadata files.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
}

