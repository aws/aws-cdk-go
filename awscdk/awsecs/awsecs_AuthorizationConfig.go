package awsecs


// The authorization configuration details for the Amazon EFS file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationConfig := &authorizationConfig{
//   	accessPointId: jsii.String("accessPointId"),
//   	iam: jsii.String("iam"),
//   }
//
type AuthorizationConfig struct {
	// The access point ID to use.
	//
	// If an access point is specified, the root directory value will be
	// relative to the directory set for the access point.
	// If specified, transit encryption must be enabled in the EFSVolumeConfiguration.
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// Whether or not to use the Amazon ECS task IAM role defined in a task definition when mounting the Amazon EFS file system.
	//
	// If enabled, transit encryption must be enabled in the EFSVolumeConfiguration.
	//
	// Valid values: ENABLED | DISABLED.
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
}

