package awsmanagedblockchain


// Configuration properties for Hyperledger Fabric for a member in a Managed Blockchain network using the Hyperledger Fabric framework.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberFabricConfigurationProperty := &memberFabricConfigurationProperty{
//   	adminPassword: jsii.String("adminPassword"),
//   	adminUsername: jsii.String("adminUsername"),
//   }
//
type CfnMember_MemberFabricConfigurationProperty struct {
	// The password for the member's initial administrative user.
	//
	// The `AdminPassword` must be at least eight characters long and no more than 32 characters. It must contain at least one uppercase letter, one lowercase letter, and one digit. It cannot have a single quotation mark (‘), a double quotation marks (“), a forward slash(/), a backward slash(\), @, or a space.
	AdminPassword *string `field:"required" json:"adminPassword" yaml:"adminPassword"`
	// The user name for the member's initial administrative user.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
}

