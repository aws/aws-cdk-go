package awsguardduty


// Properties for defining a `CfnThreatEntitySet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnThreatEntitySetProps := &CfnThreatEntitySetProps{
//   	Format: jsii.String("format"),
//   	Location: jsii.String("location"),
//
//   	// the properties below are optional
//   	Activate: jsii.Boolean(false),
//   	DetectorId: jsii.String("detectorId"),
//   	ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	Name: jsii.String("name"),
//   	Tags: []tagItemProperty{
//   		&tagItemProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatentityset.html
//
type CfnThreatEntitySetProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatentityset.html#cfn-guardduty-threatentityset-format
	//
	Format *string `field:"required" json:"format" yaml:"format"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatentityset.html#cfn-guardduty-threatentityset-location
	//
	Location *string `field:"required" json:"location" yaml:"location"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatentityset.html#cfn-guardduty-threatentityset-activate
	//
	Activate interface{} `field:"optional" json:"activate" yaml:"activate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatentityset.html#cfn-guardduty-threatentityset-detectorid
	//
	DetectorId *string `field:"optional" json:"detectorId" yaml:"detectorId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatentityset.html#cfn-guardduty-threatentityset-expectedbucketowner
	//
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatentityset.html#cfn-guardduty-threatentityset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatentityset.html#cfn-guardduty-threatentityset-tags
	//
	Tags *[]*CfnThreatEntitySet_TagItemProperty `field:"optional" json:"tags" yaml:"tags"`
}

