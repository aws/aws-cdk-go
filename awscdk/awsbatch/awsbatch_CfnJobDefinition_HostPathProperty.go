package awsbatch


// Specifies the configuration of a Kubernetes `hostPath` volume.
//
// A `hostPath` volume mounts an existing file or directory from the host node's filesystem into your pod. For more information, see [hostPath](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/volumes/#hostpath) in the *Kubernetes documentation* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostPathProperty := &hostPathProperty{
//   	path: jsii.String("path"),
//   }
//
type CfnJobDefinition_HostPathProperty struct {
	// The path of the file or directory on the host to mount into containers on the pod.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

