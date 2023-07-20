package awss3outposts


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &FilterProperty{
//   	AndOperator: &FilterAndOperatorProperty{
//   		Tags: []filterTagProperty{
//   			&filterTagProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Prefix: jsii.String("prefix"),
//   	},
//   	Prefix: jsii.String("prefix"),
//   	Tag: &filterTagProperty{
//   		Key: jsii.String("key"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-filter.html
//
type CfnBucket_FilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-filter.html#cfn-s3outposts-bucket-filter-andoperator
	//
	AndOperator interface{} `field:"optional" json:"andOperator" yaml:"andOperator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-filter.html#cfn-s3outposts-bucket-filter-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-filter.html#cfn-s3outposts-bucket-filter-tag
	//
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

