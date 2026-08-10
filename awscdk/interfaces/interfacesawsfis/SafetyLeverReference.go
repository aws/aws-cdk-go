package interfacesawsfis


// A reference to a SafetyLever resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   safetyLeverReference := &SafetyLeverReference{
//   	SafetyLeverArn: jsii.String("safetyLeverArn"),
//   }
//
type SafetyLeverReference struct {
	// The Arn of the SafetyLever resource.
	SafetyLeverArn *string `field:"required" json:"safetyLeverArn" yaml:"safetyLeverArn"`
}

