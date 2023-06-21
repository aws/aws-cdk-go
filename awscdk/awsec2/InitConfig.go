package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A collection of configuration elements.
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
type InitConfig interface {
	// Add one or more elements to the config.
	Add(elements ...InitElement)
	// Whether this configset has elements or not.
	IsEmpty() *bool
}

// The jsii proxy struct for InitConfig
type jsiiProxy_InitConfig struct {
	_ byte // padding
}

func NewInitConfig(elements *[]InitElement) InitConfig {
	_init_.Initialize()

	if err := validateNewInitConfigParameters(elements); err != nil {
		panic(err)
	}
	j := jsiiProxy_InitConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitConfig",
		[]interface{}{elements},
		&j,
	)

	return &j
}

func NewInitConfig_Override(i InitConfig, elements *[]InitElement) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitConfig",
		[]interface{}{elements},
		i,
	)
}

func (i *jsiiProxy_InitConfig) Add(elements ...InitElement) {
	args := []interface{}{}
	for _, a := range elements {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"add",
		args,
	)
}

func (i *jsiiProxy_InitConfig) IsEmpty() *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"isEmpty",
		nil, // no parameters
		&returns,
	)

	return returns
}

