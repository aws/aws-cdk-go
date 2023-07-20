package awsbatch


// The volume mounts for a container for an Amazon EKS job.
//
// For more information about volumes and volume mounts in Kubernetes, see [Volumes](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/volumes/) in the *Kubernetes documentation* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksContainerVolumeMountProperty := &EksContainerVolumeMountProperty{
//   	MountPath: jsii.String("mountPath"),
//   	Name: jsii.String("name"),
//   	ReadOnly: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainervolumemount.html
//
type CfnJobDefinition_EksContainerVolumeMountProperty struct {
	// The path on the container where the volume is mounted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainervolumemount.html#cfn-batch-jobdefinition-ekscontainervolumemount-mountpath
	//
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
	// The name the volume mount.
	//
	// This must match the name of one of the volumes in the pod.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainervolumemount.html#cfn-batch-jobdefinition-ekscontainervolumemount-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// If this value is `true` , the container has read-only access to the volume.
	//
	// Otherwise, the container can write to the volume. The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainervolumemount.html#cfn-batch-jobdefinition-ekscontainervolumemount-readonly
	//
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

