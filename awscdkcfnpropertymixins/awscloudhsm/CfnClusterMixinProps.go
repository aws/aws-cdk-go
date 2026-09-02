package awscloudhsm


// Properties for CfnClusterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnClusterMixinProps := &CfnClusterMixinProps{
//   	BackupRetentionPolicy: &BackupRetentionPolicyProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.String("value"),
//   	},
//   	HsmType: jsii.String("hsmType"),
//   	Mode: jsii.String("mode"),
//   	NetworkType: jsii.String("networkType"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudhsm-cluster.html
//
type CfnClusterMixinProps struct {
	// A policy that defines how the service retains backups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudhsm-cluster.html#cfn-cloudhsm-cluster-backupretentionpolicy
	//
	BackupRetentionPolicy interface{} `field:"optional" json:"backupRetentionPolicy" yaml:"backupRetentionPolicy"`
	// The type of HSM to use in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudhsm-cluster.html#cfn-cloudhsm-cluster-hsmtype
	//
	HsmType *string `field:"optional" json:"hsmType" yaml:"hsmType"`
	// The mode to use in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudhsm-cluster.html#cfn-cloudhsm-cluster-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The NetworkType to create a cluster with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudhsm-cluster.html#cfn-cloudhsm-cluster-networktype
	//
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The identifiers (IDs) of the subnets where the cluster is created.
	//
	// You must specify at least one subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudhsm-cluster.html#cfn-cloudhsm-cluster-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Tags to apply to the CloudHSM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudhsm-cluster.html#cfn-cloudhsm-cluster-tags
	//
	Tags *[]*CfnClusterPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

