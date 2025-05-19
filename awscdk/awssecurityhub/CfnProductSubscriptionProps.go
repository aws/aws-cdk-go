package awssecurityhub


// Properties for defining a `CfnProductSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProductSubscriptionProps := &CfnProductSubscriptionProps{
//   	ProductArn: jsii.String("productArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-productsubscription.html
//
type CfnProductSubscriptionProps struct {
	// The ARN of the product to enable the integration for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-productsubscription.html#cfn-securityhub-productsubscription-productarn
	//
	ProductArn *string `field:"required" json:"productArn" yaml:"productArn"`
}

