package previewawsappflowmixins


// The properties that are applied when Amazon EventBridge is being used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventBridgeDestinationPropertiesProperty := &EventBridgeDestinationPropertiesProperty{
//   	ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		FailOnFirstError: jsii.Boolean(false),
//   	},
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-eventbridgedestinationproperties.html
//
type CfnFlowPropsMixin_EventBridgeDestinationPropertiesProperty struct {
	// The object specified in the Amplitude flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-eventbridgedestinationproperties.html#cfn-appflow-flow-eventbridgedestinationproperties-errorhandlingconfig
	//
	ErrorHandlingConfig interface{} `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// The object specified in the Amazon EventBridge flow destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-eventbridgedestinationproperties.html#cfn-appflow-flow-eventbridgedestinationproperties-object
	//
	Object *string `field:"optional" json:"object" yaml:"object"`
}

