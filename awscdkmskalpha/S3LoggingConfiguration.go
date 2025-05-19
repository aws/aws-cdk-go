package awscdkmskalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Details of the Amazon S3 destination for broker logs.
//
// Example:
//   var vpc vpc
//   var bucket iBucket
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
//   	ClusterName: jsii.String("myCluster"),
//   	KafkaVersion: msk.KafkaVersion_V3_9_X(),
//   	Vpc: Vpc,
//   	Logging: &BrokerLogging{
//   		S3: &S3LoggingConfiguration{
//   			Bucket: *Bucket,
//   		},
//   	},
//   })
//
// Experimental.
type S3LoggingConfiguration struct {
	// The S3 bucket that is the destination for broker logs.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 prefix that is the destination for broker logs.
	// Default: - no prefix.
	//
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

