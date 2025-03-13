package awspcs


// Properties for defining a `CfnComputeNodeGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnComputeNodeGroupProps := &CfnComputeNodeGroupProps{
//   	ClusterId: jsii.String("clusterId"),
//   	CustomLaunchTemplate: &CustomLaunchTemplateProperty{
//   		Id: jsii.String("id"),
//   		Version: jsii.String("version"),
//   	},
//   	IamInstanceProfileArn: jsii.String("iamInstanceProfileArn"),
//   	InstanceConfigs: []interface{}{
//   		&InstanceConfigProperty{
//   			InstanceType: jsii.String("instanceType"),
//   		},
//   	},
//   	ScalingConfiguration: &ScalingConfigurationProperty{
//   		MaxInstanceCount: jsii.Number(123),
//   		MinInstanceCount: jsii.Number(123),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	AmiId: jsii.String("amiId"),
//   	Name: jsii.String("name"),
//   	PurchaseOption: jsii.String("purchaseOption"),
//   	SlurmConfiguration: &SlurmConfigurationProperty{
//   		SlurmCustomSettings: []interface{}{
//   			&SlurmCustomSettingProperty{
//   				ParameterName: jsii.String("parameterName"),
//   				ParameterValue: jsii.String("parameterValue"),
//   			},
//   		},
//   	},
//   	SpotOptions: &SpotOptionsProperty{
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html
//
type CfnComputeNodeGroupProps struct {
	// The ID of the cluster of the compute node group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-clusterid
	//
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// An Amazon EC2 launch template AWS PCS uses to launch compute nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-customlaunchtemplate
	//
	CustomLaunchTemplate interface{} `field:"required" json:"customLaunchTemplate" yaml:"customLaunchTemplate"`
	// The Amazon Resource Name (ARN) of the IAM instance profile used to pass an IAM role when launching EC2 instances.
	//
	// The role contained in your instance profile must have pcs:RegisterComputeNodeGroupInstance permissions attached to provision instances correctly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-iaminstanceprofilearn
	//
	IamInstanceProfileArn *string `field:"required" json:"iamInstanceProfileArn" yaml:"iamInstanceProfileArn"`
	// A list of EC2 instance configurations that AWS PCS can provision in the compute node group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-instanceconfigs
	//
	InstanceConfigs interface{} `field:"required" json:"instanceConfigs" yaml:"instanceConfigs"`
	// Specifies the boundaries of the compute node group auto scaling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-scalingconfiguration
	//
	ScalingConfiguration interface{} `field:"required" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// The list of subnet IDs where instances are provisioned by the compute node group.
	//
	// The subnets must be in the same VPC as the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the Amazon Machine Image (AMI) that AWS PCS uses to launch instances.
	//
	// If not provided, AWS PCS uses the AMI ID specified in the custom launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-amiid
	//
	AmiId *string `field:"optional" json:"amiId" yaml:"amiId"`
	// The name that identifies the compute node group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies how EC2 instances are purchased on your behalf.
	//
	// AWS PCS supports On-Demand and Spot instances. For more information, see Instance purchasing options in the Amazon Elastic Compute Cloud User Guide. If you don't provide this option, it defaults to On-Demand.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-purchaseoption
	//
	PurchaseOption *string `field:"optional" json:"purchaseOption" yaml:"purchaseOption"`
	// Additional options related to the Slurm scheduler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-slurmconfiguration
	//
	SlurmConfiguration interface{} `field:"optional" json:"slurmConfiguration" yaml:"slurmConfiguration"`
	// Additional configuration when you specify `SPOT` as the `purchaseOption` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-spotoptions
	//
	SpotOptions interface{} `field:"optional" json:"spotOptions" yaml:"spotOptions"`
	// 1 or more tags added to the resource.
	//
	// Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

