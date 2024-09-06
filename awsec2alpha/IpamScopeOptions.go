package awsec2alpha


// Being used in IPAM class to add pools to default scope created by IPAM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ec2_alpha "github.com/aws/aws-cdk-go/awsec2alpha"
//
//   ipamScopeOptions := &IpamScopeOptions{
//   	IpamScopeName: jsii.String("ipamScopeName"),
//   }
//
// Experimental.
type IpamScopeOptions struct {
	// IPAM scope name that will be used for tagging.
	// Default: none.
	//
	// Experimental.
	IpamScopeName *string `field:"optional" json:"ipamScopeName" yaml:"ipamScopeName"`
}

