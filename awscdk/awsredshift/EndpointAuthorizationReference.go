package awsredshift


// A reference to a EndpointAuthorization resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointAuthorizationReference := &EndpointAuthorizationReference{
//   	Account: jsii.String("account"),
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   }
//
type EndpointAuthorizationReference struct {
	// The Account of the EndpointAuthorization resource.
	Account *string `field:"required" json:"account" yaml:"account"`
	// The ClusterIdentifier of the EndpointAuthorization resource.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
}

