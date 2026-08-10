package interfacesawssigner


// A reference to a SigningJob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signingJobReference := &SigningJobReference{
//   	SigningJobArn: jsii.String("signingJobArn"),
//   }
//
type SigningJobReference struct {
	// The Arn of the SigningJob resource.
	SigningJobArn *string `field:"required" json:"signingJobArn" yaml:"signingJobArn"`
}

