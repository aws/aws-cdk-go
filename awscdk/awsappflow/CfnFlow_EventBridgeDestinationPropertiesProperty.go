package awsappflow


// The properties that are applied when Amazon EventBridge is being used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBridgeDestinationPropertiesProperty := &EventBridgeDestinationPropertiesProperty{
//   	Object: jsii.String("object"),
//
//   	// the properties below are optional
//   	ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		FailOnFirstError: jsii.Boolean(false),
//   	},
//   }
//
type CfnFlow_EventBridgeDestinationPropertiesProperty struct {
	// The object specified in the Amazon EventBridge flow destination.
	Object *string `field:"required" json:"object" yaml:"object"`
	// The object specified in the Amplitude flow source.
	ErrorHandlingConfig interface{} `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

