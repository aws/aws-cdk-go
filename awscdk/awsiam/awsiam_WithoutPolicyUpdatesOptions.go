package awsiam


// Options for the `withoutPolicyUpdates()` modifier of a Role.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   withoutPolicyUpdatesOptions := &withoutPolicyUpdatesOptions{
//   	addGrantsToResources: jsii.Boolean(false),
//   }
//
type WithoutPolicyUpdatesOptions struct {
	// Add grants to resources instead of dropping them.
	//
	// If this is `false` or not specified, grant permissions added to this role are ignored.
	// It is your own responsibility to make sure the role has the required permissions.
	//
	// If this is `true`, any grant permissions will be added to the resource instead.
	AddGrantsToResources *bool `field:"optional" json:"addGrantsToResources" yaml:"addGrantsToResources"`
}

