package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   provisionedThroughputProperty := &ProvisionedThroughputProperty{
//   	ReadCapacityUnits: jsii.Number(123),
//   	WriteCapacityUnits: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-simpletable-provisionedthroughput.html
//
type CfnSimpleTablePropsMixin_ProvisionedThroughputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-simpletable-provisionedthroughput.html#cfn-serverless-simpletable-provisionedthroughput-readcapacityunits
	//
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-simpletable-provisionedthroughput.html#cfn-serverless-simpletable-provisionedthroughput-writecapacityunits
	//
	WriteCapacityUnits *float64 `field:"optional" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
}

