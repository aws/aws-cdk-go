package interfacesawscloudwatch


// A reference to a AccessGrant resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessGrantReference := &AccessGrantReference{
//   	AccessGrantArn: jsii.String("accessGrantArn"),
//   }
//
type AccessGrantReference struct {
	// The Arn of the AccessGrant resource.
	AccessGrantArn *string `field:"required" json:"accessGrantArn" yaml:"accessGrantArn"`
}

