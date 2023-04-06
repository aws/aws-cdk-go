// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha


// Properties for a new Fleet alias.
//
// Example:
//   var fleet buildFleet
//
//
//   // Add an alias to an existing fleet using a dedicated fleet method
//   liveAlias := fleet.AddAlias(jsii.String("live"))
//
//   // You can also create a standalone alias
//   // You can also create a standalone alias
//   gamelift.NewAlias(this, jsii.String("TerminalAlias"), &AliasProps{
//   	AliasName: jsii.String("terminal-alias"),
//   	TerminalMessage: jsii.String("A terminal message"),
//   })
//
// Experimental.
type AliasProps struct {
	// Name of this alias.
	// Experimental.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// A human-readable description of the alias.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A fleet that the alias points to. If specified, the alias resolves to one specific fleet.
	//
	// At least one of `fleet` and `terminalMessage` must be provided.
	// Experimental.
	Fleet IFleet `field:"optional" json:"fleet" yaml:"fleet"`
	// The message text to be used with a terminal routing strategy.
	//
	// At least one of `fleet` and `terminalMessage` must be provided.
	// Experimental.
	TerminalMessage *string `field:"optional" json:"terminalMessage" yaml:"terminalMessage"`
}

