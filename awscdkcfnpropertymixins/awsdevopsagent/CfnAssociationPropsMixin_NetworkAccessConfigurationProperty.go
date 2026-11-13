package awsdevopsagent


// Specifies how AWS DevOps Agent reaches your application using a release management environment.
//
// Exactly one access type must be specified.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   networkAccessConfigurationProperty := &NetworkAccessConfigurationProperty{
//   	PrivateAccess: &PrivateNetworkAccessProperty{
//   		PrivateConnectionName: jsii.String("privateConnectionName"),
//   		RuntimeRoleArn: jsii.String("runtimeRoleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-networkaccessconfiguration.html
//
type CfnAssociationPropsMixin_NetworkAccessConfigurationProperty struct {
	// Private network access to your application inside a VPC, reached through a private connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-networkaccessconfiguration.html#cfn-devopsagent-association-networkaccessconfiguration-privateaccess
	//
	PrivateAccess interface{} `field:"optional" json:"privateAccess" yaml:"privateAccess"`
}

