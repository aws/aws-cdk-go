package previewawsiotanalyticsmixins


// Information which is used to filter message data, to segregate it according to the time frame in which it arrives.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterProperty := &FilterProperty{
//   	DeltaTime: &DeltaTimeProperty{
//   		OffsetSeconds: jsii.Number(123),
//   		TimeExpression: jsii.String("timeExpression"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-filter.html
//
type CfnDatasetPropsMixin_FilterProperty struct {
	// Used to limit data to that which has arrived since the last execution of the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-filter.html#cfn-iotanalytics-dataset-filter-deltatime
	//
	DeltaTime interface{} `field:"optional" json:"deltaTime" yaml:"deltaTime"`
}

