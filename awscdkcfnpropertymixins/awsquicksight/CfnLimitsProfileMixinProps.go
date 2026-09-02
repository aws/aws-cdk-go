package awsquicksight


// Properties for CfnLimitsProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnLimitsProfileMixinProps := &CfnLimitsProfileMixinProps{
//   	AccountId: jsii.String("accountId"),
//   	Description: jsii.String("description"),
//   	ProfileName: jsii.String("profileName"),
//   	ResourceLimits: map[string]interface{}{
//   		"resourceLimitsKey": &ProfileLimitValueProperty{
//   			"maxValue": jsii.Number(123),
//   			"unit": jsii.String("unit"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-limitsprofile.html
//
type CfnLimitsProfileMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-limitsprofile.html#cfn-quicksight-limitsprofile-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-limitsprofile.html#cfn-quicksight-limitsprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-limitsprofile.html#cfn-quicksight-limitsprofile-profilename
	//
	ProfileName *string `field:"optional" json:"profileName" yaml:"profileName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-limitsprofile.html#cfn-quicksight-limitsprofile-resourcelimits
	//
	ResourceLimits interface{} `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
}

