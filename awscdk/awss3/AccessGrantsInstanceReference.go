package awss3


// A reference to a AccessGrantsInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessGrantsInstanceReference := &AccessGrantsInstanceReference{
//   	AccessGrantsInstanceArn: jsii.String("accessGrantsInstanceArn"),
//   }
//
type AccessGrantsInstanceReference struct {
	// The AccessGrantsInstanceArn of the AccessGrantsInstance resource.
	AccessGrantsInstanceArn *string `field:"required" json:"accessGrantsInstanceArn" yaml:"accessGrantsInstanceArn"`
}

