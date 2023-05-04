package awsquicksight


// Determines the color that is applied to a particular data value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataColorProperty := &DataColorProperty{
//   	Color: jsii.String("color"),
//   	DataValue: jsii.Number(123),
//   }
//
type CfnTemplate_DataColorProperty struct {
	// The color that is applied to the data value.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The data value that the color is applied to.
	DataValue *float64 `field:"optional" json:"dataValue" yaml:"dataValue"`
}

