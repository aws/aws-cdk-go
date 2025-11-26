package previewawsworkspaceswebmixins


// The filter that specifies the events to monitor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnSessionLoggerPropsMixin_EventFilterProperty struct {
	// The filter that monitors all of the available events, including any new events emitted in the future.
	//
	// The `All` and `Include` properties are not required, but one of them should be present. `{}` is a valid input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-eventfilter.html#cfn-workspacesweb-sessionlogger-eventfilter-all
	//
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// The filter that monitors only the listed set of events.
	//
	// New events are not auto-monitored. The `All` and `Include` properties are not required, but one of them should be present.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-eventfilter.html#cfn-workspacesweb-sessionlogger-eventfilter-include
	//
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

