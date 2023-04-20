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
// Experimental.
type CustomWidgetProps struct {
	// The Arn of the AWS Lambda function that returns HTML or JSON that will be displayed in the widget.
	// Experimental.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The title of the widget.
	// Experimental.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Height of the widget.
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Parameters passed to the lambda function.
	// Experimental.
	Params interface{} `field:"optional" json:"params" yaml:"params"`
	// Update the widget on refresh.
	// Experimental.
	UpdateOnRefresh *bool `field:"optional" json:"updateOnRefresh" yaml:"updateOnRefresh"`
	// Update the widget on resize.
	// Experimental.
	UpdateOnResize *bool `field:"optional" json:"updateOnResize" yaml:"updateOnResize"`
	// Update the widget on time range change.
	// Experimental.
	UpdateOnTimeRangeChange *bool `field:"optional" json:"updateOnTimeRangeChange" yaml:"updateOnTimeRangeChange"`
	// Width of the widget, in a grid of 24 units wide.
	// Experimental.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

