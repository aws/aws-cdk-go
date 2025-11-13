package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWorkloadIdentity`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkloadIdentityProps := &CfnWorkloadIdentityProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AllowedResourceOauth2ReturnUrls: []*string{
//   		jsii.String("allowedResourceOauth2ReturnUrls"),
//   	},
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
type CfnWorkloadIdentityProps struct {
	// The name of the workload identity.
	//
	// The name must be unique within your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-workloadidentity.html#cfn-bedrockagentcore-workloadidentity-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of allowed OAuth2 return URLs for resources associated with this workload identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-workloadidentity.html#cfn-bedrockagentcore-workloadidentity-allowedresourceoauth2returnurls
	//
	AllowedResourceOauth2ReturnUrls *[]*string `field:"optional" json:"allowedResourceOauth2ReturnUrls" yaml:"allowedResourceOauth2ReturnUrls"`
	// The tags for the workload identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-workloadidentity.html#cfn-bedrockagentcore-workloadidentity-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

