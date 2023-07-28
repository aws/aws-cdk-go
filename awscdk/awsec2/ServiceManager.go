package awsec2


// The service manager that will be used by InitServices.
//
// The value needs to match the service manager used by your operating
// system.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//
//
//   ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2023(),
//
//   	Init: ec2.CloudFormationInit_FromElements(ec2.InitService_SystemdConfigFile(jsii.String("simpleserver"), &SystemdConfigFileOptions{
//   		Command: jsii.String("/usr/bin/python3 -m http.server 8080"),
//   		Cwd: jsii.String("/var/www/html"),
//   	}), ec2.InitService_Enable(jsii.String("simpleserver"), &InitServiceOptions{
//   		ServiceManager: ec2.ServiceManager_SYSTEMD,
//   	}), ec2.InitFile_FromString(jsii.String("/var/www/html/index.html"), jsii.String("Hello! It's working!"))),
//   })
//
type ServiceManager string

const (
	// Use SysVinit.
	//
	// This is the default for Linux systems.
	ServiceManager_SYSVINIT ServiceManager = "SYSVINIT"
	// Use Windows.
	//
	// This is the default for Windows systems.
	ServiceManager_WINDOWS ServiceManager = "WINDOWS"
	// Use systemd.
	ServiceManager_SYSTEMD ServiceManager = "SYSTEMD"
)

