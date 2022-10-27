package awsappflow


// The properties that are applied when Salesforce is being used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceDestinationPropertiesProperty := &salesforceDestinationPropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	dataTransferApi: jsii.String("dataTransferApi"),
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
type CfnFlow_SalesforceDestinationPropertiesProperty struct {
	// The object specified in the Salesforce flow destination.
	Object *string `field:"required" json:"object" yaml:"object"`
	// `CfnFlow.SalesforceDestinationPropertiesProperty.DataTransferApi`.
	DataTransferApi *string `field:"optional" json:"dataTransferApi" yaml:"dataTransferApi"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the Salesforce destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
	ErrorHandlingConfig interface{} `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// The name of the field that Amazon AppFlow uses as an ID when performing a write operation such as update or delete.
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// This specifies the type of write operation to be performed in Salesforce.
	//
	// When the value is `UPSERT` , then `idFieldNames` is required.
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

