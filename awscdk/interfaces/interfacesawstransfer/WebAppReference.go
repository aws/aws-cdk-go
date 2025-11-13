package interfacesawstransfer


// A reference to a WebApp resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webAppReference := &WebAppReference{
//   	WebAppArn: jsii.String("webAppArn"),
//   }
//
type WebAppReference struct {
	// The Arn of the WebApp resource.
	WebAppArn *string `field:"required" json:"webAppArn" yaml:"webAppArn"`
}

