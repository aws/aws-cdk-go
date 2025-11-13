package interfacesawsnotificationscontacts


// A reference to a EmailContact resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailContactReference := &EmailContactReference{
//   	EmailContactArn: jsii.String("emailContactArn"),
//   }
//
type EmailContactReference struct {
	// The Arn of the EmailContact resource.
	EmailContactArn *string `field:"required" json:"emailContactArn" yaml:"emailContactArn"`
}

