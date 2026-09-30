package awslicensemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLicenseAssetGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnLicenseAssetGroupMixinProps := &CfnLicenseAssetGroupMixinProps{
//   	AssociatedLicenseAssetRulesetArns: []*string{
//   		jsii.String("associatedLicenseAssetRulesetArns"),
//   	},
//   	Description: jsii.String("description"),
//   	LicenseAssetGroupConfigurations: []interface{}{
//   		&LicenseAssetGroupConfigurationProperty{
//   			UsageDimension: jsii.String("usageDimension"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Properties: []interface{}{
//   		&LicenseAssetGroupPropertyProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetgroup.html
//
type CfnLicenseAssetGroupMixinProps struct {
	// ARNs of associated license asset rulesets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetgroup.html#cfn-licensemanager-licenseassetgroup-associatedlicenseassetrulesetarns
	//
	AssociatedLicenseAssetRulesetArns *[]*string `field:"optional" json:"associatedLicenseAssetRulesetArns" yaml:"associatedLicenseAssetRulesetArns"`
	// License asset group description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetgroup.html#cfn-licensemanager-licenseassetgroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// License asset group configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetgroup.html#cfn-licensemanager-licenseassetgroup-licenseassetgroupconfigurations
	//
	LicenseAssetGroupConfigurations interface{} `field:"optional" json:"licenseAssetGroupConfigurations" yaml:"licenseAssetGroupConfigurations"`
	// License asset group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetgroup.html#cfn-licensemanager-licenseassetgroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// License asset group properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetgroup.html#cfn-licensemanager-licenseassetgroup-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// Tags to add to the license asset group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetgroup.html#cfn-licensemanager-licenseassetgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

