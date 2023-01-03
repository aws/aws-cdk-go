package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emptyDirProperty := &emptyDirProperty{
//   	medium: jsii.String("medium"),
//   	sizeLimit: jsii.String("sizeLimit"),
//   }
//
type CfnJobDefinition_EmptyDirProperty struct {
	// `CfnJobDefinition.EmptyDirProperty.Medium`.
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	// `CfnJobDefinition.EmptyDirProperty.SizeLimit`.
	SizeLimit *string `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

