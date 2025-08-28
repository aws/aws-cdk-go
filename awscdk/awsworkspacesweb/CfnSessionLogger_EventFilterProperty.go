package awsworkspacesweb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   eventFilterProperty := &EventFilterProperty{
//   	All: all,
//   	Include: []*string{
//   		jsii.String("include"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-eventfilter.html
//
type CfnSessionLogger_EventFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-eventfilter.html#cfn-workspacesweb-sessionlogger-eventfilter-all
	//
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-eventfilter.html#cfn-workspacesweb-sessionlogger-eventfilter-include
	//
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

