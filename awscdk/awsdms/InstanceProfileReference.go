package awsdms


// A reference to a InstanceProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceProfileReference := &InstanceProfileReference{
//   	InstanceProfileArn: jsii.String("instanceProfileArn"),
//   }
//
type InstanceProfileReference struct {
	// The InstanceProfileArn of the InstanceProfile resource.
	InstanceProfileArn *string `field:"required" json:"instanceProfileArn" yaml:"instanceProfileArn"`
}

