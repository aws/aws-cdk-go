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
type CfnEndpointConfig_ClarifyShapBaselineConfigProperty struct {
	// `CfnEndpointConfig.ClarifyShapBaselineConfigProperty.MimeType`.
	MimeType *string `field:"optional" json:"mimeType" yaml:"mimeType"`
	// `CfnEndpointConfig.ClarifyShapBaselineConfigProperty.ShapBaseline`.
	ShapBaseline *string `field:"optional" json:"shapBaseline" yaml:"shapBaseline"`
	// `CfnEndpointConfig.ClarifyShapBaselineConfigProperty.ShapBaselineUri`.
	ShapBaselineUri *string `field:"optional" json:"shapBaselineUri" yaml:"shapBaselineUri"`
}

