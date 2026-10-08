package awscdkredshiftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Options for S3 logging.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkredshiftalpha"
//
//   var vpc IVpc
//   var bucket IBucket
//
//
//   cluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
//   	MasterUser: &Login{
//   		MasterUsername: jsii.String("admin"),
//   	},
//   	Vpc: Vpc,
//   	Logging: awscdkredshiftalpha.ClusterLogging_S3(&S3LoggingOptions{
//   		Bucket: *Bucket,
//   		KeyPrefix: jsii.String("redshift-logs/"),
//   	}),
//   })
//
// Experimental.
type S3LoggingOptions struct {
	// The S3 bucket where the log files are stored.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The prefix applied to the log file names.
	// Default: - No prefix.
	//
	// Experimental.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

