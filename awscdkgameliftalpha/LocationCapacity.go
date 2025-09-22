package awscdkgameliftalpha


// Current resource capacity settings in a specified fleet or location.
//
// The location value might refer to a fleet's remote location or its home Region.
//
// Example:
//   var build build
//
//
//   // Locations can be added directly through constructor
//   fleet := gamelift.NewBuildFleet(this, jsii.String("Game server fleet"), &BuildFleetProps{
//   	FleetName: jsii.String("test-fleet"),
//   	Content: build,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_C4, ec2.InstanceSize_LARGE),
//   	RuntimeConfiguration: &RuntimeConfiguration{
//   		ServerProcesses: []serverProcess{
//   			&serverProcess{
//   				LaunchPath: jsii.String("/local/game/GameLiftExampleServer.x86_64"),
//   			},
//   		},
//   	},
//   	Locations: []location{
//   		&location{
//   			Region: jsii.String("eu-west-1"),
//   			Capacity: &LocationCapacity{
//   				DesiredCapacity: jsii.Number(5),
//   				MinSize: jsii.Number(2),
//   				MaxSize: jsii.Number(10),
//   			},
//   		},
//   		&location{
//   			Region: jsii.String("us-east-1"),
//   			Capacity: &LocationCapacity{
//   				DesiredCapacity: jsii.Number(5),
//   				MinSize: jsii.Number(2),
//   				MaxSize: jsii.Number(10),
//   			},
//   		},
//   	},
//   })
//
//   // Or through dedicated methods
//   fleet.AddLocation(jsii.String("ap-southeast-1"), jsii.Number(5), jsii.Number(2), jsii.Number(10))
//
// Experimental.
type LocationCapacity struct {
	// The number of Amazon EC2 instances you want to maintain in the specified fleet location.
	//
	// This value must fall between the minimum and maximum size limits.
	// Default: the default value is 0.
	//
	// Experimental.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// The maximum number of instances that are allowed in the specified fleet location.
	// Default: the default value is 1.
	//
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of instances that are allowed in the specified fleet location.
	// Default: the default value is 0.
	//
	// Experimental.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
}

