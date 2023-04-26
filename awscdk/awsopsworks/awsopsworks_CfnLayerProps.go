package awsopsworks

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLayer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customJson interface{}
//
//   cfnLayerProps := &CfnLayerProps{
//   	AutoAssignElasticIps: jsii.Boolean(false),
//   	AutoAssignPublicIps: jsii.Boolean(false),
//   	EnableAutoHealing: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	Shortname: jsii.String("shortname"),
//   	StackId: jsii.String("stackId"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	CustomInstanceProfileArn: jsii.String("customInstanceProfileArn"),
//   	CustomJson: customJson,
//   	CustomRecipes: &RecipesProperty{
//   		Configure: []*string{
//   			jsii.String("configure"),
//   		},
//   		Deploy: []*string{
//   			jsii.String("deploy"),
//   		},
//   		Setup: []*string{
//   			jsii.String("setup"),
//   		},
//   		Shutdown: []*string{
//   			jsii.String("shutdown"),
//   		},
//   		Undeploy: []*string{
//   			jsii.String("undeploy"),
//   		},
//   	},
//   	CustomSecurityGroupIds: []*string{
//   		jsii.String("customSecurityGroupIds"),
//   	},
//   	InstallUpdatesOnBoot: jsii.Boolean(false),
//   	LifecycleEventConfiguration: &LifecycleEventConfigurationProperty{
//   		ShutdownEventConfiguration: &ShutdownEventConfigurationProperty{
//   			DelayUntilElbConnectionsDrained: jsii.Boolean(false),
//   			ExecutionTimeout: jsii.Number(123),
//   		},
//   	},
//   	LoadBasedAutoScaling: &LoadBasedAutoScalingProperty{
//   		DownScaling: &AutoScalingThresholdsProperty{
//   			CpuThreshold: jsii.Number(123),
//   			IgnoreMetricsTime: jsii.Number(123),
//   			InstanceCount: jsii.Number(123),
//   			LoadThreshold: jsii.Number(123),
//   			MemoryThreshold: jsii.Number(123),
//   			ThresholdsWaitTime: jsii.Number(123),
//   		},
//   		Enable: jsii.Boolean(false),
//   		UpScaling: &AutoScalingThresholdsProperty{
//   			CpuThreshold: jsii.Number(123),
//   			IgnoreMetricsTime: jsii.Number(123),
//   			InstanceCount: jsii.Number(123),
//   			LoadThreshold: jsii.Number(123),
//   			MemoryThreshold: jsii.Number(123),
//   			ThresholdsWaitTime: jsii.Number(123),
//   		},
//   	},
//   	Packages: []*string{
//   		jsii.String("packages"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseEbsOptimizedInstances: jsii.Boolean(false),
//   	VolumeConfigurations: []interface{}{
//   		&VolumeConfigurationProperty{
//   			Encrypted: jsii.Boolean(false),
//   			Iops: jsii.Number(123),
//   			MountPoint: jsii.String("mountPoint"),
//   			NumberOfDisks: jsii.Number(123),
//   			RaidLevel: jsii.Number(123),
//   			Size: jsii.Number(123),
//   			VolumeType: jsii.String("volumeType"),
//   		},
//   	},
//   }
//
type CfnLayerProps struct {
	// Whether to automatically assign an [Elastic IP address](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html) to the layer's instances. For more information, see [How to Edit a Layer](https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-basics-edit.html) .
	AutoAssignElasticIps interface{} `field:"required" json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// For stacks that are running in a VPC, whether to automatically assign a public IP address to the layer's instances.
	//
	// For more information, see [How to Edit a Layer](https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-basics-edit.html) .
	AutoAssignPublicIps interface{} `field:"required" json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Whether to disable auto healing for the layer.
	EnableAutoHealing interface{} `field:"required" json:"enableAutoHealing" yaml:"enableAutoHealing"`
	// The layer name, which is used by the console.
	//
	// Layer names can be a maximum of 32 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
	// For custom layers only, use this parameter to specify the layer's short name, which is used internally by AWS OpsWorks Stacks and by Chef recipes.
	//
	// The short name is also used as the name for the directory where your app files are installed. It can have a maximum of 32 characters, which are limited to the alphanumeric characters, '-', '_', and '.'.
	//
	// Built-in layer short names are defined by AWS OpsWorks Stacks. For more information, see the [Layer Reference](https://docs.aws.amazon.com/opsworks/latest/userguide/layers.html) .
	Shortname *string `field:"required" json:"shortname" yaml:"shortname"`
	// The layer stack ID.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
	// The layer type.
	//
	// A stack cannot have more than one built-in layer of the same type. It can have any number of custom layers. Built-in layers are not available in Chef 12 stacks.
	Type *string `field:"required" json:"type" yaml:"type"`
	// One or more user-defined key-value pairs to be added to the stack attributes.
	//
	// To create a cluster layer, set the `EcsClusterArn` attribute to the cluster's ARN.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The ARN of an IAM profile to be used for the layer's EC2 instances.
	//
	// For more information about IAM ARNs, see [Using Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) .
	CustomInstanceProfileArn *string `field:"optional" json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// A JSON-formatted string containing custom stack configuration and deployment attributes to be installed on the layer's instances.
	//
	// For more information, see [Using Custom JSON](https://docs.aws.amazon.com/opsworks/latest/userguide/workingcookbook-json-override.html) . This feature is supported as of version 1.7.42 of the AWS CLI .
	CustomJson interface{} `field:"optional" json:"customJson" yaml:"customJson"`
	// A `LayerCustomRecipes` object that specifies the layer custom recipes.
	CustomRecipes interface{} `field:"optional" json:"customRecipes" yaml:"customRecipes"`
	// An array containing the layer custom security group IDs.
	CustomSecurityGroupIds *[]*string `field:"optional" json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Whether to install operating system and package updates when the instance boots.
	//
	// The default value is `true` . To control when updates are installed, set this value to `false` . You must then update your instances manually by using [CreateDeployment](https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/CreateDeployment) to run the `update_dependencies` stack command or by manually running `yum` (Amazon Linux) or `apt-get` (Ubuntu) on the instances.
	//
	// > To ensure that your instances have the latest security updates, we strongly recommend using the default value of `true` .
	InstallUpdatesOnBoot interface{} `field:"optional" json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// A `LifeCycleEventConfiguration` object that you can use to configure the Shutdown event to specify an execution timeout and enable or disable Elastic Load Balancer connection draining.
	LifecycleEventConfiguration interface{} `field:"optional" json:"lifecycleEventConfiguration" yaml:"lifecycleEventConfiguration"`
	// The load-based scaling configuration for the AWS OpsWorks layer.
	LoadBasedAutoScaling interface{} `field:"optional" json:"loadBasedAutoScaling" yaml:"loadBasedAutoScaling"`
	// An array of `Package` objects that describes the layer packages.
	Packages *[]*string `field:"optional" json:"packages" yaml:"packages"`
	// Specifies one or more sets of tags (key–value pairs) to associate with this AWS OpsWorks layer.
	//
	// Use tags to manage your resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Whether to use Amazon EBS-optimized instances.
	UseEbsOptimizedInstances interface{} `field:"optional" json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
	// A `VolumeConfigurations` object that describes the layer's Amazon EBS volumes.
	VolumeConfigurations interface{} `field:"optional" json:"volumeConfigurations" yaml:"volumeConfigurations"`
}

