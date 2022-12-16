package awsopensearchservice


// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the OpenSearch Service domain.
//
// For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eBSOptionsProperty := &eBSOptionsProperty{
//   	ebsEnabled: jsii.Boolean(false),
//   	iops: jsii.Number(123),
//   	throughput: jsii.Number(123),
//   	volumeSize: jsii.Number(123),
//   	volumeType: jsii.String("volumeType"),
//   }
//
type CfnDomain_EBSOptionsProperty struct {
	// Specifies whether Amazon EBS volumes are attached to data nodes in the OpenSearch Service domain.
	EbsEnabled interface{} `field:"optional" json:"ebsEnabled" yaml:"ebsEnabled"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	//
	// This property applies only to the Provisioned IOPS (SSD) EBS volume type.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// `CfnDomain.EBSOptionsProperty.Throughput`.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The size (in GiB) of the EBS volume for each data node.
	//
	// The minimum and maximum size of an EBS volume depends on the EBS volume type and the instance type to which it is attached. For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// The EBS volume type to use with the OpenSearch Service domain, such as standard, gp2, or io1.
	//
	// For more information about each type, see [Amazon EBS volume types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the *Amazon EC2 User Guide for Linux Instances* .
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

