package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialWindowOptionsProperty := &GeospatialWindowOptionsProperty{
//   	Bounds: &GeospatialCoordinateBoundsProperty{
//   		East: jsii.Number(123),
//   		North: jsii.Number(123),
//   		South: jsii.Number(123),
//   		West: jsii.Number(123),
//   	},
//   	MapZoomMode: jsii.String("mapZoomMode"),
//   }
//
type CfnDashboard_GeospatialWindowOptionsProperty struct {
	// `CfnDashboard.GeospatialWindowOptionsProperty.Bounds`.
	Bounds interface{} `field:"optional" json:"bounds" yaml:"bounds"`
	// `CfnDashboard.GeospatialWindowOptionsProperty.MapZoomMode`.
	MapZoomMode *string `field:"optional" json:"mapZoomMode" yaml:"mapZoomMode"`
}

