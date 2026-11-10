package interfacesawsroute53resolver


// A reference to a FirewallConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firewallConfigReference := &FirewallConfigReference{
//   	FirewallConfigArn: jsii.String("firewallConfigArn"),
//   }
//
type FirewallConfigReference struct {
	// The Arn of the FirewallConfig resource.
	FirewallConfigArn *string `field:"required" json:"firewallConfigArn" yaml:"firewallConfigArn"`
}

