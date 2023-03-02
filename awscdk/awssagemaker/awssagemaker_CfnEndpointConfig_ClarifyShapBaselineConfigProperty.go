package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clarifyShapBaselineConfigProperty := &clarifyShapBaselineConfigProperty{
//   	mimeType: jsii.String("mimeType"),
//   	shapBaseline: jsii.String("shapBaseline"),
//   	shapBaselineUri: jsii.String("shapBaselineUri"),
//   }
//
type CfnEndpointConfig_ClarifyShapBaselineConfigProperty struct {
	// `CfnEndpointConfig.ClarifyShapBaselineConfigProperty.MimeType`.
	MimeType *string `field:"optional" json:"mimeType" yaml:"mimeType"`
	// `CfnEndpointConfig.ClarifyShapBaselineConfigProperty.ShapBaseline`.
	ShapBaseline *string `field:"optional" json:"shapBaseline" yaml:"shapBaseline"`
	// `CfnEndpointConfig.ClarifyShapBaselineConfigProperty.ShapBaselineUri`.
	ShapBaselineUri *string `field:"optional" json:"shapBaselineUri" yaml:"shapBaselineUri"`
}

