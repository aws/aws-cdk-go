package awscdkgameliftalpha


// A remote location where a multi-location fleet can deploy EC2 instances for game hosting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//
//   location := &Location{
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	Capacity: &LocationCapacity{
//   		DesiredCapacity: jsii.Number(123),
//   		MaxSize: jsii.Number(123),
//   		MinSize: jsii.Number(123),
//   	},
//   }
//
// Experimental.
type Location struct {
	// An AWS Region code.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Current resource capacity settings in a specified fleet or location.
	//
	// The location value might refer to a fleet's remote location or its home Region.
	// Experimental.
	Capacity *LocationCapacity `field:"optional" json:"capacity" yaml:"capacity"`
}

