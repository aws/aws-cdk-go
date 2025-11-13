package interfacesawsshield


// A reference to a ProactiveEngagement resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   proactiveEngagementReference := &ProactiveEngagementReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type ProactiveEngagementReference struct {
	// The AccountId of the ProactiveEngagement resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

