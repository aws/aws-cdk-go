package awsses


// When included in a receipt rule, this action parses the received message and starts an email contact in Amazon Connect on your behalf.
//
// > When you receive emails, the maximum email size (including headers) is 40 MB. Additionally, emails may only have up to 10 attachments. Emails larger than 40 MB or with more than 10 attachments will be bounced.
//
// We recommend that you configure this action via Amazon Connect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectActionProperty := &ConnectActionProperty{
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	InstanceArn: jsii.String("instanceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-connectaction.html
//
type CfnReceiptRule_ConnectActionProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role to be used by Amazon Simple Email Service while starting email contacts to the Amazon Connect instance.
	//
	// This role should have permission to invoke `connect:StartEmailContact` for the given Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-connectaction.html#cfn-ses-receiptrule-connectaction-iamrolearn
	//
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The Amazon Resource Name (ARN) for the Amazon Connect instance that Amazon SES integrates with for starting email contacts.
	//
	// For more information about Amazon Connect instances, see the [Amazon Connect Administrator Guide](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-instances.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-connectaction.html#cfn-ses-receiptrule-connectaction-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
}

