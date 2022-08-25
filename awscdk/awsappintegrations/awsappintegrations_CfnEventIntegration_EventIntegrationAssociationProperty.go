package awsappintegrations


// The event integration association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventIntegrationAssociationProperty := &eventIntegrationAssociationProperty{
//   	clientAssociationMetadata: []interface{}{
//   		&metadataProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	clientId: jsii.String("clientId"),
//   	eventBridgeRuleName: jsii.String("eventBridgeRuleName"),
//   	eventIntegrationAssociationArn: jsii.String("eventIntegrationAssociationArn"),
//   	eventIntegrationAssociationId: jsii.String("eventIntegrationAssociationId"),
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

