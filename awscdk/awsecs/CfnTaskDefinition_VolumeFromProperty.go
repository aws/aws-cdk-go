package awsecs


// Details on a data volume from another container in the same task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeFromProperty := &VolumeFromProperty{
//   	ReadOnly: jsii.Boolean(false),
//   	SourceContainer: jsii.String("sourceContainer"),
//   }
//
type CfnTaskDefinition_VolumeFromProperty struct {
	// If this value is `true` , the container has read-only access to the volume.
	//
	// If this value is `false` , then the container can write to the volume. The default value is `false` .
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The name of another container within the same task definition to mount volumes from.
	SourceContainer *string `field:"optional" json:"sourceContainer" yaml:"sourceContainer"`
}

