package awss3outposts


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterAndOperatorProperty := &FilterAndOperatorProperty{
//   	Tags: []filterTagProperty{
//   		&filterTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   }
//
type CfnBucket_FilterAndOperatorProperty struct {
	// `CfnBucket.FilterAndOperatorProperty.Tags`.
	Tags *[]*CfnBucket_FilterTagProperty `field:"required" json:"tags" yaml:"tags"`
	// `CfnBucket.FilterAndOperatorProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

