package previewawsiottwinmakermixins


// The current status of the entity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var error interface{}
//
//   statusProperty := &StatusProperty{
//   	Error: error,
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-status.html
//
type CfnEntityPropsMixin_StatusProperty struct {
	// The error message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-status.html#cfn-iottwinmaker-entity-status-error
	//
	Error interface{} `field:"optional" json:"error" yaml:"error"`
	// The current state of the entity, component, component type, or workspace.
	//
	// Valid Values: `CREATING | UPDATING | DELETING | ACTIVE | ERROR`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-status.html#cfn-iottwinmaker-entity-status-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

