package awsdevopsagent


// Configuration for a release management environment.
//
// A release management environment defines how AWS DevOps Agent reaches your application, and other associations reference it to run automatic verification testing against that application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   releaseManagementConfigurationProperty := &ReleaseManagementConfigurationProperty{
//   	Name: jsii.String("name"),
//   	NetworkAccess: &NetworkAccessConfigurationProperty{
//   		PrivateAccess: &PrivateNetworkAccessProperty{
//   			PrivateConnectionName: jsii.String("privateConnectionName"),
//   			RuntimeRoleArn: jsii.String("runtimeRoleArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-releasemanagementconfiguration.html
//
type CfnAssociationPropsMixin_ReleaseManagementConfigurationProperty struct {
	// The name for this release management environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-releasemanagementconfiguration.html#cfn-devopsagent-association-releasemanagementconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies how AWS DevOps Agent reaches your application using a release management environment.
	//
	// Exactly one access type must be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-releasemanagementconfiguration.html#cfn-devopsagent-association-releasemanagementconfiguration-networkaccess
	//
	NetworkAccess interface{} `field:"optional" json:"networkAccess" yaml:"networkAccess"`
}

