package interfacesawsresiliencehubv2


// A reference to a UserJourney resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userJourneyReference := &UserJourneyReference{
//   	SystemIdentifier: jsii.String("systemIdentifier"),
//   	UserJourneyId: jsii.String("userJourneyId"),
//   }
//
type UserJourneyReference struct {
	// The SystemIdentifier of the UserJourney resource.
	SystemIdentifier *string `field:"required" json:"systemIdentifier" yaml:"systemIdentifier"`
	// The UserJourneyId of the UserJourney resource.
	UserJourneyId *string `field:"required" json:"userJourneyId" yaml:"userJourneyId"`
}

