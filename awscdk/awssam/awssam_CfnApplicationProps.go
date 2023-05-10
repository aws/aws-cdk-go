package awssam


// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	Location: jsii.String("location"),
//
//   	// the properties below are optional
//   	NotificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TimeoutInMinutes: jsii.Number(123),
//   }
//
type CfnApplicationProps struct {
	// `AWS::Serverless::Application.Location`.
	Location interface{} `field:"required" json:"location" yaml:"location"`
	// `AWS::Serverless::Application.NotificationArns`.
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// `AWS::Serverless::Application.Parameters`.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// `AWS::Serverless::Application.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::Serverless::Application.TimeoutInMinutes`.
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

