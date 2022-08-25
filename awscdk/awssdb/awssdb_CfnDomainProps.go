package awssdb


// Properties for defining a `CfnDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainProps := &cfnDomainProps{
//   	description: jsii.String("description"),
//   }
//
type CfnDomainProps struct {
	// Information about the SimpleDB domain.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

