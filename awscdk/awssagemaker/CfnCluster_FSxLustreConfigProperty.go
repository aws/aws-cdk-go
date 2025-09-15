package awssagemaker


// Configuration settings for an Amazon FSx for Lustre file system to be used with the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fSxLustreConfigProperty := &FSxLustreConfigProperty{
//   	PerUnitStorageThroughput: jsii.Number(123),
//   	SizeInGiB: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-fsxlustreconfig.html
//
type CfnCluster_FSxLustreConfigProperty struct {
	// The throughput capacity of the Amazon FSx for Lustre file system, measured in MB/s per TiB of storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-fsxlustreconfig.html#cfn-sagemaker-cluster-fsxlustreconfig-perunitstoragethroughput
	//
	PerUnitStorageThroughput *float64 `field:"required" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// The storage capacity of the Amazon FSx for Lustre file system, specified in gibibytes (GiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-fsxlustreconfig.html#cfn-sagemaker-cluster-fsxlustreconfig-sizeingib
	//
	SizeInGiB *float64 `field:"required" json:"sizeInGiB" yaml:"sizeInGiB"`
}

