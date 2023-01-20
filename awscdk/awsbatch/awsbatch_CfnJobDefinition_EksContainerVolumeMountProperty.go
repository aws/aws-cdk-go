package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksContainerVolumeMountProperty := &eksContainerVolumeMountProperty{
//   	mountPath: jsii.String("mountPath"),
//   	name: jsii.String("name"),
//   	readOnly: jsii.Boolean(false),
//   }
//
type CfnJobDefinition_EksContainerVolumeMountProperty struct {
	// `CfnJobDefinition.EksContainerVolumeMountProperty.MountPath`.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
	// `CfnJobDefinition.EksContainerVolumeMountProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnJobDefinition.EksContainerVolumeMountProperty.ReadOnly`.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

