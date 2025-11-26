package previewawsgameliftstreamsmixins


// Represents the default Amazon GameLift Streams application that a stream group hosts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   defaultApplicationProperty := &DefaultApplicationProperty{
//   	Arn: jsii.String("arn"),
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-defaultapplication.html
//
type CfnStreamGroupPropsMixin_DefaultApplicationProperty struct {
	// An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely identifies the application resource. Example ARN: `arn:aws:gameliftstreams:us-west-2:111122223333:application/a-9ZY8X7Wv6` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-defaultapplication.html#cfn-gameliftstreams-streamgroup-defaultapplication-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// An ID that uniquely identifies the application resource.
	//
	// Example ID: `a-9ZY8X7Wv6` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-defaultapplication.html#cfn-gameliftstreams-streamgroup-defaultapplication-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

