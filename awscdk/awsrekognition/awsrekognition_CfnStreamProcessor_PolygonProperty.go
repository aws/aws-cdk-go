package awsrekognition


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   polygonProperty := &polygonProperty{
//   	polygon: []interface{}{
//   		&pointProperty{
//   			x: jsii.Number(123),
//   			y: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnStreamProcessor_PolygonProperty struct {
	// `CfnStreamProcessor.PolygonProperty.Polygon`.
	Polygon interface{} `field:"required" json:"polygon" yaml:"polygon"`
}

