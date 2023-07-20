package awsappflow


// The connector-specific profile properties required when using Salesforce.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceConnectorProfilePropertiesProperty := &SalesforceConnectorProfilePropertiesProperty{
//   	InstanceUrl: jsii.String("instanceUrl"),
//   	IsSandboxEnvironment: jsii.Boolean(false),
//   	UsePrivateLinkForMetadataAndAuthorization: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofileproperties.html
//
type CfnConnectorProfile_SalesforceConnectorProfilePropertiesProperty struct {
	// The location of the Salesforce resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofileproperties.html#cfn-appflow-connectorprofile-salesforceconnectorprofileproperties-instanceurl
	//
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
	// Indicates whether the connector profile applies to a sandbox or production environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofileproperties.html#cfn-appflow-connectorprofile-salesforceconnectorprofileproperties-issandboxenvironment
	//
	IsSandboxEnvironment interface{} `field:"optional" json:"isSandboxEnvironment" yaml:"isSandboxEnvironment"`
	// Indicates whether to make Metadata And Authorization calls over Pivate Network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofileproperties.html#cfn-appflow-connectorprofile-salesforceconnectorprofileproperties-useprivatelinkformetadataandauthorization
	//
	UsePrivateLinkForMetadataAndAuthorization interface{} `field:"optional" json:"usePrivateLinkForMetadataAndAuthorization" yaml:"usePrivateLinkForMetadataAndAuthorization"`
}

