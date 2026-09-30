package awsglue


// Configuration for parsing JSON responses from REST API calls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseConfigurationProperty := &ResponseConfigurationProperty{
//   	ResultPath: jsii.String("resultPath"),
//
//   	// the properties below are optional
//   	ErrorPath: jsii.String("errorPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-responseconfiguration.html
//
type CfnConnectionType_ResponseConfigurationProperty struct {
	// JSON path expression for result data location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-responseconfiguration.html#cfn-glue-connectiontype-responseconfiguration-resultpath
	//
	ResultPath *string `field:"required" json:"resultPath" yaml:"resultPath"`
	// JSON path expression for error information location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-responseconfiguration.html#cfn-glue-connectiontype-responseconfiguration-errorpath
	//
	ErrorPath *string `field:"optional" json:"errorPath" yaml:"errorPath"`
}

