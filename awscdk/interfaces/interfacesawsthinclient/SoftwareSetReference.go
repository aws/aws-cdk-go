package interfacesawsthinclient


// A reference to a SoftwareSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   softwareSetReference := &SoftwareSetReference{
//   	SoftwareSetArn: jsii.String("softwareSetArn"),
//   }
//
type SoftwareSetReference struct {
	// The Arn of the SoftwareSet resource.
	SoftwareSetArn *string `field:"required" json:"softwareSetArn" yaml:"softwareSetArn"`
}

