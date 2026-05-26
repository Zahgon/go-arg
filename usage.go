package arg

import (
	"io"
)

// the width of the left column
const colWidth = 25

// Fail prints usage information to p.Config.Out and exits with status code 2.
func (p *Parser) Fail(msg string) { _ = "STUB: not implemented"; return }

// FailSubcommand prints usage information for a specified subcommand to p.Config.Out,
// then exits with status code 2. To write usage information for a top-level
// subcommand, provide just the name of that subcommand. To write usage
// information for a subcommand that is nested under another subcommand, provide
// a sequence of subcommand names starting with the top-level subcommand and so
// on down the tree.
func (p *Parser) FailSubcommand(msg string, subcommand ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteUsage writes usage information to the given writer
func (p *Parser) WriteUsage(w io.Writer) { _ = "STUB: not implemented"; return }

// WriteUsageForSubcommand writes the usage information for a specified
// subcommand. To write usage information for a top-level subcommand, provide
// just the name of that subcommand. To write usage information for a subcommand
// that is nested under another subcommand, provide a sequence of subcommand
// names starting with the top-level subcommand and so on down the tree.
func (p *Parser) WriteUsageForSubcommand(w io.Writer, subcommand ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// print the beginning of the usage string

// write the option component of the usage message

// prefix with a space

// prefix with a space

// When we parse positionals, we check that:
//  1. required positionals come before non-required positionals
//  2. there is at most one multiple-value positional
//  3. if there is a multiple-value positional then it comes after all other positionals
// Here we merely print the usage string, so we do not explicitly re-enforce those rules

// write the positionals in following form:
//    REQUIRED1 REQUIRED2
//    REQUIRED1 REQUIRED2 [OPTIONAL1 [OPTIONAL2]]
//    REQUIRED1 REQUIRED2 REPEATED [REPEATED ...]
//    REQUIRED1 REQUIRED2 [REPEATEDOPTIONAL [REPEATEDOPTIONAL ...]]
//    REQUIRED1 REQUIRED2 [OPTIONAL1 [REPEATEDOPTIONAL [REPEATEDOPTIONAL ...]]]

// if the program supports subcommands, give a hint to the user about their existence

// print prints a line like this:
//
//	--option FOO            A description of the option [default: 123]
//
// If the text on the left is longer than a certain threshold, the description is moved to the next line:
//
//	--verylongoptionoption VERY_LONG_VARIABLE
//	                        A description of the option [default: 123]
//
// If multiple "extras" are provided then they are put inside a single set of square brackets:
//
//	--option FOO            A description of the option [default: 123, env: FOO]
func print(w io.Writer, item, description string, bracketed ...string) {
	_ = "STUB: not implemented"
	return
}

func withDefault(s string) string { _ = "STUB: not implemented"; return "" }

func withEnv(env string) string { _ = "STUB: not implemented"; return "" }

// WriteHelp writes the usage string followed by the full help string for each option
func (p *Parser) WriteHelp(w io.Writer) { _ = "STUB: not implemented"; return }

// WriteHelpForSubcommand writes the usage string followed by the full help
// string for a specified subcommand. To write help for a top-level subcommand,
// provide just the name of that subcommand. To write help for a subcommand that
// is nested under another subcommand, provide a sequence of subcommand names
// starting with the top-level subcommand and so on down the tree.
func (p *Parser) WriteHelpForSubcommand(w io.Writer, subcommand ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// obtain a flattened list of options from all ancestors
// also determine if any ancestor has a version option spec

// write the list of positionals

// write the list of options with the short-only ones first to match the usage string

// write the list of global options

// write the list of built in options

// write the list of environment only variables

// write the list of subcommands

// skip this subcommand in the help message

func (p *Parser) printOption(w io.Writer, spec *spec) { _ = "STUB: not implemented"; return }

func (p *Parser) printEnvOnlyVar(w io.Writer, spec *spec) { _ = "STUB: not implemented"; return }

func synopsis(spec *spec, form string) string {
	_ = "STUB: not implemented"
	// if the user omits the placeholder tag then we pick one automatically,
	// but if the user explicitly specifies an empty placeholder then we
	// leave out the placeholder in the help message
	return ""
}
