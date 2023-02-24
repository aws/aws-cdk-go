package awsappintegrations


// The event integration association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventIntegrationAssociationProperty := &EventIntegrationAssociationProperty{
//   	ClientAssociationMetadata: []interface{}{
//   		&MetadataProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ClientId: jsii.String("clientId"),
//   	EventBridgeRuleName: jsii.String("eventBridgeRuleName"),
//   	EventIntegrationAssociationArn: jsii.String("eventIntegrationAssociationArn"),
//   	EventIntegrationAssociationId: jsii.String("eventIntegrationAssociationId"),
//   }
//
type CfnEventIntegration_EventIntegrationAssociationProperty struct {
	// The metadata associated with the client.
	ClientAssociationMetadata interface{} `field:"optional" json:"clientAssociationMetadata" yaml:"clientAssociationMetadata"`
	// The identifier for the client that is associated with the event integration.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The name of the EventBridge rule.
	EventBridgeRuleName *string `field:"optional" json:"eventBridgeRuleName" yaml:"eventBridgeRuleName"`
	// The Amazon Resource Name (ARN) for the event integration association.
	EventIntegrationAssociationArn *string `field:"optional" json:"eventIntegrationAssociationArn" yaml:"eventIntegrationAssociationArn"`
	// The identifier for the event integration association.
	EventIntegrationAssociationId *string `field:"optional" json:"eventIntegrationAssociationId" yaml:"eventIntegrationAssociationId"`
}

