package mixinsawsguardduty


// Properties for CfnTrustedEntitySetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTrustedEntitySetMixinProps := &CfnTrustedEntitySetMixinProps{
//   	Activate: jsii.Boolean(false),
//   	DetectorId: jsii.String("detectorId"),
//   	ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	Format: jsii.String("format"),
//   	Location: jsii.String("location"),
//   	Name: jsii.String("name"),
//   	Tags: []TagItemProperty{
//   		&TagItemProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html
//
type CfnTrustedEntitySetMixinProps struct {
	// A boolean value that determines if GuardDuty can start using this list for custom threat detection.
	//
	// For GuardDuty to prevent generating findings based on an activity associated with these entries, this list must be active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-activate
	//
	Activate interface{} `field:"optional" json:"activate" yaml:"activate"`
	// The unique regional detector ID of the GuardDuty account for which you want to create a trusted entity set.
	//
	// To find the `detectorId` in the current Region, see the Settings page in the GuardDuty console, or run the [ListDetectors](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListDetectors.html) API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-detectorid
	//
	DetectorId *string `field:"optional" json:"detectorId" yaml:"detectorId"`
	// The AWS account ID that owns the Amazon S3 bucket specified in the *Location* field.
	//
	// Whether or not you provide the account ID for this optional field, GuardDuty validates that the account ID associated with the `DetectorId` value owns the S3 bucket in the `Location` field. If GuardDuty finds that this S3 bucket doesn't belong to the specified account ID, you will get an error at the time of activating this list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-expectedbucketowner
	//
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// The format of the file that contains the trusted entity set.
	//
	// For information about supported formats, see [List formats](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_upload-lists.html#prepare_list) in the *Amazon GuardDuty User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The URI of the file that contains the trusted entity set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-location
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
	// A user-friendly name to identify the trusted entity set.
	//
	// Valid characters include lowercase letters, uppercase letters, numbers, dash(-), and underscore (_).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags to be added to a new trusted entity set resource.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-tags
	//
	Tags *[]*CfnTrustedEntitySetPropsMixin_TagItemProperty `field:"optional" json:"tags" yaml:"tags"`
}

