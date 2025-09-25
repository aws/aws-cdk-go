package awscontroltower


// A reference to a LandingZone resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   landingZoneReference := &LandingZoneReference{
//   	LandingZoneArn: jsii.String("landingZoneArn"),
//   	LandingZoneIdentifier: jsii.String("landingZoneIdentifier"),
//   }
//
type LandingZoneReference struct {
	// The ARN of the LandingZone resource.
	LandingZoneArn *string `field:"required" json:"landingZoneArn" yaml:"landingZoneArn"`
	// The LandingZoneIdentifier of the LandingZone resource.
	LandingZoneIdentifier *string `field:"required" json:"landingZoneIdentifier" yaml:"landingZoneIdentifier"`
}

