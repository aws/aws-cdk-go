package mixinsawsapptest


// Specifies the mainframe action properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mainframeActionPropertiesProperty := &MainframeActionPropertiesProperty{
//   	DmsTaskArn: jsii.String("dmsTaskArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeactionproperties.html
//
type CfnTestCasePropsMixin_MainframeActionPropertiesProperty struct {
	// The DMS task ARN of the mainframe action properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeactionproperties.html#cfn-apptest-testcase-mainframeactionproperties-dmstaskarn
	//
	DmsTaskArn *string `field:"optional" json:"dmsTaskArn" yaml:"dmsTaskArn"`
}

