package previewawsdatazonemixins


// The configuration details of the Amazon SageMaker data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerRunConfigurationInputProperty := &SageMakerRunConfigurationInputProperty{
//   	TrackingAssets: map[string][]*string{
//   		"trackingAssetsKey": []*string{
//   			jsii.String("trackingAssets"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-sagemakerrunconfigurationinput.html
//
type CfnDataSourcePropsMixin_SageMakerRunConfigurationInputProperty struct {
	// The tracking assets of the Amazon SageMaker run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-sagemakerrunconfigurationinput.html#cfn-datazone-datasource-sagemakerrunconfigurationinput-trackingassets
	//
	TrackingAssets interface{} `field:"optional" json:"trackingAssets" yaml:"trackingAssets"`
}

