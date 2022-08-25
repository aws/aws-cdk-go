package awsstepfunctions


// Options for selecting the choice paths.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   afterwardsOptions := &afterwardsOptions{
//   	includeErrorHandlers: jsii.Boolean(false),
//   	includeOtherwise: jsii.Boolean(false),
//   }
//
type AfterwardsOptions struct {
	// Whether to include error handling states.
	//
	// If this is true, all states which are error handlers (added through 'onError')
	// and states reachable via error handlers will be included as well.
	IncludeErrorHandlers *bool `field:"optional" json:"includeErrorHandlers" yaml:"includeErrorHandlers"`
	// Whether to include the default/otherwise transition for the current Choice state.
	//
	// If this is true and the current Choice does not have a default outgoing
	// transition, one will be added included when .next() is called on the chain.
	IncludeOtherwise *bool `field:"optional" json:"includeOtherwise" yaml:"includeOtherwise"`
}

