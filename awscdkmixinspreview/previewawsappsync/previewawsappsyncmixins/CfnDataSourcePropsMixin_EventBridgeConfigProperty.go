package previewawsappsyncmixins


// The data source.
//
// This can be an API destination, resource, or AWS service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventBridgeConfigProperty := &EventBridgeConfigProperty{
//   	EventBusArn: jsii.String("eventBusArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-eventbridgeconfig.html
//
type CfnDataSourcePropsMixin_EventBridgeConfigProperty struct {
	// The event bus pipeline's ARN.
	//
	// For more information about event buses, see [EventBridge event buses](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-bus.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-eventbridgeconfig.html#cfn-appsync-datasource-eventbridgeconfig-eventbusarn
	//
	EventBusArn *string `field:"optional" json:"eventBusArn" yaml:"eventBusArn"`
}

