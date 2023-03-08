package awss3outposts


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &filterProperty{
//   	andOperator: &filterAndOperatorProperty{
//   		tags: []filterTagProperty{
//   			&filterTagProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		prefix: jsii.String("prefix"),
//   	},
//   	prefix: jsii.String("prefix"),
//   	tag: &filterTagProperty{
//   		key: jsii.String("key"),
//   		value: jsii.String("value"),
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

