package awscloudwatch


// The properties for a CustomWidget.
//
// Example:
//   var dashboard dashboard
//
//
//   // Import or create a lambda function
//   fn := lambda.Function_FromFunctionArn(dashboard, jsii.String("Function"), jsii.String("arn:aws:lambda:us-east-1:123456789012:function:MyFn"))
//
//   dashboard.AddWidgets(cloudwatch.NewCustomWidget(&CustomWidgetProps{
//   	FunctionArn: fn.FunctionArn,
//   	Title: jsii.String("My lambda baked widget"),
//   }))
//
type CustomWidgetProps struct {
	// The Arn of the AWS Lambda function that returns HTML or JSON that will be displayed in the widget.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The title of the widget.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Height of the widget.
	// Default: - 6 for Alarm and Graph widgets.
	// 3 for single value widgets where most recent value of a metric is displayed.
	//
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Parameters passed to the lambda function.
	// Default: - no parameters are passed to the lambda function.
	//
	Params interface{} `field:"optional" json:"params" yaml:"params"`
	// Update the widget on refresh.
	// Default: true.
	//
	UpdateOnRefresh *bool `field:"optional" json:"updateOnRefresh" yaml:"updateOnRefresh"`
	// Update the widget on resize.
	// Default: true.
	//
	UpdateOnResize *bool `field:"optional" json:"updateOnResize" yaml:"updateOnResize"`
	// Update the widget on time range change.
	// Default: true.
	//
	UpdateOnTimeRangeChange *bool `field:"optional" json:"updateOnTimeRangeChange" yaml:"updateOnTimeRangeChange"`
	// Width of the widget, in a grid of 24 units wide.
	// Default: 6.
	//
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

