package previewawsopsworksmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLayerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var customJson interface{}
//
//   cfnLayerMixinProps := &CfnLayerMixinProps{
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	AutoAssignElasticIps: jsii.Boolean(false),
//   	AutoAssignPublicIps: jsii.Boolean(false),
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
//   	EnableAutoHealing: jsii.Boolean(false),
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
//   	Name: jsii.String("name"),
//   	Packages: []*string{
//   		jsii.String("packages"),
//   	},
//   	Shortname: jsii.String("shortname"),
//   	StackId: jsii.String("stackId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html
//
type CfnLayerMixinProps struct {
	// One or more user-defined key-value pairs to be added to the stack attributes.
	//
	// To create a cluster layer, set the `EcsClusterArn` attribute to the cluster's ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// Whether to automatically assign an [Elastic IP address](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html) to the layer's instances. For more information, see [How to Edit a Layer](https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-basics-edit.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-autoassignelasticips
	//
	AutoAssignElasticIps interface{} `field:"optional" json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// For stacks that are running in a VPC, whether to automatically assign a public IP address to the layer's instances.
	//
	// For more information, see [How to Edit a Layer](https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-basics-edit.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-autoassignpublicips
	//
	AutoAssignPublicIps interface{} `field:"optional" json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// The ARN of an IAM profile to be used for the layer's EC2 instances.
	//
	// For more information about IAM ARNs, see [Using Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-custominstanceprofilearn
	//
	CustomInstanceProfileArn *string `field:"optional" json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// A JSON-formatted string containing custom stack configuration and deployment attributes to be installed on the layer's instances.
	//
	// For more information, see [Using Custom JSON](https://docs.aws.amazon.com/opsworks/latest/userguide/workingcookbook-json-override.html) . This feature is supported as of version 1.7.42 of the AWS CLI .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-customjson
	//
	CustomJson interface{} `field:"optional" json:"customJson" yaml:"customJson"`
	// A `LayerCustomRecipes` object that specifies the layer custom recipes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-customrecipes
	//
	CustomRecipes interface{} `field:"optional" json:"customRecipes" yaml:"customRecipes"`
	// An array containing the layer custom security group IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-customsecuritygroupids
	//
	CustomSecurityGroupIds *[]*string `field:"optional" json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Whether to disable auto healing for the layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-enableautohealing
	//
	EnableAutoHealing interface{} `field:"optional" json:"enableAutoHealing" yaml:"enableAutoHealing"`
	// Whether to install operating system and package updates when the instance boots.
	//
	// The default value is `true` . To control when updates are installed, set this value to `false` . You must then update your instances manually by using `CreateDeployment` to run the `update_dependencies` stack command or by manually running `yum` (Amazon Linux) or `apt-get` (Ubuntu) on the instances.
	//
	// > To ensure that your instances have the latest security updates, we strongly recommend using the default value of `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-installupdatesonboot
	//
	InstallUpdatesOnBoot interface{} `field:"optional" json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// A `LifeCycleEventConfiguration` object that you can use to configure the Shutdown event to specify an execution timeout and enable or disable Elastic Load Balancer connection draining.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-lifecycleeventconfiguration
	//
	LifecycleEventConfiguration interface{} `field:"optional" json:"lifecycleEventConfiguration" yaml:"lifecycleEventConfiguration"`
	// The load-based scaling configuration for the OpsWorks layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-loadbasedautoscaling
	//
	LoadBasedAutoScaling interface{} `field:"optional" json:"loadBasedAutoScaling" yaml:"loadBasedAutoScaling"`
	// The layer name, which is used by the console.
	//
	// Layer names can be a maximum of 32 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of `Package` objects that describes the layer packages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-packages
	//
	Packages *[]*string `field:"optional" json:"packages" yaml:"packages"`
	// For custom layers only, use this parameter to specify the layer's short name, which is used internally by OpsWorks Stacks and by Chef recipes.
	//
	// The short name is also used as the name for the directory where your app files are installed. It can have a maximum of 32 characters, which are limited to the alphanumeric characters, '-', '_', and '.'.
	//
	// Built-in layer short names are defined by OpsWorks Stacks. For more information, see the [Layer Reference](https://docs.aws.amazon.com/opsworks/latest/userguide/layers.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-shortname
	//
	Shortname *string `field:"optional" json:"shortname" yaml:"shortname"`
	// The layer stack ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-stackid
	//
	StackId *string `field:"optional" json:"stackId" yaml:"stackId"`
	// Specifies one or more sets of tags (key–value pairs) to associate with this OpsWorks layer.
	//
	// Use tags to manage your resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The layer type.
	//
	// A stack cannot have more than one built-in layer of the same type. It can have any number of custom layers. Built-in layers are not available in Chef 12 stacks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Whether to use Amazon EBS-optimized instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-useebsoptimizedinstances
	//
	UseEbsOptimizedInstances interface{} `field:"optional" json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
	// A `VolumeConfigurations` object that describes the layer's Amazon EBS volumes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-volumeconfigurations
	//
	VolumeConfigurations interface{} `field:"optional" json:"volumeConfigurations" yaml:"volumeConfigurations"`
}

