package previewawsvpclatticemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDomainVerificationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDomainVerificationMixinProps := &CfnDomainVerificationMixinProps{
//   	DomainName: jsii.String("domainName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-domainverification.html
//
type CfnDomainVerificationMixinProps struct {
	// The domain name being verified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-domainverification.html#cfn-vpclattice-domainverification-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The tags associated with the domain verification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-domainverification.html#cfn-vpclattice-domainverification-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

