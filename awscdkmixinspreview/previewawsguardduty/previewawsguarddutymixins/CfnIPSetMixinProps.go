package previewawsguarddutymixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIPSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIPSetMixinProps := &CfnIPSetMixinProps{
//   	Activate: jsii.Boolean(false),
//   	DetectorId: jsii.String("detectorId"),
//   	ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	Format: jsii.String("format"),
//   	Location: jsii.String("location"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html
//
type CfnIPSetMixinProps struct {
	// A boolean value that determines if GuardDuty can start using this list for custom threat detection.
	//
	// For GuardDuty to prevent generating findings based on an activity associated with these entries, this list must be active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-activate
	//
	Activate interface{} `field:"optional" json:"activate" yaml:"activate"`
	// The unique ID of the detector of the GuardDuty account for which you want to create an IPSet.
	//
	// To find the `detectorId` in the current Region, see the
	// Settings page in the GuardDuty console, or run the [ListDetectors](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListDetectors.html) API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-detectorid
	//
	DetectorId *string `field:"optional" json:"detectorId" yaml:"detectorId"`
	// The AWS account ID that owns the Amazon S3 bucket specified in the *Location* field.
	//
	// When you provide this account ID, GuardDuty will validate that the S3 bucket belongs to this account. If you don't specify an account ID owner, GuardDuty doesn't perform any validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-expectedbucketowner
	//
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// The format of the file that contains the IPSet.
	//
	// For information about supported formats, see [List formats](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_upload-lists.html#prepare_list) in the *Amazon GuardDuty User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The URI of the file that contains the IPSet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-location
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The user-friendly name to identify the IPSet.
	//
	// The name of your list must be unique within an AWS account and Region. Valid characters are alphanumeric, whitespace, dash (-), and underscores (_).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags to be added to a new threat entity set resource.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

