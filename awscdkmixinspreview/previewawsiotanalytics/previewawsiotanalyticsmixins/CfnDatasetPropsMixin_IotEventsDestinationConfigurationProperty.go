package previewawsiotanalyticsmixins


// Configuration information for delivery of dataset contents to AWS IoT Events .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iotEventsDestinationConfigurationProperty := &IotEventsDestinationConfigurationProperty{
//   	InputName: jsii.String("inputName"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-ioteventsdestinationconfiguration.html
//
type CfnDatasetPropsMixin_IotEventsDestinationConfigurationProperty struct {
	// The name of the AWS IoT Events input to which dataset contents are delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-ioteventsdestinationconfiguration.html#cfn-iotanalytics-dataset-ioteventsdestinationconfiguration-inputname
	//
	InputName *string `field:"optional" json:"inputName" yaml:"inputName"`
	// The ARN of the role that grants ITA permission to deliver dataset contents to an AWS IoT Events input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-ioteventsdestinationconfiguration.html#cfn-iotanalytics-dataset-ioteventsdestinationconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

