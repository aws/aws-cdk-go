package awsecs


// This parameter is specified when you're using [Amazon FSx for Windows File Server](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/what-is.html) file system for task storage.
//
// For more information and the input format, see [Amazon FSx for Windows File Server volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/wfsx-volumes.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fSxWindowsFileServerVolumeConfigurationProperty := &FSxWindowsFileServerVolumeConfigurationProperty{
//   	FileSystemId: jsii.String("fileSystemId"),
//   	RootDirectory: jsii.String("rootDirectory"),
//
//   	// the properties below are optional
//   	AuthorizationConfig: &FSxAuthorizationConfigProperty{
//   		CredentialsParameter: jsii.String("credentialsParameter"),
//   		Domain: jsii.String("domain"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-fsxwindowsfileservervolumeconfiguration.html
//
type CfnTaskDefinition_FSxWindowsFileServerVolumeConfigurationProperty struct {
	// The Amazon FSx for Windows File Server file system ID to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-fsxwindowsfileservervolumeconfiguration.html#cfn-ecs-taskdefinition-fsxwindowsfileservervolumeconfiguration-filesystemid
	//
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// The directory within the Amazon FSx for Windows File Server file system to mount as the root directory inside the host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-fsxwindowsfileservervolumeconfiguration.html#cfn-ecs-taskdefinition-fsxwindowsfileservervolumeconfiguration-rootdirectory
	//
	RootDirectory *string `field:"required" json:"rootDirectory" yaml:"rootDirectory"`
	// The authorization configuration details for the Amazon FSx for Windows File Server file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-fsxwindowsfileservervolumeconfiguration.html#cfn-ecs-taskdefinition-fsxwindowsfileservervolumeconfiguration-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

