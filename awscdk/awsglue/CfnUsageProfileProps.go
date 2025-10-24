package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnUsageProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUsageProfileProps := &CfnUsageProfileProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Configuration: &ProfileConfigurationProperty{
//   		JobConfiguration: map[string]interface{}{
//   			"jobConfigurationKey": &ConfigurationObjectProperty{
//   				"allowedValues": []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				"defaultValue": jsii.String("defaultValue"),
//   				"maxValue": jsii.String("maxValue"),
//   				"minValue": jsii.String("minValue"),
//   			},
//   		},
//   		SessionConfiguration: map[string]interface{}{
//   			"sessionConfigurationKey": &ConfigurationObjectProperty{
//   				"allowedValues": []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				"defaultValue": jsii.String("defaultValue"),
//   				"maxValue": jsii.String("maxValue"),
//   				"minValue": jsii.String("minValue"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-usageprofile.html
//
type CfnUsageProfileProps struct {
	// The name of the usage profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-usageprofile.html#cfn-glue-usageprofile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-usageprofile.html#cfn-glue-usageprofile-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// A description of the usage profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-usageprofile.html#cfn-glue-usageprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags to be applied to this UsageProfiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-usageprofile.html#cfn-glue-usageprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

