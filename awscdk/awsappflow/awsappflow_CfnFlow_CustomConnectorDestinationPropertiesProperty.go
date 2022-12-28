package awsappflow


// The properties that are applied when the custom connector is being used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customConnectorDestinationPropertiesProperty := &customConnectorDestinationPropertiesProperty{
//   	entityName: jsii.String("entityName"),
//
//   	// the properties below are optional
//   	customProperties: map[string]*string{
//   		"customPropertiesKey": jsii.String("customProperties"),
//   	},
//   	errorHandlingConfig: &errorHandlingConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		failOnFirstError: jsii.Boolean(false),
//   	},
//   	idFieldNames: []*string{
//   		jsii.String("idFieldNames"),
//   	},
//   	writeOperationType: jsii.String("writeOperationType"),
//   }
//
type CfnFlow_CustomConnectorDestinationPropertiesProperty struct {
	// The entity specified in the custom connector as a destination in the flow.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// The custom properties that are specific to the connector when it's used as a destination in the flow.
	CustomProperties interface{} `field:"optional" json:"customProperties" yaml:"customProperties"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the custom connector as destination.
	ErrorHandlingConfig interface{} `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// The name of the field that Amazon AppFlow uses as an ID when performing a write operation such as update, delete, or upsert.
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// Specifies the type of write operation to be performed in the custom connector when it's used as destination.
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

