package awsappsync


// Options for GraphQL Types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var intermediateType iIntermediateType
//
//   graphqlTypeOptions := &graphqlTypeOptions{
//   	intermediateType: intermediateType,
//   	isList: jsii.Boolean(false),
//   	isRequired: jsii.Boolean(false),
//   	isRequiredList: jsii.Boolean(false),
//   }
//
// Experimental.
type GraphqlTypeOptions struct {
	// property determining if this attribute is a list i.e. if true, attribute would be [Type].
	// Experimental.
	IsList *bool `field:"optional" json:"isList" yaml:"isList"`
	// property determining if this attribute is non-nullable i.e. if true, attribute would be Type!
	// Experimental.
	IsRequired *bool `field:"optional" json:"isRequired" yaml:"isRequired"`
	// property determining if this attribute is a non-nullable list i.e. if true, attribute would be [ Type ]! or if isRequired true, attribe would be [ Type! ]!
	// Experimental.
	IsRequiredList *bool `field:"optional" json:"isRequiredList" yaml:"isRequiredList"`
	// the intermediate type linked to this attribute.
	// Experimental.
	IntermediateType IIntermediateType `field:"optional" json:"intermediateType" yaml:"intermediateType"`
}

