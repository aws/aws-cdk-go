package awsiotsitewise


// Attaches a data source to the container filesystem for a task at a customer-supplied relative path under the service-owned mount root.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mountProperty := &MountProperty{
//   	Name: jsii.String("name"),
//   	RelativePath: jsii.String("relativePath"),
//   	Source: &MountSourceProperty{
//   		S3AccessPoint: &S3AccessPointSourceProperty{
//   			AccessPointArn: jsii.String("accessPointArn"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	StorageType: jsii.String("storageType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-mount.html
//
type CfnTaskPropsMixin_MountProperty struct {
	// A unique name for the mount within the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-mount.html#cfn-iotsitewise-task-mount-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The relative path under the service-owned mount root where this mount is attached inside the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-mount.html#cfn-iotsitewise-task-mount-relativepath
	//
	RelativePath *string `field:"optional" json:"relativePath" yaml:"relativePath"`
	// The data source configuration for a mount.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-mount.html#cfn-iotsitewise-task-mount-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The type of storage used for the mount inside the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-mount.html#cfn-iotsitewise-task-mount-storagetype
	//
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
}

