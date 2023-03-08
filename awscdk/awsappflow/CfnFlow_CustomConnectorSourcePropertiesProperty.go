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
//   }
//
type CfnFlow_CustomConnectorSourcePropertiesProperty struct {
	// The entity specified in the custom connector as a source in the flow.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// Custom properties that are required to use the custom connector as a source.
	CustomProperties interface{} `field:"optional" json:"customProperties" yaml:"customProperties"`
}

