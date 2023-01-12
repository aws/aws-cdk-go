package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Create Linux/UNIX users and to assign user IDs.
//
// Users are created as non-interactive system users with a shell of
// /sbin/nologin. This is by design and cannot be modified.
//
// Not supported for Windows systems.
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
type InitUser interface {
	InitElement
	// Returns the init element type for this element.
	ElementType() *string
}

// The jsii proxy struct for InitUser
type jsiiProxy_InitUser struct {
	jsiiProxy_InitElement
}

func (j *jsiiProxy_InitUser) ElementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementType",
		&returns,
	)
	return returns
}


func NewInitUser(userName *string, userOptions *InitUserOptions) InitUser {
	_init_.Initialize()

	if err := validateNewInitUserParameters(userName, userOptions); err != nil {
		panic(err)
	}
	j := jsiiProxy_InitUser{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitUser",
		[]interface{}{userName, userOptions},
		&j,
	)

	return &j
}

func NewInitUser_Override(i InitUser, userName *string, userOptions *InitUserOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitUser",
		[]interface{}{userName, userOptions},
		i,
	)
}

// Create a user from user name.
func InitUser_FromName(userName *string, options *InitUserOptions) InitUser {
	_init_.Initialize()

	if err := validateInitUser_FromNameParameters(userName, options); err != nil {
		panic(err)
	}
	var returns InitUser

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitUser",
		"fromName",
		[]interface{}{userName, options},
		&returns,
	)

	return returns
}

