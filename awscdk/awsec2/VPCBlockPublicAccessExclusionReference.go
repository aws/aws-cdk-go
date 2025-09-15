package awsec2


// A reference to a VPCBlockPublicAccessExclusion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCBlockPublicAccessExclusionReference := &VPCBlockPublicAccessExclusionReference{
//   	ExclusionId: jsii.String("exclusionId"),
//   }
//
type VPCBlockPublicAccessExclusionReference struct {
	// The ExclusionId of the VPCBlockPublicAccessExclusion resource.
	ExclusionId *string `field:"required" json:"exclusionId" yaml:"exclusionId"`
}

