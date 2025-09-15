package awss3


// A reference to a AccessGrantsLocation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessGrantsLocationReference := &AccessGrantsLocationReference{
//   	AccessGrantsLocationArn: jsii.String("accessGrantsLocationArn"),
//   	AccessGrantsLocationId: jsii.String("accessGrantsLocationId"),
//   }
//
type AccessGrantsLocationReference struct {
	// The ARN of the AccessGrantsLocation resource.
	AccessGrantsLocationArn *string `field:"required" json:"accessGrantsLocationArn" yaml:"accessGrantsLocationArn"`
	// The AccessGrantsLocationId of the AccessGrantsLocation resource.
	AccessGrantsLocationId *string `field:"required" json:"accessGrantsLocationId" yaml:"accessGrantsLocationId"`
}

