package awsgamelift


// Amazon GameLift Anywhere configuration options for your Anywhere fleets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anywhereConfigurationProperty := &AnywhereConfigurationProperty{
//   	Cost: jsii.String("cost"),
//   }
//
type CfnFleet_AnywhereConfigurationProperty struct {
	// The cost to run your fleet per hour.
	//
	// Amazon GameLift uses the provided cost of your fleet to balance usage in queues. For more information about queues, see [Setting up queues](https://docs.aws.amazon.com/gamelift/latest/developerguide/queues-intro.html) in the *Amazon GameLift Developer Guide* .
	Cost *string `field:"required" json:"cost" yaml:"cost"`
}

