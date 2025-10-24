package awss3outposts


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterAndOperatorProperty := &FilterAndOperatorProperty{
//   	Tags: []FilterTagProperty{
//   		&FilterTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-filterandoperator.html
//
type CfnBucket_FilterAndOperatorProperty struct {
	// All of these tags must exist in the object's tag set in order for the rule to apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-filterandoperator.html#cfn-s3outposts-bucket-filterandoperator-tags
	//
	Tags *[]*CfnBucket_FilterTagProperty `field:"required" json:"tags" yaml:"tags"`
	// Prefix identifies one or more objects to which the rule applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-filterandoperator.html#cfn-s3outposts-bucket-filterandoperator-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

