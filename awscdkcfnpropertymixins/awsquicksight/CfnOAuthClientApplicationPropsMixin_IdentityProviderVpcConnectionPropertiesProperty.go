package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   identityProviderVpcConnectionPropertiesProperty := &IdentityProviderVpcConnectionPropertiesProperty{
//   	VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-oauthclientapplication-identityprovidervpcconnectionproperties.html
//
type CfnOAuthClientApplicationPropsMixin_IdentityProviderVpcConnectionPropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-oauthclientapplication-identityprovidervpcconnectionproperties.html#cfn-quicksight-oauthclientapplication-identityprovidervpcconnectionproperties-vpcconnectionarn
	//
	VpcConnectionArn *string `field:"optional" json:"vpcConnectionArn" yaml:"vpcConnectionArn"`
}

