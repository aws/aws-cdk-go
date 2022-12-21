// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   resourceA := awscdk.NewCfnResource(this, jsii.String("ResourceA"), resourceProps)
//   resourceB := awscdk.NewCfnResource(this, jsii.String("ResourceB"), resourceProps)
//
//   resourceB.addDependency(resourceA)
//
type CfnResourceProps struct {
	// CloudFormation resource type (e.g. `AWS::S3::Bucket`).
	Type *string `field:"required" json:"type" yaml:"type"`
	// Resource properties.
	Properties *map[string]interface{} `field:"optional" json:"properties" yaml:"properties"`
}

