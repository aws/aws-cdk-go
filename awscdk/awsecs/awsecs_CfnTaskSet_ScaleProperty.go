package awsecs


// A floating-point percentage of the desired number of tasks to place and keep running in the task set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scaleProperty := &scaleProperty{
//   	unit: jsii.String("unit"),
//   	value: jsii.Number(123),
//   }
//
type CfnTaskSet_ScaleProperty struct {
	// The unit of measure for the scale value.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// The value, specified as a percent total of a service's `desiredCount` , to scale the task set.
	//
	// Accepted values are numbers between 0 and 100.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

