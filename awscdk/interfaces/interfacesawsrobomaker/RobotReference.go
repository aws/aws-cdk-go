package interfacesawsrobomaker


// A reference to a Robot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   robotReference := &RobotReference{
//   	RobotArn: jsii.String("robotArn"),
//   }
//
type RobotReference struct {
	// The ARN of the Robot resource.
	RobotArn *string `field:"required" json:"robotArn" yaml:"robotArn"`
}

