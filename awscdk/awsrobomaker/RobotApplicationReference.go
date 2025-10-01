package awsrobomaker


// A reference to a RobotApplication resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   robotApplicationReference := &RobotApplicationReference{
//   	RobotApplicationArn: jsii.String("robotApplicationArn"),
//   }
//
type RobotApplicationReference struct {
	// The Arn of the RobotApplication resource.
	RobotApplicationArn *string `field:"required" json:"robotApplicationArn" yaml:"robotApplicationArn"`
}

