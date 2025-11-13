package interfacesawss3


// A reference to a AccessGrant resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessGrantReference := &AccessGrantReference{
//   	AccessGrantArn: jsii.String("accessGrantArn"),
//   	AccessGrantId: jsii.String("accessGrantId"),
//   }
//
type AccessGrantReference struct {
	// The ARN of the AccessGrant resource.
	AccessGrantArn *string `field:"required" json:"accessGrantArn" yaml:"accessGrantArn"`
	// The AccessGrantId of the AccessGrant resource.
	AccessGrantId *string `field:"required" json:"accessGrantId" yaml:"accessGrantId"`
}

