package previewawslexmixins


// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   textLogDestinationProperty := &TextLogDestinationProperty{
//   	CloudWatch: &CloudWatchLogGroupLogDestinationProperty{
//   		CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		LogPrefix: jsii.String("logPrefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-textlogdestination.html
//
type CfnBotAliasPropsMixin_TextLogDestinationProperty struct {
	// Defines the Amazon CloudWatch Logs log group where text and metadata logs are delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-textlogdestination.html#cfn-lex-botalias-textlogdestination-cloudwatch
	//
	CloudWatch interface{} `field:"optional" json:"cloudWatch" yaml:"cloudWatch"`
}

