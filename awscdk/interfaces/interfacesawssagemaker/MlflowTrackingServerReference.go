package interfacesawssagemaker


// A reference to a MlflowTrackingServer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mlflowTrackingServerReference := &MlflowTrackingServerReference{
//   	TrackingServerName: jsii.String("trackingServerName"),
//   }
//
type MlflowTrackingServerReference struct {
	// The TrackingServerName of the MlflowTrackingServer resource.
	TrackingServerName *string `field:"required" json:"trackingServerName" yaml:"trackingServerName"`
}

