package awsappflow


// The connector-specific profile properties required when using SAPOData.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAPODataConnectorProfilePropertiesProperty := &SAPODataConnectorProfilePropertiesProperty{
//   	ApplicationHostUrl: jsii.String("applicationHostUrl"),
//   	ApplicationServicePath: jsii.String("applicationServicePath"),
//   	ClientNumber: jsii.String("clientNumber"),
//   	DisableSso: jsii.Boolean(false),
//   	LogonLanguage: jsii.String("logonLanguage"),
//   	OAuthProperties: &OAuthPropertiesProperty{
//   		AuthCodeUrl: jsii.String("authCodeUrl"),
//   		OAuthScopes: []*string{
//   			jsii.String("oAuthScopes"),
//   		},
//   		TokenUrl: jsii.String("tokenUrl"),
//   	},
//   	PortNumber: jsii.Number(123),
//   	PrivateLinkServiceName: jsii.String("privateLinkServiceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-sapodataconnectorprofileproperties.html
//
type CfnConnectorProfile_SAPODataConnectorProfilePropertiesProperty struct {
	// The location of the SAPOData resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-sapodataconnectorprofileproperties.html#cfn-appflow-connectorprofile-sapodataconnectorprofileproperties-applicationhosturl
	//
	ApplicationHostUrl *string `field:"optional" json:"applicationHostUrl" yaml:"applicationHostUrl"`
	// The application path to catalog service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-sapodataconnectorprofileproperties.html#cfn-appflow-connectorprofile-sapodataconnectorprofileproperties-applicationservicepath
	//
	ApplicationServicePath *string `field:"optional" json:"applicationServicePath" yaml:"applicationServicePath"`
	// The client number for the client creating the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-sapodataconnectorprofileproperties.html#cfn-appflow-connectorprofile-sapodataconnectorprofileproperties-clientnumber
	//
	ClientNumber *string `field:"optional" json:"clientNumber" yaml:"clientNumber"`
	// If you set this parameter to true, Amazon AppFlow bypasses the single sign-on (SSO) settings in your SAP account when it accesses your SAP OData instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-sapodataconnectorprofileproperties.html#cfn-appflow-connectorprofile-sapodataconnectorprofileproperties-disablesso
	//
	DisableSso interface{} `field:"optional" json:"disableSso" yaml:"disableSso"`
	// The logon language of SAPOData instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-sapodataconnectorprofileproperties.html#cfn-appflow-connectorprofile-sapodataconnectorprofileproperties-logonlanguage
	//
	LogonLanguage *string `field:"optional" json:"logonLanguage" yaml:"logonLanguage"`
	// The SAPOData OAuth properties required for OAuth type authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-sapodataconnectorprofileproperties.html#cfn-appflow-connectorprofile-sapodataconnectorprofileproperties-oauthproperties
	//
	OAuthProperties interface{} `field:"optional" json:"oAuthProperties" yaml:"oAuthProperties"`
	// The port number of the SAPOData instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-sapodataconnectorprofileproperties.html#cfn-appflow-connectorprofile-sapodataconnectorprofileproperties-portnumber
	//
	PortNumber *float64 `field:"optional" json:"portNumber" yaml:"portNumber"`
	// The SAPOData Private Link service name to be used for private data transfers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-sapodataconnectorprofileproperties.html#cfn-appflow-connectorprofile-sapodataconnectorprofileproperties-privatelinkservicename
	//
	PrivateLinkServiceName *string `field:"optional" json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
}

