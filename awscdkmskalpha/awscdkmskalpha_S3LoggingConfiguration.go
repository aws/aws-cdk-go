// The CDK Construct Library for AWS::MSK
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
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	logging: &brokerLogging{
//   		s3: &s3LoggingConfiguration{
//   			bucket: bucket,
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
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

