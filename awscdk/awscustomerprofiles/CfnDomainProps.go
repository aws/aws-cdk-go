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
//   	DefaultExpirationDays: jsii.Number(123),
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	DataStore: &DataStoreProperty{
//   		Enabled: jsii.Boolean(false),
//   		Readiness: &ReadinessProperty{
//   			Message: jsii.String("message"),
//   			ProgressPercentage: jsii.Number(123),
//   		},
//   	},
//   	DeadLetterQueueUrl: jsii.String("deadLetterQueueUrl"),
//   	DefaultEncryptionKey: jsii.String("defaultEncryptionKey"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html
//
type CfnDomainProps struct {
	// The default number of days until the data within the domain expires.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-defaultexpirationdays
	//
	DefaultExpirationDays *float64 `field:"required" json:"defaultExpirationDays" yaml:"defaultExpirationDays"`
	// The unique name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Configuration and status of the data store for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-datastore
	//
	DataStore interface{} `field:"optional" json:"dataStore" yaml:"dataStore"`
	// The URL of the SQS dead letter queue, which is used for reporting errors associated with ingesting data from third party applications.
	//
	// You must set up a policy on the `DeadLetterQueue` for the `SendMessage` operation to enable Amazon Connect Customer Profiles to send messages to the `DeadLetterQueue` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-deadletterqueueurl
	//
	DeadLetterQueueUrl *string `field:"optional" json:"deadLetterQueueUrl" yaml:"deadLetterQueueUrl"`
	// The default encryption key, which is an AWS managed key, is used when no specific type of encryption key is specified.
	//
	// It is used to encrypt all data before it is placed in permanent or semi-permanent storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-defaultencryptionkey
	//
	DefaultEncryptionKey *string `field:"optional" json:"defaultEncryptionKey" yaml:"defaultEncryptionKey"`
	// The process of matching duplicate profiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-matching
	//
	Matching interface{} `field:"optional" json:"matching" yaml:"matching"`
	// The process of matching duplicate profiles using Rule-Based matching.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-rulebasedmatching
	//
	RuleBasedMatching interface{} `field:"optional" json:"ruleBasedMatching" yaml:"ruleBasedMatching"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

