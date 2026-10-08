package awscdkredshiftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkredshiftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The logging configuration for a Redshift cluster.
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
type ClusterLogging interface {
}

// The jsii proxy struct for ClusterLogging
type jsiiProxy_ClusterLogging struct {
	_ byte // padding
}

// Experimental.
func NewClusterLogging_Override(c ClusterLogging) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-redshift-alpha.ClusterLogging",
		nil, // no parameters
		c,
	)
}

// Send logs to CloudWatch.
// Experimental.
func ClusterLogging_CloudWatch(options *CloudWatchLoggingOptions) ClusterLogging {
	_init_.Initialize()

	if err := validateClusterLogging_CloudWatchParameters(options); err != nil {
		panic(err)
	}
	var returns ClusterLogging

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-redshift-alpha.ClusterLogging",
		"cloudWatch",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Send logs to S3.
// Experimental.
func ClusterLogging_S3(options *S3LoggingOptions) ClusterLogging {
	_init_.Initialize()

	if err := validateClusterLogging_S3Parameters(options); err != nil {
		panic(err)
	}
	var returns ClusterLogging

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-redshift-alpha.ClusterLogging",
		"s3",
		[]interface{}{options},
		&returns,
	)

	return returns
}

