package awsappflow


// The properties that are applied when the custom connector is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customConnectorSourcePropertiesProperty := &CustomConnectorSourcePropertiesProperty{
//   	EntityName: jsii.String("entityName"),
//
//   	// the properties below are optional
//   	CustomProperties: map[string]*string{
//   		"customPropertiesKey": jsii.String("customProperties"),
//   	},
//   	DataTransferApi: &DataTransferApiProperty{
//   		Name: jsii.String("name"),
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-customconnectorsourceproperties.html
//
type CfnFlow_CustomConnectorSourcePropertiesProperty struct {
	// The entity specified in the custom connector as a source in the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-customconnectorsourceproperties.html#cfn-appflow-flow-customconnectorsourceproperties-entityname
	//
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// Custom properties that are required to use the custom connector as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-customconnectorsourceproperties.html#cfn-appflow-flow-customconnectorsourceproperties-customproperties
	//
	CustomProperties interface{} `field:"optional" json:"customProperties" yaml:"customProperties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-customconnectorsourceproperties.html#cfn-appflow-flow-customconnectorsourceproperties-datatransferapi
	//
	DataTransferApi interface{} `field:"optional" json:"dataTransferApi" yaml:"dataTransferApi"`
}

