// The CDK Construct Library for AWS::GameLift
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
//   fleet := gamelift.NewBuildFleet(this, jsii.String("Game server fleet"), &buildFleetProps{
//   	fleetName: jsii.String("test-fleet"),
//   	content: build,
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_C4, ec2.instanceSize_LARGE),
//   	runtimeConfiguration: &runtimeConfiguration{
//   		serverProcesses: []serverProcess{
//   			&serverProcess{
//   				launchPath: jsii.String("/local/game/GameLiftExampleServer.x86_64"),
//   			},
//   		},
//   	},
//   	locations: []location{
//   		&location{
//   			region: jsii.String("eu-west-1"),
//   			capacity: &locationCapacity{
//   				desiredCapacity: jsii.Number(5),
//   				minSize: jsii.Number(2),
//   				maxSize: jsii.Number(10),
//   			},
//   		},
//   		&location{
//   			region: jsii.String("us-east-1"),
//   			capacity: &locationCapacity{
//   				desiredCapacity: jsii.Number(5),
//   				minSize: jsii.Number(2),
//   				maxSize: jsii.Number(10),
//   			},
//   		},
//   	},
//   })
//
//   // Or through dedicated methods
//   fleet.addLocation(jsii.String("ap-southeast-1"), jsii.Number(5), jsii.Number(2), jsii.Number(10))
//
// Experimental.
type LocationCapacity struct {
	// The number of Amazon EC2 instances you want to maintain in the specified fleet location.
	//
	// This value must fall between the minimum and maximum size limits.
	// Experimental.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// The maximum number of instances that are allowed in the specified fleet location.
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of instances that are allowed in the specified fleet location.
	// Experimental.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
}

