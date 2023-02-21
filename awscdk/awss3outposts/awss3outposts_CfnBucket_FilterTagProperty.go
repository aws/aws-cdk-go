package awss3outposts


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterTagProperty := &filterTagProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnBucket_FilterTagProperty struct {
	// `CfnBucket.FilterTagProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnBucket.FilterTagProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

