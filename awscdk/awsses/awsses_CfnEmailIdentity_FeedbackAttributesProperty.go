package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   feedbackAttributesProperty := &feedbackAttributesProperty{
//   	emailForwardingEnabled: jsii.Boolean(false),
//   }
//
type CfnEmailIdentity_FeedbackAttributesProperty struct {
	// `CfnEmailIdentity.FeedbackAttributesProperty.EmailForwardingEnabled`.
	EmailForwardingEnabled interface{} `field:"optional" json:"emailForwardingEnabled" yaml:"emailForwardingEnabled"`
}

