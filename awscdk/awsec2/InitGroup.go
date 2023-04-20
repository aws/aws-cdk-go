package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Create Linux/UNIX groups and assign group IDs.
//
// Not supported for Windows systems.
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
type InitGroup interface {
	InitElement
	// Returns the init element type for this element.
	ElementType() *string
}

// The jsii proxy struct for InitGroup
type jsiiProxy_InitGroup struct {
	jsiiProxy_InitElement
}

func (j *jsiiProxy_InitGroup) ElementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementType",
		&returns,
	)
	return returns
}


func NewInitGroup(groupName *string, groupId *float64) InitGroup {
	_init_.Initialize()

	if err := validateNewInitGroupParameters(groupName); err != nil {
		panic(err)
	}
	j := jsiiProxy_InitGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitGroup",
		[]interface{}{groupName, groupId},
		&j,
	)

	return &j
}

func NewInitGroup_Override(i InitGroup, groupName *string, groupId *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitGroup",
		[]interface{}{groupName, groupId},
		i,
	)
}

// Create a group from its name, and optionally, group id.
func InitGroup_FromName(groupName *string, groupId *float64) InitGroup {
	_init_.Initialize()

	if err := validateInitGroup_FromNameParameters(groupName); err != nil {
		panic(err)
	}
	var returns InitGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitGroup",
		"fromName",
		[]interface{}{groupName, groupId},
		&returns,
	)

	return returns
}

