package awsroute53


// A complex type that lists the coordinates for a geoproximity resource record.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coordinatesProperty := &CoordinatesProperty{
//   	Latitude: jsii.String("latitude"),
//   	Longitude: jsii.String("longitude"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-coordinates.html
//
type CfnRecordSet_CoordinatesProperty struct {
	// Specifies a coordinate of the north–south position of a geographic point on the surface of the Earth (-90 - 90).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-coordinates.html#cfn-route53-recordset-coordinates-latitude
	//
	Latitude *string `field:"required" json:"latitude" yaml:"latitude"`
	// Specifies a coordinate of the east–west position of a geographic point on the surface of the Earth (-180 - 180).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-coordinates.html#cfn-route53-recordset-coordinates-longitude
	//
	Longitude *string `field:"required" json:"longitude" yaml:"longitude"`
}

