package mixinsawsappflow


// The connector-specific profile properties required when using Salesforce.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   salesforceConnectorProfilePropertiesProperty := &SalesforceConnectorProfilePropertiesProperty{
//   	InstanceUrl: jsii.String("instanceUrl"),
//   	IsSandboxEnvironment: jsii.Boolean(false),
//   	UsePrivateLinkForMetadataAndAuthorization: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofileproperties.html
//
type CfnConnectorProfilePropsMixin_SalesforceConnectorProfilePropertiesProperty struct {
	// The location of the Salesforce resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofileproperties.html#cfn-appflow-connectorprofile-salesforceconnectorprofileproperties-instanceurl
	//
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
	// Indicates whether the connector profile applies to a sandbox or production environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofileproperties.html#cfn-appflow-connectorprofile-salesforceconnectorprofileproperties-issandboxenvironment
	//
	IsSandboxEnvironment interface{} `field:"optional" json:"isSandboxEnvironment" yaml:"isSandboxEnvironment"`
	// If the connection mode for the connector profile is private, this parameter sets whether Amazon AppFlow uses the private network to send metadata and authorization calls to Salesforce.
	//
	// Amazon AppFlow sends private calls through AWS PrivateLink . These calls travel through AWS infrastructure without being exposed to the public internet.
	//
	// Set either of the following values:
	//
	// - **true** - Amazon AppFlow sends all calls to Salesforce over the private network.
	//
	// These private calls are:
	//
	// - Calls to get metadata about your Salesforce records. This metadata describes your Salesforce objects and their fields.
	// - Calls to get or refresh access tokens that allow Amazon AppFlow to access your Salesforce records.
	// - Calls to transfer your Salesforce records as part of a flow run.
	// - **false** - The default value. Amazon AppFlow sends some calls to Salesforce privately and other calls over the public internet.
	//
	// The public calls are:
	//
	// - Calls to get metadata about your Salesforce records.
	// - Calls to get or refresh access tokens.
	//
	// The private calls are:
	//
	// - Calls to transfer your Salesforce records as part of a flow run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofileproperties.html#cfn-appflow-connectorprofile-salesforceconnectorprofileproperties-useprivatelinkformetadataandauthorization
	//
	UsePrivateLinkForMetadataAndAuthorization interface{} `field:"optional" json:"usePrivateLinkForMetadataAndAuthorization" yaml:"usePrivateLinkForMetadataAndAuthorization"`
}

