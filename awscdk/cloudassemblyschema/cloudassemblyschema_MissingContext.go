package cloudassemblyschema


// Represents a missing piece of context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   missingContext := &missingContext{
//   	key: jsii.String("key"),
//   	props: &amiContextQuery{
//   		account: jsii.String("account"),
//   		filters: map[string][]*string{
//   			"filtersKey": []*string{
//   				jsii.String("filters"),
//   			},
//   		},
//   		region: jsii.String("region"),
//
//   		// the properties below are optional
//   		lookupRoleArn: jsii.String("lookupRoleArn"),
//   		owners: []interface{}{
//   			jsii.String("owners"),
//   		},
//   	},
//   	provider: awscdk.Cloud_assembly_schema.contextProvider_AMI_PROVIDER,
//   }
//
type MissingContext struct {
	// The missing context key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// A set of provider-specific options.
	Props interface{} `field:"required" json:"props" yaml:"props"`
	// The provider from which we expect this context key to be obtained.
	Provider ContextProvider `field:"required" json:"provider" yaml:"provider"`
}

