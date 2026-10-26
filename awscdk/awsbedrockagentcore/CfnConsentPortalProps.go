package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConsentPortal`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConsentPortalProps := &CfnConsentPortalProps{
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	IdpConfig: &IdpConfigProperty{
//   		CredentialProviderArn: jsii.String("credentialProviderArn"),
//   		Scopes: []*string{
//   			jsii.String("scopes"),
//   		},
//
//   		// the properties below are optional
//   		Audience: jsii.String("audience"),
//   	},
//   	Name: jsii.String("name"),
//   	Sources: []interface{}{
//   		&ConsentPortalSourceProperty{
//   			Identifier: jsii.String("identifier"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html
//
type CfnConsentPortalProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html#cfn-bedrockagentcore-consentportal-executionrolearn
	//
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html#cfn-bedrockagentcore-consentportal-idpconfig
	//
	IdpConfig interface{} `field:"required" json:"idpConfig" yaml:"idpConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html#cfn-bedrockagentcore-consentportal-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html#cfn-bedrockagentcore-consentportal-sources
	//
	Sources interface{} `field:"required" json:"sources" yaml:"sources"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html#cfn-bedrockagentcore-consentportal-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to assign to the consent portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html#cfn-bedrockagentcore-consentportal-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

