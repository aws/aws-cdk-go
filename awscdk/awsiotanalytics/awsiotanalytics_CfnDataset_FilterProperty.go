package awsiotanalytics


// Information which is used to filter message data, to segregate it according to the time frame in which it arrives.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &filterProperty{
//   	deltaTime: &deltaTimeProperty{
//   		offsetSeconds: jsii.Number(123),
//   		timeExpression: jsii.String("timeExpression"),
//   	},
//   }
//
type CfnDataset_FilterProperty struct {
	// Used to limit data to that which has arrived since the last execution of the action.
	DeltaTime interface{} `field:"optional" json:"deltaTime" yaml:"deltaTime"`
}

