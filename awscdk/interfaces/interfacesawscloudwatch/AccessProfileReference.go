package interfacesawscloudwatch


// A reference to a AccessProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessProfileReference := &AccessProfileReference{
//   	AccessProfileArn: jsii.String("accessProfileArn"),
//   	ProfileId: jsii.String("profileId"),
//   	SpaceId: jsii.String("spaceId"),
//   }
//
type AccessProfileReference struct {
	// The ARN of the AccessProfile resource.
	AccessProfileArn *string `field:"required" json:"accessProfileArn" yaml:"accessProfileArn"`
	// The ProfileId of the AccessProfile resource.
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// The SpaceId of the AccessProfile resource.
	SpaceId *string `field:"required" json:"spaceId" yaml:"spaceId"`
}

