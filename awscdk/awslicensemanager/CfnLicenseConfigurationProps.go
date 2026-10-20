package awslicensemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLicenseConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLicenseConfigurationProps := &CfnLicenseConfigurationProps{
//   	LicenseCountingType: jsii.String("licenseCountingType"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DisassociateWhenNotFound: jsii.Boolean(false),
//   	LicenseCount: jsii.Number(123),
//   	LicenseCountHardLimit: jsii.Boolean(false),
//   	LicenseExpiry: jsii.Number(123),
//   	LicenseRules: []*string{
//   		jsii.String("licenseRules"),
//   	},
//   	ProductInformationList: []interface{}{
//   		&ProductInformationProperty{
//   			ProductInformationFilterList: []interface{}{
//   				&ProductInformationFilterProperty{
//   					ProductInformationFilterComparator: jsii.String("productInformationFilterComparator"),
//   					ProductInformationFilterName: jsii.String("productInformationFilterName"),
//
//   					// the properties below are optional
//   					ProductInformationFilterValue: []*string{
//   						jsii.String("productInformationFilterValue"),
//   					},
//   				},
//   			},
//   			ResourceType: jsii.String("resourceType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseconfiguration.html
//
type CfnLicenseConfigurationProps struct {
	// Dimension used to track the license inventory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseconfiguration.html#cfn-licensemanager-licenseconfiguration-licensecountingtype
	//
	LicenseCountingType *string `field:"required" json:"licenseCountingType" yaml:"licenseCountingType"`
	// Name of the license configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseconfiguration.html#cfn-licensemanager-licenseconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the license configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseconfiguration.html#cfn-licensemanager-licenseconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When true, disassociates a resource when software is uninstalled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseconfiguration.html#cfn-licensemanager-licenseconfiguration-disassociatewhennotfound
	//
	DisassociateWhenNotFound interface{} `field:"optional" json:"disassociateWhenNotFound" yaml:"disassociateWhenNotFound"`
	// Number of licenses managed by the license configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseconfiguration.html#cfn-licensemanager-licenseconfiguration-licensecount
	//
	LicenseCount *float64 `field:"optional" json:"licenseCount" yaml:"licenseCount"`
	// Indicates whether hard or soft license enforcement is used.
	//
	// Exceeding a hard limit blocks the launch of new instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseconfiguration.html#cfn-licensemanager-licenseconfiguration-licensecounthardlimit
	//
	LicenseCountHardLimit interface{} `field:"optional" json:"licenseCountHardLimit" yaml:"licenseCountHardLimit"`
	// License configuration expiry as an epoch timestamp (seconds since 1970-01-01T00:00:00Z).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseconfiguration.html#cfn-licensemanager-licenseconfiguration-licenseexpiry
	//
	LicenseExpiry *float64 `field:"optional" json:"licenseExpiry" yaml:"licenseExpiry"`
	// License rules.
	//
	// The syntax is #name=value (for example, #allowedTenancy=EC2-DedicatedHost).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseconfiguration.html#cfn-licensemanager-licenseconfiguration-licenserules
	//
	LicenseRules *[]*string `field:"optional" json:"licenseRules" yaml:"licenseRules"`
	// Product information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseconfiguration.html#cfn-licensemanager-licenseconfiguration-productinformationlist
	//
	ProductInformationList interface{} `field:"optional" json:"productInformationList" yaml:"productInformationList"`
	// Tags to add to the license configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseconfiguration.html#cfn-licensemanager-licenseconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

