package awssecretsmanager


// A reference to a RotationSchedule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rotationScheduleReference := &RotationScheduleReference{
//   	RotationScheduleId: jsii.String("rotationScheduleId"),
//   }
//
type RotationScheduleReference struct {
	// The Id of the RotationSchedule resource.
	RotationScheduleId *string `field:"required" json:"rotationScheduleId" yaml:"rotationScheduleId"`
}

