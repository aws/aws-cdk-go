package awsappflow


// The connector-specific profile properties when using Amazon Redshift.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftConnectorProfilePropertiesProperty := &redshiftConnectorProfilePropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//   	databaseUrl: jsii.String("databaseUrl"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   }
//
type CfnConnectorProfile_RedshiftConnectorProfilePropertiesProperty struct {
	// A name for the associated Amazon S3 bucket.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The JDBC URL of the Amazon Redshift cluster.
	DatabaseUrl *string `field:"required" json:"databaseUrl" yaml:"databaseUrl"`
	// The Amazon Resource Name (ARN) of the IAM role.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The object key for the destination bucket in which Amazon AppFlow places the files.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

