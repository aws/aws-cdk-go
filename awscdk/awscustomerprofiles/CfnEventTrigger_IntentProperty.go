package awscustomerprofiles


// Identifies a predicted intent by domain and type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intentProperty := &IntentProperty{
//   	IntentDomain: jsii.String("intentDomain"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intent.html
//
type CfnEventTrigger_IntentProperty struct {
	// The domain of the intent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intent.html#cfn-customerprofiles-eventtrigger-intent-intentdomain
	//
	IntentDomain *string `field:"required" json:"intentDomain" yaml:"intentDomain"`
	// The type of the intent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intent.html#cfn-customerprofiles-eventtrigger-intent-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

