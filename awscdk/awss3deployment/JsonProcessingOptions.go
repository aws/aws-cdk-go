package awss3deployment


// Define options which can be passed using the method `Source.jsonData()`.
//
// Example:
//   var bucket bucket
//   var param stringParameter
//
//
//   // Example with a secret value that contains double quotes
//   deployment := s3deploy.NewBucketDeployment(this, jsii.String("JsonDeployment"), &BucketDeploymentProps{
//   	Sources: []iSource{
//   		s3deploy.Source_JsonData(jsii.String("config.json"), map[string]interface{}{
//   			"api_endpoint": jsii.String("https://api.example.com"),
//   			"secretValue": param.stringValue,
//   			 // value with double quotes
//   			"config": map[string]interface{}{
//   				"enabled": jsii.Boolean(true),
//   				"features": []*string{
//   					jsii.String("feature1"),
//   					jsii.String("feature2"),
//   				},
//   			},
//   		}, &JsonProcessingOptions{
//   			Escape: jsii.Boolean(true),
//   		}),
//   	},
//   	DestinationBucket: bucket,
//   })
//
type JsonProcessingOptions struct {
	// If set to `true`, the marker substitution will make sure the value inserted in the file will be a valid JSON string.
	// Default: - false.
	//
	Escape *bool `field:"optional" json:"escape" yaml:"escape"`
}

