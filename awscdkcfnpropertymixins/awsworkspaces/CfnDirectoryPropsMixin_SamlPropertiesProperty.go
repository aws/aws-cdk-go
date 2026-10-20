package awsworkspaces


// Describes the enablement status, user access URL, and relay state parameter name that are used for configuring federation with an SAML 2.0 identity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   samlPropertiesProperty := &SamlPropertiesProperty{
//   	RelayStateParameterName: jsii.String("relayStateParameterName"),
//   	Status: jsii.String("status"),
//   	UserAccessUrl: jsii.String("userAccessUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-samlproperties.html
//
type CfnDirectoryPropsMixin_SamlPropertiesProperty struct {
	// The relay state parameter name supported by the SAML 2.0 identity provider (IdP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-samlproperties.html#cfn-workspaces-directory-samlproperties-relaystateparametername
	//
	RelayStateParameterName *string `field:"optional" json:"relayStateParameterName" yaml:"relayStateParameterName"`
	// Indicates the status of SAML 2.0 authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-samlproperties.html#cfn-workspaces-directory-samlproperties-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The SAML 2.0 identity provider (IdP) user access URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-samlproperties.html#cfn-workspaces-directory-samlproperties-useraccessurl
	//
	UserAccessUrl *string `field:"optional" json:"userAccessUrl" yaml:"userAccessUrl"`
}

