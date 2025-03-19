package awsses


// A CloudWatch dimension upon which to categorize your emails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchDimension := &CloudWatchDimension{
//   	DefaultValue: jsii.String("defaultValue"),
//   	Name: jsii.String("name"),
//   	Source: awscdk.Aws_ses.CloudWatchDimensionSource_EMAIL_HEADER,
//   }
//
type CloudWatchDimension struct {
	// The default value of the dimension that is published to Amazon CloudWatch if you do not provide the value of the dimension when you send an email.
	DefaultValue *string `field:"required" json:"defaultValue" yaml:"defaultValue"`
	// The name of an Amazon CloudWatch dimension associated with an email sending metric.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The place where Amazon SES finds the value of a dimension to publish to Amazon CloudWatch.
	Source CloudWatchDimensionSource `field:"required" json:"source" yaml:"source"`
}

