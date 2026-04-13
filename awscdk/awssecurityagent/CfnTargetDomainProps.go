package awssecurityagent

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTargetDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTargetDomainProps := &CfnTargetDomainProps{
//   	TargetDomainName: jsii.String("targetDomainName"),
//   	VerificationMethod: jsii.String("verificationMethod"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-targetdomain.html
//
type CfnTargetDomainProps struct {
	// Domain name of the target domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-targetdomain.html#cfn-securityagent-targetdomain-targetdomainname
	//
	TargetDomainName *string `field:"required" json:"targetDomainName" yaml:"targetDomainName"`
	// Verification method for the target domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-targetdomain.html#cfn-securityagent-targetdomain-verificationmethod
	//
	VerificationMethod *string `field:"required" json:"verificationMethod" yaml:"verificationMethod"`
	// Tags for the target domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-targetdomain.html#cfn-securityagent-targetdomain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

