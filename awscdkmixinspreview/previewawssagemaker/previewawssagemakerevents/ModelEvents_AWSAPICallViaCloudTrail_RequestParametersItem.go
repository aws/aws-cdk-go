package previewawssagemakerevents


// Type definition for RequestParametersItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParametersItem := &RequestParametersItem{
//   	InitialInstanceCount: []*string{
//   		jsii.String("initialInstanceCount"),
//   	},
//   	InitialVariantWeight: []*string{
//   		jsii.String("initialVariantWeight"),
//   	},
//   	InstanceType: []*string{
//   		jsii.String("instanceType"),
//   	},
//   	ModelName: []*string{
//   		jsii.String("modelName"),
//   	},
//   	VariantName: []*string{
//   		jsii.String("variantName"),
//   	},
//   }
//
// Experimental.
type ModelEvents_AWSAPICallViaCloudTrail_RequestParametersItem struct {
	// initialInstanceCount property.
	//
	// Specify an array of string values to match this event if the actual value of initialInstanceCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InitialInstanceCount *[]*string `field:"optional" json:"initialInstanceCount" yaml:"initialInstanceCount"`
	// initialVariantWeight property.
	//
	// Specify an array of string values to match this event if the actual value of initialVariantWeight is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InitialVariantWeight *[]*string `field:"optional" json:"initialVariantWeight" yaml:"initialVariantWeight"`
	// instanceType property.
	//
	// Specify an array of string values to match this event if the actual value of instanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// modelName property.
	//
	// Specify an array of string values to match this event if the actual value of modelName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Model reference.
	//
	// Experimental.
	ModelName *[]*string `field:"optional" json:"modelName" yaml:"modelName"`
	// variantName property.
	//
	// Specify an array of string values to match this event if the actual value of variantName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VariantName *[]*string `field:"optional" json:"variantName" yaml:"variantName"`
}

