package awsgamelift


// Defines the mount point configuration within a container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerMountPointProperty := &ContainerMountPointProperty{
//   	InstancePath: jsii.String("instancePath"),
//
//   	// the properties below are optional
//   	AccessLevel: jsii.String("accessLevel"),
//   	ContainerPath: jsii.String("containerPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containermountpoint.html
//
type CfnContainerGroupDefinition_ContainerMountPointProperty struct {
	// The path on the host that will be mounted in the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containermountpoint.html#cfn-gamelift-containergroupdefinition-containermountpoint-instancepath
	//
	InstancePath *string `field:"required" json:"instancePath" yaml:"instancePath"`
	// The access permissions for the mounted path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containermountpoint.html#cfn-gamelift-containergroupdefinition-containermountpoint-accesslevel
	//
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// The path inside the container where the mount is accessible.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containermountpoint.html#cfn-gamelift-containergroupdefinition-containermountpoint-containerpath
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
}

