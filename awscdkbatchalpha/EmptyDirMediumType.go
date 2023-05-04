// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// What medium the volume will live in.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   jobDefn := batch.NewEksJobDefinition(this, jsii.String("eksf2"), &EksJobDefinitionProps{
//   	Container: batch.NewEksContainerDefinition(this, jsii.String("container"), &EksContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		Volumes: []eksVolume{
//   			batch.*eksVolume_EmptyDir(&EmptyDirVolumeOptions{
//   				Name: jsii.String("myEmptyDirVolume"),
//   				MountPath: jsii.String("/mount/path"),
//   				Medium: batch.EmptyDirMediumType_MEMORY,
//   				Readonly: jsii.Boolean(true),
//   				SizeLimit: cdk.Size_Mebibytes(jsii.Number(2048)),
//   			}),
//   		},
//   	}),
//   })
//
// Experimental.
type EmptyDirMediumType string

const (
	// Use the disk storage of the node.
	//
	// Items written here will survive node reboots.
	// Experimental.
	EmptyDirMediumType_DISK EmptyDirMediumType = "DISK"
	// Use the `tmpfs` volume that is backed by RAM of the node.
	//
	// Items written here will *not* survive node reboots.
	// Experimental.
	EmptyDirMediumType_MEMORY EmptyDirMediumType = "MEMORY"
)

