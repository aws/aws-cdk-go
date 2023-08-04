package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for applying CloudFormation init to an instance or instance group.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//   var machineImage iMachineImage
//
//
//   ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//
//   	// Showing the most complex setup, if you have simpler requirements
//   	// you can use `CloudFormationInit.fromElements()`.
//   	Init: ec2.CloudFormationInit_FromConfigSets(&ConfigSetProps{
//   		ConfigSets: map[string][]*string{
//   			// Applies the configs below in this order
//   			"default": []*string{
//   				jsii.String("yumPreinstall"),
//   				jsii.String("config"),
//   			},
//   		},
//   		Configs: map[string]initConfig{
//   			"yumPreinstall": ec2.NewInitConfig([]InitElement{
//   				ec2.InitPackage_yum(jsii.String("git")),
//   			}),
//   			"config": ec2.NewInitConfig([]InitElement{
//   				ec2.InitFile_fromObject(jsii.String("/etc/stack.json"), map[string]interface{}{
//   					"stackId": awscdk.*stack_of(this).stackId,
//   					"stackName": awscdk.*stack_of(this).stackName,
//   					"region": awscdk.*stack_of(this).region,
//   				}),
//   				ec2.InitGroup_fromName(jsii.String("my-group")),
//   				ec2.InitUser_fromName(jsii.String("my-user")),
//   				ec2.InitPackage_rpm(jsii.String("http://mirrors.ukfast.co.uk/sites/dl.fedoraproject.org/pub/epel/8/Everything/x86_64/Packages/r/rubygem-git-1.5.0-2.el8.noarch.rpm")),
//   			}),
//   		},
//   	}),
//   	InitOptions: &ApplyCloudFormationInitOptions{
//   		// Optional, which configsets to activate (['default'] by default)
//   		ConfigSets: []*string{
//   			jsii.String("default"),
//   		},
//
//   		// Optional, how long the installation is expected to take (5 minutes by default)
//   		Timeout: awscdk.Duration_Minutes(jsii.Number(30)),
//
//   		// Optional, whether to include the --url argument when running cfn-init and cfn-signal commands (false by default)
//   		IncludeUrl: jsii.Boolean(true),
//
//   		// Optional, whether to include the --role argument when running cfn-init and cfn-signal commands (false by default)
//   		IncludeRole: jsii.Boolean(true),
//   	},
//   })
//
type ApplyCloudFormationInitOptions struct {
	// ConfigSet to activate.
	// Default: ['default'].
	//
	ConfigSets *[]*string `field:"optional" json:"configSets" yaml:"configSets"`
	// Force instance replacement by embedding a config fingerprint.
	//
	// If `true` (the default), a hash of the config will be embedded into the
	// UserData, so that if the config changes, the UserData changes.
	//
	// - If the EC2 instance is instance-store backed or
	//   `userDataCausesReplacement` is set, this will cause the instance to be
	//   replaced and the new configuration to be applied.
	// - If the instance is EBS-backed and `userDataCausesReplacement` is not
	//   set, the change of UserData will make the instance restart but not be
	//   replaced, and the configuration will not be applied automatically.
	//
	// If `false`, no hash will be embedded, and if the CloudFormation Init
	// config changes nothing will happen to the running instance. If a
	// config update introduces errors, you will not notice until after the
	// CloudFormation deployment successfully finishes and the next instance
	// fails to launch.
	// Default: true.
	//
	EmbedFingerprint *bool `field:"optional" json:"embedFingerprint" yaml:"embedFingerprint"`
	// Don't fail the instance creation when cfn-init fails.
	//
	// You can use this to prevent CloudFormation from rolling back when
	// instances fail to start up, to help in debugging.
	// Default: false.
	//
	IgnoreFailures *bool `field:"optional" json:"ignoreFailures" yaml:"ignoreFailures"`
	// Include --role argument when running cfn-init and cfn-signal commands.
	//
	// This will be the IAM instance profile attached to the EC2 instance.
	// Default: false.
	//
	IncludeRole *bool `field:"optional" json:"includeRole" yaml:"includeRole"`
	// Include --url argument when running cfn-init and cfn-signal commands.
	//
	// This will be the cloudformation endpoint in the deployed region
	// e.g. https://cloudformation.us-east-1.amazonaws.com
	// Default: false.
	//
	IncludeUrl *bool `field:"optional" json:"includeUrl" yaml:"includeUrl"`
	// Print the results of running cfn-init to the Instance System Log.
	//
	// By default, the output of running cfn-init is written to a log file
	// on the instance. Set this to `true` to print it to the System Log
	// (visible from the EC2 Console), `false` to not print it.
	//
	// (Be aware that the system log is refreshed at certain points in
	// time of the instance life cycle, and successful execution may
	// not always show up).
	// Default: true.
	//
	PrintLog *bool `field:"optional" json:"printLog" yaml:"printLog"`
	// Timeout waiting for the configuration to be applied.
	// Default: Duration.minutes(5)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

