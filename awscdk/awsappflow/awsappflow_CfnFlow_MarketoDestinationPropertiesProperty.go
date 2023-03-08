package awsappflow


// The properties that Amazon AppFlow applies when you use Marketo as a flow destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   marketoDestinationPropertiesProperty := &marketoDestinationPropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	errorHandlingConfig: &errorHandlingConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		failOnFirstError: jsii.Boolean(false),
//   	},
//   }
//
type CfnFlow_MarketoDestinationPropertiesProperty struct {
	// The object specified in the Marketo flow destination.
	Object *string `field:"required" json:"object" yaml:"object"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
	ErrorHandlingConfig interface{} `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

