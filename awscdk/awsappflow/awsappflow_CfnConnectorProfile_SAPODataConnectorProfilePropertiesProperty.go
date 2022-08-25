package awsappflow


// The connector-specific profile properties required when using SAPOData.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAPODataConnectorProfilePropertiesProperty := &sAPODataConnectorProfilePropertiesProperty{
//   	applicationHostUrl: jsii.String("applicationHostUrl"),
//   	applicationServicePath: jsii.String("applicationServicePath"),
//   	clientNumber: jsii.String("clientNumber"),
//   	logonLanguage: jsii.String("logonLanguage"),
//   	oAuthProperties: &oAuthPropertiesProperty{
//   		authCodeUrl: jsii.String("authCodeUrl"),
//   		oAuthScopes: []*string{
//   			jsii.String("oAuthScopes"),
//   		},
//   		tokenUrl: jsii.String("tokenUrl"),
//   	},
//   	portNumber: jsii.Number(123),
//   	privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   }
//
type CfnConnectorProfile_SAPODataConnectorProfilePropertiesProperty struct {
	// The location of the SAPOData resource.
	ApplicationHostUrl *string `field:"optional" json:"applicationHostUrl" yaml:"applicationHostUrl"`
	// The application path to catalog service.
	ApplicationServicePath *string `field:"optional" json:"applicationServicePath" yaml:"applicationServicePath"`
	// The client number for the client creating the connection.
	ClientNumber *string `field:"optional" json:"clientNumber" yaml:"clientNumber"`
	// The logon language of SAPOData instance.
	LogonLanguage *string `field:"optional" json:"logonLanguage" yaml:"logonLanguage"`
	// The SAPOData OAuth properties required for OAuth type authentication.
	OAuthProperties interface{} `field:"optional" json:"oAuthProperties" yaml:"oAuthProperties"`
	// The port number of the SAPOData instance.
	PortNumber *float64 `field:"optional" json:"portNumber" yaml:"portNumber"`
	// The SAPOData Private Link service name to be used for private data transfers.
	PrivateLinkServiceName *string `field:"optional" json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
}

