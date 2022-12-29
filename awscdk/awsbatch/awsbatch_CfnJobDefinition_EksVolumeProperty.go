package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksVolumeProperty := &eksVolumeProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	emptyDir: &emptyDirProperty{
//   		medium: jsii.String("medium"),
//   		sizeLimit: jsii.String("sizeLimit"),
//   	},
//   	hostPath: &hostPathProperty{
//   		path: jsii.String("path"),
//   	},
//   	secret: &secretProperty{
//   		name: jsii.String("name"),
//   		valueFrom: jsii.String("valueFrom"),
//   	},
//   }
//
type CfnJobDefinition_EksVolumeProperty struct {
	// `CfnJobDefinition.EksVolumeProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnJobDefinition.EksVolumeProperty.EmptyDir`.
	EmptyDir interface{} `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// `CfnJobDefinition.EksVolumeProperty.HostPath`.
	HostPath interface{} `field:"optional" json:"hostPath" yaml:"hostPath"`
	// `CfnJobDefinition.EksVolumeProperty.Secret`.
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
}

