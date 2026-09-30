package interfacesawsredshift


// A reference to a QEV2IdcApplication resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   qEV2IdcApplicationReference := &QEV2IdcApplicationReference{
//   	Qev2IdcApplicationArn: jsii.String("qev2IdcApplicationArn"),
//   }
//
type QEV2IdcApplicationReference struct {
	// The Qev2IdcApplicationArn of the QEV2IdcApplication resource.
	Qev2IdcApplicationArn *string `field:"required" json:"qev2IdcApplicationArn" yaml:"qev2IdcApplicationArn"`
}

