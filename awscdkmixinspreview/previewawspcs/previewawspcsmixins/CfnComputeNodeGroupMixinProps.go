package previewawspcsmixins


// Properties for CfnComputeNodeGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnComputeNodeGroupMixinProps := &CfnComputeNodeGroupMixinProps{
//   	AmiId: jsii.String("amiId"),
//   	ClusterId: jsii.String("clusterId"),
//   	CustomLaunchTemplate: &CustomLaunchTemplateProperty{
//   		TemplateId: jsii.String("templateId"),
//   		Version: jsii.String("version"),
//   	},
//   	IamInstanceProfileArn: jsii.String("iamInstanceProfileArn"),
//   	InstanceConfigs: []interface{}{
//   		&InstanceConfigProperty{
//   			InstanceType: jsii.String("instanceType"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	PurchaseOption: jsii.String("purchaseOption"),
//   	ScalingConfiguration: &ScalingConfigurationProperty{
//   		MaxInstanceCount: jsii.Number(123),
//   		MinInstanceCount: jsii.Number(123),
//   	},
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
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html
//
type CfnComputeNodeGroupMixinProps struct {
	// The ID of the Amazon Machine Image (AMI) that AWS PCS uses to launch instances.
	//
	// If not provided, AWS PCS uses the AMI ID specified in the custom launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-amiid
	//
	AmiId *string `field:"optional" json:"amiId" yaml:"amiId"`
	// The ID of the cluster of the compute node group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-clusterid
	//
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// An Amazon EC2 launch template AWS PCS uses to launch compute nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-customlaunchtemplate
	//
	CustomLaunchTemplate interface{} `field:"optional" json:"customLaunchTemplate" yaml:"customLaunchTemplate"`
	// The Amazon Resource Name (ARN) of the IAM instance profile used to pass an IAM role when launching EC2 instances.
	//
	// The role contained in your instance profile must have the `pcs:RegisterComputeNodeGroupInstance` permission and the role name must start with `AWSPCS` or must have the path `/aws-pcs/` . For more information, see [IAM instance profiles for AWS PCS](https://docs.aws.amazon.com//pcs/latest/userguide/security-instance-profiles.html) in the *AWS PCS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-iaminstanceprofilearn
	//
	IamInstanceProfileArn *string `field:"optional" json:"iamInstanceProfileArn" yaml:"iamInstanceProfileArn"`
	// A list of EC2 instance configurations that AWS PCS can provision in the compute node group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-instanceconfigs
	//
	InstanceConfigs interface{} `field:"optional" json:"instanceConfigs" yaml:"instanceConfigs"`
	// The name that identifies the compute node group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies how EC2 instances are purchased on your behalf.
	//
	// AWS PCS supports On-Demand Instances, Spot Instances, and Amazon EC2 Capacity Blocks for ML. For more information, see [Amazon EC2 billing and purchasing options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-purchasing-options.html) in the *Amazon Elastic Compute Cloud User Guide* . For more information about AWS PCS support for Capacity Blocks, see [Using Amazon EC2 Capacity Blocks for ML with AWS PCS](https://docs.aws.amazon.com/pcs/latest/userguide/capacity-blocks.html) in the *AWS PCS User Guide* . If you don't provide this option, it defaults to On-Demand.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-purchaseoption
	//
	PurchaseOption *string `field:"optional" json:"purchaseOption" yaml:"purchaseOption"`
	// Specifies the boundaries of the compute node group auto scaling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-scalingconfiguration
	//
	ScalingConfiguration interface{} `field:"optional" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// Additional options related to the Slurm scheduler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-slurmconfiguration
	//
	SlurmConfiguration interface{} `field:"optional" json:"slurmConfiguration" yaml:"slurmConfiguration"`
	// Additional configuration when you specify `SPOT` as the `purchaseOption` for the `CreateComputeNodeGroup` API action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-spotoptions
	//
	SpotOptions interface{} `field:"optional" json:"spotOptions" yaml:"spotOptions"`
	// The list of subnet IDs where instances are provisioned by the compute node group.
	//
	// The subnets must be in the same VPC as the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// 1 or more tags added to the resource.
	//
	// Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html#cfn-pcs-computenodegroup-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

