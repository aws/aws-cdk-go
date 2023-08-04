package awscdk


// Example:
//   type myConstruct struct {
//   	resource
//   	tags
//   }
//
//   func newMyConstruct(scope construct, id *string) *myConstruct {
//   	this := &myConstruct{}
//   	newResource_Override(this, scope, id)
//
//   	awscdk.NewCfnResource(this, jsii.String("Resource"), &cfnResourceProps{
//   		Type: jsii.String("Whatever::The::Type"),
//   		Properties: map[string]interface{}{
//   			// ...
//   			"Tags": this.tags.renderedTags,
//   		},
//   	})
//   	return this
//   }
//
type CfnResourceProps struct {
	// CloudFormation resource type (e.g. `AWS::S3::Bucket`).
	Type *string `field:"required" json:"type" yaml:"type"`
	// Resource properties.
	// Default: - No resource properties.
	//
	Properties *map[string]interface{} `field:"optional" json:"properties" yaml:"properties"`
}

