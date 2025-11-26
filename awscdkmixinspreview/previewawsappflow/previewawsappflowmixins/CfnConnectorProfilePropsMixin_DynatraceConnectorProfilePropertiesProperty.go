package previewawsappflowmixins


// The connector-specific profile properties required by Dynatrace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dynatraceConnectorProfilePropertiesProperty := &DynatraceConnectorProfilePropertiesProperty{
//   	InstanceUrl: jsii.String("instanceUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-dynatraceconnectorprofileproperties.html
//
type CfnConnectorProfilePropsMixin_DynatraceConnectorProfilePropertiesProperty struct {
	// The location of the Dynatrace resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-dynatraceconnectorprofileproperties.html#cfn-appflow-connectorprofile-dynatraceconnectorprofileproperties-instanceurl
	//
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
}

