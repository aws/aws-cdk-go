package awssagemaker


// Configuration for mounting an Amazon FSx Lustre file system to the instances in the SageMaker HyperPod cluster instance group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterFsxLustreConfigProperty := &ClusterFsxLustreConfigProperty{
//   	DnsName: jsii.String("dnsName"),
//   	MountName: jsii.String("mountName"),
//
//   	// the properties below are optional
//   	MountPath: jsii.String("mountPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterfsxlustreconfig.html
//
type CfnCluster_ClusterFsxLustreConfigProperty struct {
	// The DNS name of the FSx for Lustre file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterfsxlustreconfig.html#cfn-sagemaker-cluster-clusterfsxlustreconfig-dnsname
	//
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// The mount name of the FSx for Lustre file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterfsxlustreconfig.html#cfn-sagemaker-cluster-clusterfsxlustreconfig-mountname
	//
	MountName *string `field:"required" json:"mountName" yaml:"mountName"`
	// The mount path for the FSx for Lustre file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterfsxlustreconfig.html#cfn-sagemaker-cluster-clusterfsxlustreconfig-mountpath
	//
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

