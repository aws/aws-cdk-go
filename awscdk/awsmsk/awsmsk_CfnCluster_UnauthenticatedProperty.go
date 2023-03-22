package awsmsk


// Details for allowing no client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   unauthenticatedProperty := &unauthenticatedProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnCluster_UnauthenticatedProperty struct {
	// Unauthenticated is enabled or not.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

