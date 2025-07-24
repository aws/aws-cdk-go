package awslogsdestinations


// Options that may be provided to LambdaDestination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaDestinationOptions := &LambdaDestinationOptions{
//   	AddPermissions: jsii.Boolean(false),
//   }
//
type LambdaDestinationOptions struct {
	// Whether or not to add Lambda Permissions.
	// Default: true.
	//
	AddPermissions *bool `field:"optional" json:"addPermissions" yaml:"addPermissions"`
}

