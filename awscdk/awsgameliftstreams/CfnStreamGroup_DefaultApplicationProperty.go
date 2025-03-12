package awsgameliftstreams


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultApplicationProperty := &DefaultApplicationProperty{
//   	Arn: jsii.String("arn"),
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-defaultapplication.html
//
type CfnStreamGroup_DefaultApplicationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-defaultapplication.html#cfn-gameliftstreams-streamgroup-defaultapplication-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-defaultapplication.html#cfn-gameliftstreams-streamgroup-defaultapplication-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

