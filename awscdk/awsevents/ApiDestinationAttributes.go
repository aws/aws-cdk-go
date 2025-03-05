package awsevents


// The properties to import an existing Api Destination.
//
// Example:
//   connection := events.Connection_FromEventBusArn(this, jsii.String("Connection"), jsii.String("arn:aws:events:us-east-1:123456789012:event-bus/EventBusName"), jsii.String("arn:aws:secretsmanager:us-east-1:123456789012:secret:SecretName-f3gDy9"))
//
//   apiDestinationArn := "arn:aws:events:us-east-1:123456789012:api-destination/DestinationName"
//   destination := events.ApiDestination_FromApiDestinationAttributes(this, jsii.String("Destination"), &ApiDestinationAttributes{
//   	ApiDestinationArn: jsii.String(ApiDestinationArn),
//   	Connection: Connection,
//   })
//
//   rule := events.NewRule(this, jsii.String("OtherRule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
//   	Targets: []iRuleTarget{
//   		targets.NewApiDestination(destination),
//   	},
//   })
//
type ApiDestinationAttributes struct {
	// The ARN of the Api Destination.
	ApiDestinationArn *string `field:"required" json:"apiDestinationArn" yaml:"apiDestinationArn"`
	// The Connection to associate with the Api Destination.
	Connection IConnection `field:"required" json:"connection" yaml:"connection"`
}

