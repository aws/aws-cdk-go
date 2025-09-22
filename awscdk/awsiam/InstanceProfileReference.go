package awsiam


// A reference to a InstanceProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceProfileReference := &InstanceProfileReference{
//   	InstanceProfileArn: jsii.String("instanceProfileArn"),
//   	InstanceProfileName: jsii.String("instanceProfileName"),
//   }
//
type InstanceProfileReference struct {
	// The ARN of the InstanceProfile resource.
	InstanceProfileArn *string `field:"required" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// The InstanceProfileName of the InstanceProfile resource.
	InstanceProfileName *string `field:"required" json:"instanceProfileName" yaml:"instanceProfileName"`
}

