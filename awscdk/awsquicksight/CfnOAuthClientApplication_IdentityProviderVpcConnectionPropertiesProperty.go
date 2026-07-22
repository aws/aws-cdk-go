package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityProviderVpcConnectionPropertiesProperty := &IdentityProviderVpcConnectionPropertiesProperty{
//   	VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-oauthclientapplication-identityprovidervpcconnectionproperties.html
//
type CfnOAuthClientApplication_IdentityProviderVpcConnectionPropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-oauthclientapplication-identityprovidervpcconnectionproperties.html#cfn-quicksight-oauthclientapplication-identityprovidervpcconnectionproperties-vpcconnectionarn
	//
	VpcConnectionArn *string `field:"required" json:"vpcConnectionArn" yaml:"vpcConnectionArn"`
}

