package awsappmesh


// An object that represents a duration of time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   durationProperty := &durationProperty{
//   	unit: jsii.String("unit"),
//   	value: jsii.Number(123),
//   }
//
type CfnRoute_DurationProperty struct {
	// A unit of time.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// A number of time units.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

