package previewawsmanagedblockchainmixins


// Configuration properties of the member.
//
// Applies only to Hyperledger Fabric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   memberConfigurationProperty := &MemberConfigurationProperty{
//   	Description: jsii.String("description"),
//   	MemberFrameworkConfiguration: &MemberFrameworkConfigurationProperty{
//   		MemberFabricConfiguration: &MemberFabricConfigurationProperty{
//   			AdminPassword: jsii.String("adminPassword"),
//   			AdminUsername: jsii.String("adminUsername"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-memberconfiguration.html
//
type CfnMemberPropsMixin_MemberConfigurationProperty struct {
	// An optional description of the member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-memberconfiguration.html#cfn-managedblockchain-member-memberconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration properties of the blockchain framework relevant to the member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-memberconfiguration.html#cfn-managedblockchain-member-memberconfiguration-memberframeworkconfiguration
	//
	MemberFrameworkConfiguration interface{} `field:"optional" json:"memberFrameworkConfiguration" yaml:"memberFrameworkConfiguration"`
	// The name of the member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-memberconfiguration.html#cfn-managedblockchain-member-memberconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

