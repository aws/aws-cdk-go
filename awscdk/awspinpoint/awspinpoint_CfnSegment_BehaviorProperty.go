package awspinpoint


// Specifies behavior-based criteria for the segment, such as how recently users have used your app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   behaviorProperty := &behaviorProperty{
//   	recency: &recencyProperty{
//   		duration: jsii.String("duration"),
//   		recencyType: jsii.String("recencyType"),
//   	},
//   }
//
type CfnSegment_BehaviorProperty struct {
	// Specifies how recently segment members were active.
	Recency interface{} `field:"optional" json:"recency" yaml:"recency"`
}

