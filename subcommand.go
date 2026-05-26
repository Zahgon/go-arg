package arg

// Subcommand returns the user struct for the subcommand selected by
// the command line arguments most recently processed by the parser.
// The return value is always a pointer to a struct. If no subcommand
// was specified then it returns the top-level arguments struct. If
// no command line arguments have been processed by this parser then it
// returns nil.
func (p *Parser) Subcommand() interface{} { _ = "STUB: not implemented"; return nil }

// SubcommandNames returns the sequence of subcommands specified by the
// user. If no subcommands were given then it returns an empty slice.
func (p *Parser) SubcommandNames() []string { _ = "STUB: not implemented"; return nil }

// lookupCommand finds a subcommand based on a sequence of subcommand names. The
// first string should be a top-level subcommand, the next should be a child
// subcommand of that subcommand, and so on. If no strings are given then the
// root command is returned. If no such subcommand exists then an error is
// returned.
func (p *Parser) lookupCommand(path ...string) (*command, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
