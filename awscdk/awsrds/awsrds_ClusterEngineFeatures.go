package awsrds


// Represents Database Engine features.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterEngineFeatures := &clusterEngineFeatures{
//   	s3Export: jsii.String("s3Export"),
//   	s3Import: jsii.String("s3Import"),
//   }
//
type ClusterEngineFeatures struct {
	// Feature name for the DB instance that the IAM role to export to S3 bucket is to be associated with.
	S3Export *string `field:"optional" json:"s3Export" yaml:"s3Export"`
	// Feature name for the DB instance that the IAM role to access the S3 bucket for import is to be associated with.
	S3Import *string `field:"optional" json:"s3Import" yaml:"s3Import"`
}

