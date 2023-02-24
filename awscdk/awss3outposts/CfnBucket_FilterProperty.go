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
type CfnBucket_FilterProperty struct {
	// `CfnBucket.FilterProperty.AndOperator`.
	AndOperator interface{} `field:"optional" json:"andOperator" yaml:"andOperator"`
	// `CfnBucket.FilterProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// `CfnBucket.FilterProperty.Tag`.
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

