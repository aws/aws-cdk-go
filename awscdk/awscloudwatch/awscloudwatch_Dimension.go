package awscloudwatch


// Metric dimension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   dimension := &dimension{
//   	name: jsii.String("name"),
//   	value: value,
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-dimension.html
//
type Dimension struct {
	// Name of the dimension.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the dimension.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

