package awsappflow


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customConnectorSourcePropertiesProperty := &customConnectorSourcePropertiesProperty{
//   	entityName: jsii.String("entityName"),
//
//   	// the properties below are optional
//   	customProperties: map[string]*string{
//   		"customPropertiesKey": jsii.String("customProperties"),
//   	},
//   }
//
type CfnFlow_CustomConnectorSourcePropertiesProperty struct {
	// `CfnFlow.CustomConnectorSourcePropertiesProperty.EntityName`.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// `CfnFlow.CustomConnectorSourcePropertiesProperty.CustomProperties`.
	CustomProperties interface{} `field:"optional" json:"customProperties" yaml:"customProperties"`
}

