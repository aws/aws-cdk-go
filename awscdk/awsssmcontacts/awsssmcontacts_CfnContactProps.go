package awsssmcontacts


// Properties for defining a `CfnContact`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContactProps := &cfnContactProps{
//   	alias: jsii.String("alias"),
//   	displayName: jsii.String("displayName"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	plan: []interface{}{
//   		&stageProperty{
//   			durationInMinutes: jsii.Number(123),
//
//   			// the properties below are optional
//   			targets: []interface{}{
//   				&targetsProperty{
//   					channelTargetInfo: &channelTargetInfoProperty{
//   						channelId: jsii.String("channelId"),
//   						retryIntervalInMinutes: jsii.Number(123),
//   					},
//   					contactTargetInfo: &contactTargetInfoProperty{
//   						contactId: jsii.String("contactId"),
//   						isEssential: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnContactProps struct {
	// The unique and identifiable alias of the contact or escalation plan.
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The full name of the contact or escalation plan.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Refers to the type of contact.
	//
	// A single contact is type `PERSONAL` and an escalation plan is type `ESCALATION` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// A list of stages.
	//
	// A contact has an engagement plan with stages that contact specified contact channels. An escalation plan uses stages that contact specified contacts.
	Plan interface{} `field:"optional" json:"plan" yaml:"plan"`
}

