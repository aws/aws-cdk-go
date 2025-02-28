package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataProtectionSettings`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataProtectionSettingsProps := &CfnDataProtectionSettingsProps{
//   	AdditionalEncryptionContext: map[string]*string{
//   		"additionalEncryptionContextKey": jsii.String("additionalEncryptionContext"),
//   	},
//   	CustomerManagedKey: jsii.String("customerManagedKey"),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	InlineRedactionConfiguration: &InlineRedactionConfigurationProperty{
//   		InlineRedactionPatterns: []interface{}{
//   			&InlineRedactionPatternProperty{
//   				RedactionPlaceHolder: &RedactionPlaceHolderProperty{
//   					RedactionPlaceHolderType: jsii.String("redactionPlaceHolderType"),
//
//   					// the properties below are optional
//   					RedactionPlaceHolderText: jsii.String("redactionPlaceHolderText"),
//   				},
//
//   				// the properties below are optional
//   				BuiltInPatternId: jsii.String("builtInPatternId"),
//   				ConfidenceLevel: jsii.Number(123),
//   				CustomPattern: &CustomPatternProperty{
//   					PatternName: jsii.String("patternName"),
//   					PatternRegex: jsii.String("patternRegex"),
//
//   					// the properties below are optional
//   					KeywordRegex: jsii.String("keywordRegex"),
//   					PatternDescription: jsii.String("patternDescription"),
//   				},
//   				EnforcedUrls: []*string{
//   					jsii.String("enforcedUrls"),
//   				},
//   				ExemptUrls: []*string{
//   					jsii.String("exemptUrls"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		GlobalConfidenceLevel: jsii.Number(123),
//   		GlobalEnforcedUrls: []*string{
//   			jsii.String("globalEnforcedUrls"),
//   		},
//   		GlobalExemptUrls: []*string{
//   			jsii.String("globalExemptUrls"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-dataprotectionsettings.html
//
type CfnDataProtectionSettingsProps struct {
	// The additional encryption context of the data protection settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-dataprotectionsettings.html#cfn-workspacesweb-dataprotectionsettings-additionalencryptioncontext
	//
	AdditionalEncryptionContext interface{} `field:"optional" json:"additionalEncryptionContext" yaml:"additionalEncryptionContext"`
	// The customer managed key used to encrypt sensitive information in the data protection settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-dataprotectionsettings.html#cfn-workspacesweb-dataprotectionsettings-customermanagedkey
	//
	CustomerManagedKey *string `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// The description of the data protection settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-dataprotectionsettings.html#cfn-workspacesweb-dataprotectionsettings-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the data protection settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-dataprotectionsettings.html#cfn-workspacesweb-dataprotectionsettings-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The inline redaction configuration for the data protection settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-dataprotectionsettings.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionconfiguration
	//
	InlineRedactionConfiguration interface{} `field:"optional" json:"inlineRedactionConfiguration" yaml:"inlineRedactionConfiguration"`
	// The tags of the data protection settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-dataprotectionsettings.html#cfn-workspacesweb-dataprotectionsettings-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

