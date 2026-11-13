package interfacesawscloudwatch


// A reference to a Alert resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alertReference := &AlertReference{
//   	AlertArn: jsii.String("alertArn"),
//   	AlertId: jsii.String("alertId"),
//   	SpaceId: jsii.String("spaceId"),
//   }
//
type AlertReference struct {
	// The ARN of the Alert resource.
	AlertArn *string `field:"required" json:"alertArn" yaml:"alertArn"`
	// The AlertId of the Alert resource.
	AlertId *string `field:"required" json:"alertId" yaml:"alertId"`
	// The SpaceId of the Alert resource.
	SpaceId *string `field:"required" json:"spaceId" yaml:"spaceId"`
}

