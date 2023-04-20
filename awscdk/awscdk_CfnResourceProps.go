// An experiment to bundle the entire CDK into a single module
package awscdk


// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/constructs-go/constructs"
//
//   type myAspect struct {
//   }
//
//   func (this *myAspect) visit(node iConstruct) {
//   	if *node instanceof cdk.CfnResource && *node.CfnResourceType == "Foo::Bar" {
//   		this.error(*node, jsii.String("we do not want a Foo::Bar resource"))
//   	}
//   }
//
//   func (this *myAspect) error(node iConstruct, message *string) {
//   	cdk.Annotations_Of(*node).AddError(*message)
//   }
//
//   type myStack struct {
//   	stack
//   }
//
//   func newMyStack(scope construct, id *string) *myStack {
//   	this := &myStack{}
//   	cdk.NewStack_Override(this, scope, id)
//
//   	stack := cdk.NewStack()
//   	cdk.NewCfnResource(stack, jsii.String("Foo"), &CfnResourceProps{
//   		Type: jsii.String("Foo::Bar"),
//   		Properties: map[string]interface{}{
//   			"Fred": jsii.String("Thud"),
//   		},
//   	})
//   	cdk.Aspects_Of(stack).Add(NewMyAspect())
//   	return this
//   }
//
// Experimental.
type CfnResourceProps struct {
	// CloudFormation resource type (e.g. `AWS::S3::Bucket`).
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Resource properties.
	// Experimental.
	Properties *map[string]interface{} `field:"optional" json:"properties" yaml:"properties"`
}

