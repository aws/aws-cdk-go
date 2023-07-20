package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clarifyShapBaselineConfigProperty := &ClarifyShapBaselineConfigProperty{
//   	MimeType: jsii.String("mimeType"),
//   	ShapBaseline: jsii.String("shapBaseline"),
//   	ShapBaselineUri: jsii.String("shapBaselineUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapbaselineconfig.html
//
type CfnEndpointConfig_ClarifyShapBaselineConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapbaselineconfig.html#cfn-sagemaker-endpointconfig-clarifyshapbaselineconfig-mimetype
	//
	MimeType *string `field:"optional" json:"mimeType" yaml:"mimeType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapbaselineconfig.html#cfn-sagemaker-endpointconfig-clarifyshapbaselineconfig-shapbaseline
	//
	ShapBaseline *string `field:"optional" json:"shapBaseline" yaml:"shapBaseline"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapbaselineconfig.html#cfn-sagemaker-endpointconfig-clarifyshapbaselineconfig-shapbaselineuri
	//
	ShapBaselineUri *string `field:"optional" json:"shapBaselineUri" yaml:"shapBaselineUri"`
}

