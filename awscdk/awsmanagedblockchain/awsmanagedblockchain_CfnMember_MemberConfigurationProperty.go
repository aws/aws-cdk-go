package awsmanagedblockchain


// Configuration properties of the member.
//
// Applies only to Hyperledger Fabric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberConfigurationProperty := &memberConfigurationProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	memberFrameworkConfiguration: &memberFrameworkConfigurationProperty{
//   		memberFabricConfiguration: &memberFabricConfigurationProperty{
//   			adminPassword: jsii.String("adminPassword"),
//   			adminUsername: jsii.String("adminUsername"),
//   		},
//   	},
//   }
//
type CfnMember_MemberConfigurationProperty struct {
	// The name of the member.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional description of the member.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration properties of the blockchain framework relevant to the member.
	MemberFrameworkConfiguration interface{} `field:"optional" json:"memberFrameworkConfiguration" yaml:"memberFrameworkConfiguration"`
}

