package awsglue


// Configuration for parsing JSON responses from REST API calls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   responseConfigurationProperty := &ResponseConfigurationProperty{
//   	ErrorPath: jsii.String("errorPath"),
//   	ResultPath: jsii.String("resultPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-responseconfiguration.html
//
type CfnConnectionTypePropsMixin_ResponseConfigurationProperty struct {
	// JSON path expression for error information location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-responseconfiguration.html#cfn-glue-connectiontype-responseconfiguration-errorpath
	//
	ErrorPath *string `field:"optional" json:"errorPath" yaml:"errorPath"`
	// JSON path expression for result data location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-responseconfiguration.html#cfn-glue-connectiontype-responseconfiguration-resultpath
	//
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
}

