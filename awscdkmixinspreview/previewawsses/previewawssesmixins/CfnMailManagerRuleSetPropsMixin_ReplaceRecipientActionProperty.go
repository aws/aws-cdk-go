package previewawssesmixins


// This action replaces the email envelope recipients with the given list of recipients.
//
// If the condition of this action applies only to a subset of recipients, only those recipients are replaced with the recipients specified in the action. The message contents and headers are unaffected by this action, only the envelope recipients are updated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   replaceRecipientActionProperty := &ReplaceRecipientActionProperty{
//   	ReplaceWith: []*string{
//   		jsii.String("replaceWith"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-replacerecipientaction.html
//
type CfnMailManagerRuleSetPropsMixin_ReplaceRecipientActionProperty struct {
	// This action specifies the replacement recipient email addresses to insert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-replacerecipientaction.html#cfn-ses-mailmanagerruleset-replacerecipientaction-replacewith
	//
	ReplaceWith *[]*string `field:"optional" json:"replaceWith" yaml:"replaceWith"`
}

