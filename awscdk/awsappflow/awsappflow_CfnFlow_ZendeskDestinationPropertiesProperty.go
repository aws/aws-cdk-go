package awsappflow


// The properties that are applied when Zendesk is used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zendeskDestinationPropertiesProperty := &zendeskDestinationPropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
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
type CfnFlow_ZendeskDestinationPropertiesProperty struct {
	// The object specified in the Zendesk flow destination.
	Object *string `field:"required" json:"object" yaml:"object"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
	ErrorHandlingConfig interface{} `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// A list of field names that can be used as an ID field when performing a write operation.
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// The possible write operations in the destination connector.
	//
	// When this value is not provided, this defaults to the `INSERT` operation.
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

