package awscloudtrail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEventDataStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventDataStoreProps := &cfnEventDataStoreProps{
//   	advancedEventSelectors: []interface{}{
//   		&advancedEventSelectorProperty{
//   			fieldSelectors: []interface{}{
//   				&advancedFieldSelectorProperty{
//   					field: jsii.String("field"),
//
//   					// the properties below are optional
//   					endsWith: []*string{
//   						jsii.String("endsWith"),
//   					},
//   					equalTo: []*string{
//   						jsii.String("equalTo"),
//   					},
//   					notEndsWith: []*string{
//   						jsii.String("notEndsWith"),
//   					},
//   					notEquals: []*string{
//   						jsii.String("notEquals"),
//   					},
//   					notStartsWith: []*string{
//   						jsii.String("notStartsWith"),
//   					},
//   					startsWith: []*string{
//   						jsii.String("startsWith"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			name: jsii.String("name"),
//   		},
//   	},
//   	multiRegionEnabled: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	organizationEnabled: jsii.Boolean(false),
//   	retentionPeriod: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	terminationProtectionEnabled: jsii.Boolean(false),
//   }
//
type CfnEventDataStoreProps struct {
	// `AWS::CloudTrail::EventDataStore.AdvancedEventSelectors`.
	AdvancedEventSelectors interface{} `field:"optional" json:"advancedEventSelectors" yaml:"advancedEventSelectors"`
	// `AWS::CloudTrail::EventDataStore.MultiRegionEnabled`.
	MultiRegionEnabled interface{} `field:"optional" json:"multiRegionEnabled" yaml:"multiRegionEnabled"`
	// `AWS::CloudTrail::EventDataStore.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::CloudTrail::EventDataStore.OrganizationEnabled`.
	OrganizationEnabled interface{} `field:"optional" json:"organizationEnabled" yaml:"organizationEnabled"`
	// `AWS::CloudTrail::EventDataStore.RetentionPeriod`.
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// `AWS::CloudTrail::EventDataStore.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::CloudTrail::EventDataStore.TerminationProtectionEnabled`.
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
}

