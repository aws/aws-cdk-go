package interfacesawsglue


// A reference to a UsageProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   usageProfileReference := &UsageProfileReference{
//   	UsageProfileName: jsii.String("usageProfileName"),
//   }
//
type UsageProfileReference struct {
	// The Name of the UsageProfile resource.
	UsageProfileName *string `field:"required" json:"usageProfileName" yaml:"usageProfileName"`
}

