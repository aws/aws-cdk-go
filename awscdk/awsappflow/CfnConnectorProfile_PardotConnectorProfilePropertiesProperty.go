package awsappflow


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pardotConnectorProfilePropertiesProperty := &PardotConnectorProfilePropertiesProperty{
//   	BusinessUnitId: jsii.String("businessUnitId"),
//
//   	// the properties below are optional
//   	InstanceUrl: jsii.String("instanceUrl"),
//   	IsSandboxEnvironment: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html
//
type CfnConnectorProfile_PardotConnectorProfilePropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html#cfn-appflow-connectorprofile-pardotconnectorprofileproperties-businessunitid
	//
	BusinessUnitId *string `field:"required" json:"businessUnitId" yaml:"businessUnitId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html#cfn-appflow-connectorprofile-pardotconnectorprofileproperties-instanceurl
	//
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
	// Indicates whether the connector profile applies to a demo or production environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html#cfn-appflow-connectorprofile-pardotconnectorprofileproperties-issandboxenvironment
	//
	IsSandboxEnvironment interface{} `field:"optional" json:"isSandboxEnvironment" yaml:"isSandboxEnvironment"`
}

