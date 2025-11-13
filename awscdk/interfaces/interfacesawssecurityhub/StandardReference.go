package interfacesawssecurityhub


// A reference to a Standard resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   standardReference := &StandardReference{
//   	StandardsSubscriptionArn: jsii.String("standardsSubscriptionArn"),
//   }
//
type StandardReference struct {
	// The StandardsSubscriptionArn of the Standard resource.
	StandardsSubscriptionArn *string `field:"required" json:"standardsSubscriptionArn" yaml:"standardsSubscriptionArn"`
}

