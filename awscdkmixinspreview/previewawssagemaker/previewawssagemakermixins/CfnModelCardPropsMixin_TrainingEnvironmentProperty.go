package previewawssagemakermixins


// SageMaker AI training image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   trainingEnvironmentProperty := &TrainingEnvironmentProperty{
//   	ContainerImage: []*string{
//   		jsii.String("containerImage"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingenvironment.html
//
type CfnModelCardPropsMixin_TrainingEnvironmentProperty struct {
	// SageMaker AI inference image URI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingenvironment.html#cfn-sagemaker-modelcard-trainingenvironment-containerimage
	//
	ContainerImage *[]*string `field:"optional" json:"containerImage" yaml:"containerImage"`
}

