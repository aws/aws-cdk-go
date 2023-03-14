package awsbatch


// Specifies an Amazon EKS volume for a job definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksVolumeProperty := &EksVolumeProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	EmptyDir: &EmptyDirProperty{
//   		Medium: jsii.String("medium"),
//   		SizeLimit: jsii.String("sizeLimit"),
//   	},
//   	HostPath: &HostPathProperty{
//   		Path: jsii.String("path"),
//   	},
//   	Secret: &SecretProperty{
//   		Name: jsii.String("name"),
//   		ValueFrom: jsii.String("valueFrom"),
//   	},
//   }
//
type CfnJobDefinition_EksVolumeProperty struct {
	// The name of the volume.
	//
	// The name must be allowed as a DNS subdomain name. For more information, see [DNS subdomain names](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names) in the *Kubernetes documentation* .
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the configuration of a Kubernetes `emptyDir` volume.
	//
	// For more information, see [emptyDir](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) in the *Kubernetes documentation* .
	EmptyDir interface{} `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// Specifies the configuration of a Kubernetes `hostPath` volume.
	//
	// For more information, see [hostPath](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/volumes/#hostpath) in the *Kubernetes documentation* .
	HostPath interface{} `field:"optional" json:"hostPath" yaml:"hostPath"`
	// Specifies the configuration of a Kubernetes `secret` volume.
	//
	// For more information, see [secret](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/volumes/#secret) in the *Kubernetes documentation* .
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
}

