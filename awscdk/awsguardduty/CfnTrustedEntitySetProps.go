package awsguardduty


// Properties for defining a `CfnTrustedEntitySet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrustedEntitySetProps := &CfnTrustedEntitySetProps{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html
//
type CfnTrustedEntitySetProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-format
	//
	Format *string `field:"required" json:"format" yaml:"format"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-location
	//
	Location *string `field:"required" json:"location" yaml:"location"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-activate
	//
	Activate interface{} `field:"optional" json:"activate" yaml:"activate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-detectorid
	//
	DetectorId *string `field:"optional" json:"detectorId" yaml:"detectorId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-expectedbucketowner
	//
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-trustedentityset.html#cfn-guardduty-trustedentityset-tags
	//
	Tags *[]*CfnTrustedEntitySet_TagItemProperty `field:"optional" json:"tags" yaml:"tags"`
}

