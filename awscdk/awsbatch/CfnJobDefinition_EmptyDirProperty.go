package awsbatch


// Specifies the configuration of a Kubernetes `emptyDir` volume.
//
// An `emptyDir` volume is first created when a pod is assigned to a node. It exists as long as that pod is running on that node. The `emptyDir` volume is initially empty. All containers in the pod can read and write the files in the `emptyDir` volume. However, the `emptyDir` volume can be mounted at the same or different paths in each container. When a pod is removed from a node for any reason, the data in the `emptyDir` is deleted permanently. For more information, see [emptyDir](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) in the *Kubernetes documentation* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emptyDirProperty := &EmptyDirProperty{
//   	Medium: jsii.String("medium"),
//   	SizeLimit: jsii.String("sizeLimit"),
//   }
//
type CfnJobDefinition_EmptyDirProperty struct {
	// The medium to store the volume.
	//
	// The default value is an empty string, which uses the storage of the node.
	//
	// - **""** - *(Default)* Use the disk storage of the node.
	// - **"Memory"** - Use the `tmpfs` volume that's backed by the RAM of the node. Contents of the volume are lost when the node reboots, and any storage on the volume counts against the container's memory limit.
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	// The maximum size of the volume.
	//
	// By default, there's no maximum size defined.
	SizeLimit *string `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

