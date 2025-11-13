package interfacesawssso


// A reference to a Instance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceReference := &InstanceReference{
//   	InstanceArn: jsii.String("instanceArn"),
//   }
//
type InstanceReference struct {
	// The InstanceArn of the Instance resource.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
}

