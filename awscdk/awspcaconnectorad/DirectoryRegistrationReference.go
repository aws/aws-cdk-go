package awspcaconnectorad


// A reference to a DirectoryRegistration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   directoryRegistrationReference := &DirectoryRegistrationReference{
//   	DirectoryRegistrationArn: jsii.String("directoryRegistrationArn"),
//   }
//
type DirectoryRegistrationReference struct {
	// The DirectoryRegistrationArn of the DirectoryRegistration resource.
	DirectoryRegistrationArn *string `field:"required" json:"directoryRegistrationArn" yaml:"directoryRegistrationArn"`
}

