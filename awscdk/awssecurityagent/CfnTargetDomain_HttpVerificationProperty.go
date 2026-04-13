package awssecurityagent


// Represents HTTP route verification details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpVerificationProperty := &HttpVerificationProperty{
//   	RoutePath: jsii.String("routePath"),
//   	Token: jsii.String("token"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-targetdomain-httpverification.html
//
type CfnTargetDomain_HttpVerificationProperty struct {
	// Route path where verification token should be placed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-targetdomain-httpverification.html#cfn-securityagent-targetdomain-httpverification-routepath
	//
	RoutePath *string `field:"optional" json:"routePath" yaml:"routePath"`
	// Token used to verify domain ownership.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-targetdomain-httpverification.html#cfn-securityagent-targetdomain-httpverification-token
	//
	Token *string `field:"optional" json:"token" yaml:"token"`
}

