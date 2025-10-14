package awspcs


// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &CfnClusterProps{
//   	Networking: &NetworkingProperty{
//   		NetworkType: jsii.String("networkType"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	Scheduler: &SchedulerProperty{
//   		Type: jsii.String("type"),
//   		Version: jsii.String("version"),
//   	},
//   	Size: jsii.String("size"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	SlurmConfiguration: &SlurmConfigurationProperty{
//   		Accounting: &AccountingProperty{
//   			Mode: jsii.String("mode"),
//
//   			// the properties below are optional
//   			DefaultPurgeTimeInDays: jsii.Number(123),
//   		},
//   		AuthKey: &AuthKeyProperty{
//   			SecretArn: jsii.String("secretArn"),
//   			SecretVersion: jsii.String("secretVersion"),
//   		},
//   		ScaleDownIdleTimeInSeconds: jsii.Number(123),
//   		SlurmCustomSettings: []interface{}{
//   			&SlurmCustomSettingProperty{
//   				ParameterName: jsii.String("parameterName"),
//   				ParameterValue: jsii.String("parameterValue"),
//   			},
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-cluster.html
//
type CfnClusterProps struct {
	// The networking configuration for the cluster's control plane.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-cluster.html#cfn-pcs-cluster-networking
	//
	Networking interface{} `field:"required" json:"networking" yaml:"networking"`
	// The cluster management and job scheduling software associated with the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-cluster.html#cfn-pcs-cluster-scheduler
	//
	Scheduler interface{} `field:"required" json:"scheduler" yaml:"scheduler"`
	// The size of the cluster.
	//
	// - `SMALL` : 32 compute nodes and 256 jobs
	// - `MEDIUM` : 512 compute nodes and 8192 jobs
	// - `LARGE` : 2048 compute nodes and 16,384 jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-cluster.html#cfn-pcs-cluster-size
	//
	Size *string `field:"required" json:"size" yaml:"size"`
	// The name that identifies the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-cluster.html#cfn-pcs-cluster-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Additional options related to the Slurm scheduler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-cluster.html#cfn-pcs-cluster-slurmconfiguration
	//
	SlurmConfiguration interface{} `field:"optional" json:"slurmConfiguration" yaml:"slurmConfiguration"`
	// 1 or more tags added to the resource.
	//
	// Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-cluster.html#cfn-pcs-cluster-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

