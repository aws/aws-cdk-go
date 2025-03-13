package awsgreengrassv2


// Contains information about a volume that Linux processes in a container can access.
//
// When you define a volume, the AWS IoT Greengrass Core software mounts the source files to the destination inside the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaVolumeMountProperty := &LambdaVolumeMountProperty{
//   	AddGroupOwner: jsii.Boolean(false),
//   	DestinationPath: jsii.String("destinationPath"),
//   	Permission: jsii.String("permission"),
//   	SourcePath: jsii.String("sourcePath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdavolumemount.html
//
type CfnComponentVersion_LambdaVolumeMountProperty struct {
	// Whether or not to add the AWS IoT Greengrass user group as an owner of the volume.
	//
	// Default: `false`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdavolumemount.html#cfn-greengrassv2-componentversion-lambdavolumemount-addgroupowner
	//
	AddGroupOwner interface{} `field:"optional" json:"addGroupOwner" yaml:"addGroupOwner"`
	// The path to the logical volume in the file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdavolumemount.html#cfn-greengrassv2-componentversion-lambdavolumemount-destinationpath
	//
	DestinationPath *string `field:"optional" json:"destinationPath" yaml:"destinationPath"`
	// The permission to access the volume: read/only ( `ro` ) or read/write ( `rw` ).
	//
	// Default: `ro`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdavolumemount.html#cfn-greengrassv2-componentversion-lambdavolumemount-permission
	//
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
	// The path to the physical volume in the file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdavolumemount.html#cfn-greengrassv2-componentversion-lambdavolumemount-sourcepath
	//
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

