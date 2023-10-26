package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainProps := &CfnDomainProps{
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	DeadLetterQueueUrl: jsii.String("deadLetterQueueUrl"),
//   	DefaultEncryptionKey: jsii.String("defaultEncryptionKey"),
//   	DefaultExpirationDays: jsii.Number(123),
//   	Matching: &MatchingProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		AutoMerging: &AutoMergingProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			ConflictResolution: &ConflictResolutionProperty{
//   				ConflictResolvingModel: jsii.String("conflictResolvingModel"),
//
//   				// the properties below are optional
//   				SourceName: jsii.String("sourceName"),
//   			},
//   			Consolidation: &ConsolidationProperty{
//   				MatchingAttributesList: []interface{}{
//   					[]*string{
//   						jsii.String("matchingAttributesList"),
//   					},
//   				},
//   			},
//   			MinAllowedConfidenceScoreForMerging: jsii.Number(123),
//   		},
//   		ExportingConfig: &ExportingConfigProperty{
//   			S3Exporting: &S3ExportingConfigProperty{
//   				S3BucketName: jsii.String("s3BucketName"),
//
//   				// the properties below are optional
//   				S3KeyName: jsii.String("s3KeyName"),
//   			},
//   		},
//   		JobSchedule: &JobScheduleProperty{
//   			DayOfTheWeek: jsii.String("dayOfTheWeek"),
//   			Time: jsii.String("time"),
//   		},
//   	},
//   	RuleBasedMatching: &RuleBasedMatchingProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		AttributeTypesSelector: &AttributeTypesSelectorProperty{
//   			AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//
//   			// the properties below are optional
//   			Address: []*string{
//   				jsii.String("address"),
//   			},
//   			EmailAddress: []*string{
//   				jsii.String("emailAddress"),
//   			},
//   			PhoneNumber: []*string{
//   				jsii.String("phoneNumber"),
//   			},
//   		},
//   		ConflictResolution: &ConflictResolutionProperty{
//   			ConflictResolvingModel: jsii.String("conflictResolvingModel"),
//
//   			// the properties below are optional
//   			SourceName: jsii.String("sourceName"),
//   		},
//   		ExportingConfig: &ExportingConfigProperty{
//   			S3Exporting: &S3ExportingConfigProperty{
//   				S3BucketName: jsii.String("s3BucketName"),
//
//   				// the properties below are optional
//   				S3KeyName: jsii.String("s3KeyName"),
//   			},
//   		},
//   		MatchingRules: []interface{}{
//   			&MatchingRuleProperty{
//   				Rule: []*string{
//   					jsii.String("rule"),
//   				},
//   			},
//   		},
//   		MaxAllowedRuleLevelForMatching: jsii.Number(123),
//   		MaxAllowedRuleLevelForMerging: jsii.Number(123),
//   		Status: jsii.String("status"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html
//
type CfnDomainProps struct {
	// The unique name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The URL of the SQS dead letter queue, which is used for reporting errors associated with ingesting data from third party applications.
	//
	// You must set up a policy on the DeadLetterQueue for the SendMessage operation to enable Amazon Connect Customer Profiles to send messages to the DeadLetterQueue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-deadletterqueueurl
	//
	DeadLetterQueueUrl *string `field:"optional" json:"deadLetterQueueUrl" yaml:"deadLetterQueueUrl"`
	// The default encryption key, which is an AWS managed key, is used when no specific type of encryption key is specified.
	//
	// It is used to encrypt all data before it is placed in permanent or semi-permanent storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-defaultencryptionkey
	//
	DefaultEncryptionKey *string `field:"optional" json:"defaultEncryptionKey" yaml:"defaultEncryptionKey"`
	// The default number of days until the data within the domain expires.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-defaultexpirationdays
	//
	DefaultExpirationDays *float64 `field:"optional" json:"defaultExpirationDays" yaml:"defaultExpirationDays"`
	// The process of matching duplicate profiles.
	//
	// If Matching = true, Amazon Connect Customer Profiles starts a weekly batch process called Identity Resolution Job. If you do not specify a date and time for Identity Resolution Job to run, by default it runs every Saturday at 12AM UTC to detect duplicate profiles in your domains. After the Identity Resolution Job completes, use the GetMatches API to return and review the results. Or, if you have configured ExportingConfig in the MatchingRequest, you can download the results from S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-matching
	//
	Matching interface{} `field:"optional" json:"matching" yaml:"matching"`
	// The process of matching duplicate profiles using the Rule-Based matching.
	//
	// If RuleBasedMatching = true, Amazon Connect Customer Profiles will start to match and merge your profiles according to your configuration in the RuleBasedMatchingRequest. You can use the ListRuleBasedMatches and GetSimilarProfiles API to return and review the results. Also, if you have configured ExportingConfig in the RuleBasedMatchingRequest, you can download the results from S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-rulebasedmatching
	//
	RuleBasedMatching interface{} `field:"optional" json:"ruleBasedMatching" yaml:"ruleBasedMatching"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

