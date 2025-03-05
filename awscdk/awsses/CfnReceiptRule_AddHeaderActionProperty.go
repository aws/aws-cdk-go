package awsses


// When included in a receipt rule, this action adds a header to the received email.
//
// For information about adding a header using a receipt rule, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html) .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-addheaderaction.html
//
type CfnReceiptRule_AddHeaderActionProperty struct {
	// The name of the header to add to the incoming message.
	//
	// The name must contain at least one character, and can contain up to 50 characters. It consists of alphanumeric ( `a–z, A–Z, 0–9` ) characters and dashes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-addheaderaction.html#cfn-ses-receiptrule-addheaderaction-headername
	//
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The content to include in the header.
	//
	// This value can contain up to 2048 characters. It can't contain newline ( `\n` ) or carriage return ( `\r` ) characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-addheaderaction.html#cfn-ses-receiptrule-addheaderaction-headervalue
	//
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
}

