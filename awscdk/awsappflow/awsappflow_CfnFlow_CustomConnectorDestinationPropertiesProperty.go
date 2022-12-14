package awsappflow


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
	// `CfnFlow.CustomConnectorDestinationPropertiesProperty.EntityName`.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// `CfnFlow.CustomConnectorDestinationPropertiesProperty.CustomProperties`.
	CustomProperties interface{} `field:"optional" json:"customProperties" yaml:"customProperties"`
	// `CfnFlow.CustomConnectorDestinationPropertiesProperty.ErrorHandlingConfig`.
	ErrorHandlingConfig interface{} `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// `CfnFlow.CustomConnectorDestinationPropertiesProperty.IdFieldNames`.
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// `CfnFlow.CustomConnectorDestinationPropertiesProperty.WriteOperationType`.
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

