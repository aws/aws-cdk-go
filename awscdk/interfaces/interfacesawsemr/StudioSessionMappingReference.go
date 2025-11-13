package interfacesawsemr


// A reference to a StudioSessionMapping resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   studioSessionMappingReference := &StudioSessionMappingReference{
//   	IdentityName: jsii.String("identityName"),
//   	IdentityType: jsii.String("identityType"),
//   	StudioId: jsii.String("studioId"),
//   }
//
type StudioSessionMappingReference struct {
	// The IdentityName of the StudioSessionMapping resource.
	IdentityName *string `field:"required" json:"identityName" yaml:"identityName"`
	// The IdentityType of the StudioSessionMapping resource.
	IdentityType *string `field:"required" json:"identityType" yaml:"identityType"`
	// The StudioId of the StudioSessionMapping resource.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
}

