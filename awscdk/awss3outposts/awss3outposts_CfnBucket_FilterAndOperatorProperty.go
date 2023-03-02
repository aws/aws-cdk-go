package awss3outposts


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterAndOperatorProperty := &filterAndOperatorProperty{
//   	tags: []filterTagProperty{
//   		&filterTagProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnBucket_FilterAndOperatorProperty struct {
	// `CfnBucket.FilterAndOperatorProperty.Tags`.
	Tags *[]*CfnBucket_FilterTagProperty `field:"required" json:"tags" yaml:"tags"`
	// `CfnBucket.FilterAndOperatorProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

