package previewawssagemakerevents


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseElements := &ResponseElements{
//   	EndpointConfigArn: []*string{
//   		jsii.String("endpointConfigArn"),
//   	},
//   	ModelArn: []*string{
//   		jsii.String("modelArn"),
//   	},
//   	TrainingJobArn: []*string{
//   		jsii.String("trainingJobArn"),
//   	},
//   	TransformJobArn: []*string{
//   		jsii.String("transformJobArn"),
//   	},
//   }
//
// Experimental.
type ModelEvents_AWSAPICallViaCloudTrail_ResponseElements struct {
	// endpointConfigArn property.
	//
	// Specify an array of string values to match this event if the actual value of endpointConfigArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndpointConfigArn *[]*string `field:"optional" json:"endpointConfigArn" yaml:"endpointConfigArn"`
	// modelArn property.
	//
	// Specify an array of string values to match this event if the actual value of modelArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ModelArn *[]*string `field:"optional" json:"modelArn" yaml:"modelArn"`
	// trainingJobArn property.
	//
	// Specify an array of string values to match this event if the actual value of trainingJobArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TrainingJobArn *[]*string `field:"optional" json:"trainingJobArn" yaml:"trainingJobArn"`
	// transformJobArn property.
	//
	// Specify an array of string values to match this event if the actual value of transformJobArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformJobArn *[]*string `field:"optional" json:"transformJobArn" yaml:"transformJobArn"`
}

