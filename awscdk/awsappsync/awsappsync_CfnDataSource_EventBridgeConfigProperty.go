package awsappsync


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBridgeConfigProperty := &eventBridgeConfigProperty{
//   	eventBusArn: jsii.String("eventBusArn"),
//   }
//
type CfnDataSource_EventBridgeConfigProperty struct {
	// `CfnDataSource.EventBridgeConfigProperty.EventBusArn`.
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
}

