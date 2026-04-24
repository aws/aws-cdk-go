package interfacesawsrobomaker


// A reference to a RobotApplicationVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   robotApplicationVersionReference := &RobotApplicationVersionReference{
//   	RobotApplicationVersionArn: jsii.String("robotApplicationVersionArn"),
//   }
//
type RobotApplicationVersionReference struct {
	// The Arn of the RobotApplicationVersion resource.
	RobotApplicationVersionArn *string `field:"required" json:"robotApplicationVersionArn" yaml:"robotApplicationVersionArn"`
}

