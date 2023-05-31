package awsquicksight


// A structure that represents a semantic entity type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   semanticEntityTypeProperty := &SemanticEntityTypeProperty{
//   	SubTypeName: jsii.String("subTypeName"),
//   	TypeName: jsii.String("typeName"),
//   	TypeParameters: map[string]*string{
//   		"typeParametersKey": jsii.String("typeParameters"),
//   	},
//   }
//
type CfnTopic_SemanticEntityTypeProperty struct {
	// The semantic entity sub type name.
	SubTypeName *string `field:"optional" json:"subTypeName" yaml:"subTypeName"`
	// The semantic entity type name.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// The semantic entity type parameters.
	TypeParameters interface{} `field:"optional" json:"typeParameters" yaml:"typeParameters"`
}

