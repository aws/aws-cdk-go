package awsec2


// Options for CloudFormationInit.withConfigSets.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//   var instanceType instanceType
//   var machineImage iMachineImage
//
//
//   ec2.NewInstance(this, jsii.String("Instance"), &instanceProps{
//   	vpc: vpc,
//   	instanceType: instanceType,
//   	machineImage: machineImage,
//
//   	// Showing the most complex setup, if you have simpler requirements
//   	// you can use `CloudFormationInit.fromElements()`.
//   	init: ec2.cloudFormationInit.fromConfigSets(&configSetProps{
//   		configSets: map[string][]*string{
//   			// Applies the configs below in this order
//   			"default": []*string{
//   				jsii.String("yumPreinstall"),
//   				jsii.String("config"),
//   			},
//   		},
//   		configs: map[string]initConfig{
//   			"yumPreinstall": ec2.NewInitConfig([]InitElement{
//   				ec2.InitPackage.yum(jsii.String("git")),
//   			}),
//   			"config": ec2.NewInitConfig([]InitElement{
//   				ec2.InitFile.fromObject(jsii.String("/etc/stack.json"), map[string]interface{}{
//   					"stackId": awscdk.*stack.of(this).stackId,
//   					"stackName": awscdk.*stack.of(this).stackName,
//   					"region": awscdk.*stack.of(this).region,
//   				}),
//   				ec2.InitGroup.fromName(jsii.String("my-group")),
//   				ec2.InitUser.fromName(jsii.String("my-user")),
//   				ec2.InitPackage.rpm(jsii.String("http://mirrors.ukfast.co.uk/sites/dl.fedoraproject.org/pub/epel/8/Everything/x86_64/Packages/r/rubygem-git-1.5.0-2.el8.noarch.rpm")),
//   			}),
//   		},
//   	}),
//   	initOptions: &applyCloudFormationInitOptions{
//   		// Optional, which configsets to activate (['default'] by default)
//   		configSets: []*string{
//   			jsii.String("default"),
//   		},
//
//   		// Optional, how long the installation is expected to take (5 minutes by default)
//   		timeout: awscdk.Duration.minutes(jsii.Number(30)),
//
//   		// Optional, whether to include the --url argument when running cfn-init and cfn-signal commands (false by default)
//   		includeUrl: jsii.Boolean(true),
//
//   		// Optional, whether to include the --role argument when running cfn-init and cfn-signal commands (false by default)
//   		includeRole: jsii.Boolean(true),
//   	},
//   })
//
type ConfigSetProps struct {
	// The sets of configs to pick from.
	Configs *map[string]InitConfig `field:"required" json:"configs" yaml:"configs"`
	// The definitions of each config set.
	ConfigSets *map[string]*[]*string `field:"required" json:"configSets" yaml:"configSets"`
}

