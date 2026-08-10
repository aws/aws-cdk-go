package awssecurityhub


// Represents a specific health issue detected for a connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthIssueProperty := &HealthIssueProperty{
//   	Code: jsii.String("code"),
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connector-healthissue.html
//
type CfnConnector_HealthIssueProperty struct {
	// The code identifying the type of health issue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connector-healthissue.html#cfn-securityhub-connector-healthissue-code
	//
	Code *string `field:"required" json:"code" yaml:"code"`
	// A human-readable message that describes the health issue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connector-healthissue.html#cfn-securityhub-connector-healthissue-message
	//
	Message *string `field:"required" json:"message" yaml:"message"`
}

