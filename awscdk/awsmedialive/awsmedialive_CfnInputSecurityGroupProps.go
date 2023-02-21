package awsmedialive


// Properties for defining a `CfnInputSecurityGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnInputSecurityGroupProps := &cfnInputSecurityGroupProps{
//   	tags: tags,
//   	whitelistRules: []interface{}{
//   		&inputWhitelistRuleCidrProperty{
//   			cidr: jsii.String("cidr"),
//   		},
//   	},
//   }
//
type CfnInputSecurityGroupProps struct {
	// A collection of tags for this input security group.
	//
	// Each tag is a key-value pair.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The list of IPv4 CIDR addresses to include in the input security group as "allowed" addresses.
	WhitelistRules interface{} `field:"optional" json:"whitelistRules" yaml:"whitelistRules"`
}

