package previewawsappflowmixins


// The connector-specific profile properties required when using Salesforce Pardot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pardotConnectorProfilePropertiesProperty := &PardotConnectorProfilePropertiesProperty{
//   	BusinessUnitId: jsii.String("businessUnitId"),
//   	InstanceUrl: jsii.String("instanceUrl"),
//   	IsSandboxEnvironment: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html
//
type CfnConnectorProfilePropsMixin_PardotConnectorProfilePropertiesProperty struct {
	// The business unit id of Salesforce Pardot instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html#cfn-appflow-connectorprofile-pardotconnectorprofileproperties-businessunitid
	//
	BusinessUnitId *string `field:"optional" json:"businessUnitId" yaml:"businessUnitId"`
	// The location of the Salesforce Pardot resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html#cfn-appflow-connectorprofile-pardotconnectorprofileproperties-instanceurl
	//
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
	// Indicates whether the connector profile applies to a sandbox or production environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html#cfn-appflow-connectorprofile-pardotconnectorprofileproperties-issandboxenvironment
	//
	IsSandboxEnvironment interface{} `field:"optional" json:"isSandboxEnvironment" yaml:"isSandboxEnvironment"`
}

