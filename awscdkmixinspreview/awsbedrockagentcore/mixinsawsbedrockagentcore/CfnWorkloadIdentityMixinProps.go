package mixinsawsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnWorkloadIdentityPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkloadIdentityMixinProps := &CfnWorkloadIdentityMixinProps{
//   	AllowedResourceOauth2ReturnUrls: []*string{
//   		jsii.String("allowedResourceOauth2ReturnUrls"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-workloadidentity.html
//
type CfnWorkloadIdentityMixinProps struct {
	// The list of allowed OAuth2 return URLs for resources associated with this workload identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-workloadidentity.html#cfn-bedrockagentcore-workloadidentity-allowedresourceoauth2returnurls
	//
	AllowedResourceOauth2ReturnUrls *[]*string `field:"optional" json:"allowedResourceOauth2ReturnUrls" yaml:"allowedResourceOauth2ReturnUrls"`
	// The name of the workload identity.
	//
	// The name must be unique within your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-workloadidentity.html#cfn-bedrockagentcore-workloadidentity-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags for the workload identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-workloadidentity.html#cfn-bedrockagentcore-workloadidentity-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

