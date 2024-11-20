package awsgamelift


// Specifies the information required to run game servers with this container group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gameServerContainerDefinitionProperty := &GameServerContainerDefinitionProperty{
//   	ContainerName: jsii.String("containerName"),
//   	ImageUri: jsii.String("imageUri"),
//   	ServerSdkVersion: jsii.String("serverSdkVersion"),
//
//   	// the properties below are optional
//   	DependsOn: []interface{}{
//   		&ContainerDependencyProperty{
//   			Condition: jsii.String("condition"),
//   			ContainerName: jsii.String("containerName"),
//   		},
//   	},
//   	EnvironmentOverride: []interface{}{
//   		&ContainerEnvironmentProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MountPoints: []interface{}{
//   		&ContainerMountPointProperty{
//   			InstancePath: jsii.String("instancePath"),
//
//   			// the properties below are optional
//   			AccessLevel: jsii.String("accessLevel"),
//   			ContainerPath: jsii.String("containerPath"),
//   		},
//   	},
//   	PortConfiguration: &PortConfigurationProperty{
//   		ContainerPortRanges: []interface{}{
//   			&ContainerPortRangeProperty{
//   				FromPort: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
//   				ToPort: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ResolvedImageDigest: jsii.String("resolvedImageDigest"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html
//
type CfnContainerGroupDefinition_GameServerContainerDefinitionProperty struct {
	// A descriptive label for the container definition.
	//
	// Container definition names must be unique with a container group definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-containername
	//
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Specifies the image URI of this container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-imageuri
	//
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The version of the server SDK used in this container group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-serversdkversion
	//
	ServerSdkVersion *string `field:"required" json:"serverSdkVersion" yaml:"serverSdkVersion"`
	// A list of container dependencies that determines when this container starts up and shuts down.
	//
	// For container groups with multiple containers, dependencies let you define a startup/shutdown sequence across the containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-dependson
	//
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// The environment variables to pass to a container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-environmentoverride
	//
	EnvironmentOverride interface{} `field:"optional" json:"environmentOverride" yaml:"environmentOverride"`
	// A list of mount point configurations to be used in a container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-mountpoints
	//
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// Defines the ports on a container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-portconfiguration
	//
	PortConfiguration interface{} `field:"optional" json:"portConfiguration" yaml:"portConfiguration"`
	// The digest of the container image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-resolvedimagedigest
	//
	ResolvedImageDigest *string `field:"optional" json:"resolvedImageDigest" yaml:"resolvedImageDigest"`
}

