package mixinsawsses


// An object that contains information about the suppression list preferences for your account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   suppressionOptionsProperty := &SuppressionOptionsProperty{
//   	SuppressedReasons: []*string{
//   		jsii.String("suppressedReasons"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-suppressionoptions.html
//
type CfnConfigurationSetPropsMixin_SuppressionOptionsProperty struct {
	// A list that contains the reasons that email addresses are automatically added to the suppression list for your account.
	//
	// This list can contain any or all of the following:
	//
	// - `COMPLAINT` – Amazon SES adds an email address to the suppression list for your account when a message sent to that address results in a complaint.
	// - `BOUNCE` – Amazon SES adds an email address to the suppression list for your account when a message sent to that address results in a hard bounce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-suppressionoptions.html#cfn-ses-configurationset-suppressionoptions-suppressedreasons
	//
	SuppressedReasons *[]*string `field:"optional" json:"suppressedReasons" yaml:"suppressedReasons"`
}

