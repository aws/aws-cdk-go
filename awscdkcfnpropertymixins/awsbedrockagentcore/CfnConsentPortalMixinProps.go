package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConsentPortalPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnConsentPortalMixinProps := &CfnConsentPortalMixinProps{
//   	Description: jsii.String("description"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	IdpConfig: &IdpConfigProperty{
//   		Audience: jsii.String("audience"),
//   		CredentialProviderArn: jsii.String("credentialProviderArn"),
//   		Scopes: []*string{
//   			jsii.String("scopes"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Sources: []interface{}{
//   		&ConsentPortalSourceProperty{
//   			Identifier: jsii.String("identifier"),
//   			Type: jsii.String("type"),
//   		},
//   	},
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
type CfnConsentPortalMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html#cfn-bedrockagentcore-consentportal-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html#cfn-bedrockagentcore-consentportal-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html#cfn-bedrockagentcore-consentportal-idpconfig
	//
	IdpConfig interface{} `field:"optional" json:"idpConfig" yaml:"idpConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html#cfn-bedrockagentcore-consentportal-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html#cfn-bedrockagentcore-consentportal-sources
	//
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// Tags to assign to the consent portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-consentportal.html#cfn-bedrockagentcore-consentportal-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

