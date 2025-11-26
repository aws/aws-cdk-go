package previewawsopensearchservicemixins


// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the OpenSearch Service domain.
//
// For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eBSOptionsProperty := &EBSOptionsProperty{
//   	EbsEnabled: jsii.Boolean(false),
//   	Iops: jsii.Number(123),
//   	Throughput: jsii.Number(123),
//   	VolumeSize: jsii.Number(123),
//   	VolumeType: jsii.String("volumeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-ebsoptions.html
//
type CfnDomainPropsMixin_EBSOptionsProperty struct {
	// Specifies whether Amazon EBS volumes are attached to data nodes in the OpenSearch Service domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-ebsoptions.html#cfn-opensearchservice-domain-ebsoptions-ebsenabled
	//
	EbsEnabled interface{} `field:"optional" json:"ebsEnabled" yaml:"ebsEnabled"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	//
	// This property applies only to the `gp3` and provisioned IOPS EBS volume types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-ebsoptions.html#cfn-opensearchservice-domain-ebsoptions-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The throughput (in MiB/s) of the EBS volumes attached to data nodes.
	//
	// Applies only to the `gp3` volume type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-ebsoptions.html#cfn-opensearchservice-domain-ebsoptions-throughput
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The size (in GiB) of the EBS volume for each data node.
	//
	// The minimum and maximum size of an EBS volume depends on the EBS volume type and the instance type to which it is attached. For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-ebsoptions.html#cfn-opensearchservice-domain-ebsoptions-volumesize
	//
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// The EBS volume type to use with the OpenSearch Service domain.
	//
	// If you choose `gp3` , you must also specify values for `Iops` and `Throughput` . For more information about each type, see [Amazon EBS volume types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the *Amazon EC2 User Guide for Linux Instances* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-ebsoptions.html#cfn-opensearchservice-domain-ebsoptions-volumetype
	//
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

