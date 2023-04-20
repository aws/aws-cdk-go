package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A package to be installed during cfn-init time.
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
type InitPackage interface {
	InitElement
	// Returns the init element type for this element.
	ElementType() *string
	RenderPackageVersions() interface{}
}

// The jsii proxy struct for InitPackage
type jsiiProxy_InitPackage struct {
	jsiiProxy_InitElement
}

func (j *jsiiProxy_InitPackage) ElementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementType",
		&returns,
	)
	return returns
}


func NewInitPackage(type_ *string, versions *[]*string, packageName *string, serviceHandles *[]InitServiceRestartHandle) InitPackage {
	_init_.Initialize()

	if err := validateNewInitPackageParameters(type_, versions); err != nil {
		panic(err)
	}
	j := jsiiProxy_InitPackage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitPackage",
		[]interface{}{type_, versions, packageName, serviceHandles},
		&j,
	)

	return &j
}

func NewInitPackage_Override(i InitPackage, type_ *string, versions *[]*string, packageName *string, serviceHandles *[]InitServiceRestartHandle) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitPackage",
		[]interface{}{type_, versions, packageName, serviceHandles},
		i,
	)
}

// Install a package using APT.
func InitPackage_Apt(packageName *string, options *NamedPackageOptions) InitPackage {
	_init_.Initialize()

	if err := validateInitPackage_AptParameters(packageName, options); err != nil {
		panic(err)
	}
	var returns InitPackage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitPackage",
		"apt",
		[]interface{}{packageName, options},
		&returns,
	)

	return returns
}

// Install an MSI package from an HTTP URL or a location on disk.
func InitPackage_Msi(location *string, options *LocationPackageOptions) InitPackage {
	_init_.Initialize()

	if err := validateInitPackage_MsiParameters(location, options); err != nil {
		panic(err)
	}
	var returns InitPackage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitPackage",
		"msi",
		[]interface{}{location, options},
		&returns,
	)

	return returns
}

// Install a package from PyPI.
func InitPackage_Python(packageName *string, options *NamedPackageOptions) InitPackage {
	_init_.Initialize()

	if err := validateInitPackage_PythonParameters(packageName, options); err != nil {
		panic(err)
	}
	var returns InitPackage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitPackage",
		"python",
		[]interface{}{packageName, options},
		&returns,
	)

	return returns
}

// Install an RPM from an HTTP URL or a location on disk.
func InitPackage_Rpm(location *string, options *LocationPackageOptions) InitPackage {
	_init_.Initialize()

	if err := validateInitPackage_RpmParameters(location, options); err != nil {
		panic(err)
	}
	var returns InitPackage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitPackage",
		"rpm",
		[]interface{}{location, options},
		&returns,
	)

	return returns
}

// Install a package from RubyGems.
func InitPackage_RubyGem(gemName *string, options *NamedPackageOptions) InitPackage {
	_init_.Initialize()

	if err := validateInitPackage_RubyGemParameters(gemName, options); err != nil {
		panic(err)
	}
	var returns InitPackage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitPackage",
		"rubyGem",
		[]interface{}{gemName, options},
		&returns,
	)

	return returns
}

// Install a package using Yum.
func InitPackage_Yum(packageName *string, options *NamedPackageOptions) InitPackage {
	_init_.Initialize()

	if err := validateInitPackage_YumParameters(packageName, options); err != nil {
		panic(err)
	}
	var returns InitPackage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitPackage",
		"yum",
		[]interface{}{packageName, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InitPackage) RenderPackageVersions() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"renderPackageVersions",
		nil, // no parameters
		&returns,
	)

	return returns
}

