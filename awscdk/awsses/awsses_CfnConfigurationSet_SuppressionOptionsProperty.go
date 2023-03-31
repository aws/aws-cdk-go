package awsses


// An object that contains information about the suppression list preferences for your account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   suppressionOptionsProperty := &suppressionOptionsProperty{
//   	suppressedReasons: []*string{
//   		jsii.String("suppressedReasons"),
//   	},
//   }
//
type CfnConfigurationSet_SuppressionOptionsProperty struct {
	// A list that contains the reasons that email addresses are automatically added to the suppression list for your account.
	//
	// This list can contain any or all of the following:
	//
	// - `COMPLAINT` – Amazon SES adds an email address to the suppression list for your account when a message sent to that address results in a complaint.
	// - `BOUNCE` – Amazon SES adds an email address to the suppression list for your account when a message sent to that address results in a hard bounce.
	SuppressedReasons *[]*string `field:"optional" json:"suppressedReasons" yaml:"suppressedReasons"`
}

