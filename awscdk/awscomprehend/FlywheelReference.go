package awscomprehend


// A reference to a Flywheel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flywheelReference := &FlywheelReference{
//   	FlywheelArn: jsii.String("flywheelArn"),
//   }
//
type FlywheelReference struct {
	// The Arn of the Flywheel resource.
	FlywheelArn *string `field:"required" json:"flywheelArn" yaml:"flywheelArn"`
}

