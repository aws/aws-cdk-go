// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Logging bucket and S3 prefix combination.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"))
//   bucket := s3.Bucket_FromBucketName(this, jsii.String("bucket"), jsii.String("logging-bucket"))
//
//   cluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
//   	MasterUser: &Login{
//   		MasterUsername: jsii.String("admin"),
//   	},
//   	Vpc: Vpc,
//   	LoggingProperties: &LoggingProperties{
//   		LoggingBucket: bucket,
//   		LoggingKeyPrefix: jsii.String("prefix"),
//   	},
//   })
//
// Experimental.
type LoggingProperties struct {
	// Bucket to send logs to.
	//
	// Logging information includes queries and connection attempts, for the specified Amazon Redshift cluster.
	// Experimental.
	LoggingBucket awss3.IBucket `field:"required" json:"loggingBucket" yaml:"loggingBucket"`
	// Prefix used for logging.
	// Experimental.
	LoggingKeyPrefix *string `field:"required" json:"loggingKeyPrefix" yaml:"loggingKeyPrefix"`
}

