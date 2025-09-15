package awsses


// The action to add a header to a message.
//
// When executed, this action will add the given header to the message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addHeaderActionProperty := &AddHeaderActionProperty{
//   	HeaderName: jsii.String("headerName"),
//   	HeaderValue: jsii.String("headerValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-addheaderaction.html
//
type CfnMailManagerRuleSet_AddHeaderActionProperty struct {
	// The name of the header to add to an email.
	//
	// The header must be prefixed with "X-". Headers are added regardless of whether the header name pre-existed in the email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-addheaderaction.html#cfn-ses-mailmanagerruleset-addheaderaction-headername
	//
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The value of the header to add to the email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-addheaderaction.html#cfn-ses-mailmanagerruleset-addheaderaction-headervalue
	//
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
}

