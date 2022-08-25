package awsmedialive


// An IPv4 CIDR range to include in this input security group.
//
// The parent of this entity is InputSecurityGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputWhitelistRuleCidrProperty := &inputWhitelistRuleCidrProperty{
//   	cidr: jsii.String("cidr"),
//   }
//
type CfnInputSecurityGroup_InputWhitelistRuleCidrProperty struct {
	// An IPv4 CIDR range to include in this input security group.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

