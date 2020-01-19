package option

import (
	"flag"
)

type Option struct {
	Opt_f *bool
	Opt_i *bool
	Opt_n *bool
	Opt_q *bool
	Opt_Q *bool
	Opt_s *bool
	Opt_v *bool
	Opt_o *string
	Opt_p *string
	Opt_P *string
	Opt_w *bool
	Opt_W *bool
	Opt_C *bool
	Opt_L *bool
	Opt_M *bool
}
func NewOption() Option {
	f_p := new(Option)
	f := *f_p
	f.Opt_f = flag.Bool("f", false, "force overwriting of destination files. Not currently passed down to the mv/cp/ln command due to vagaries of implementations (but you can use -o-f to do that).")
	f.Opt_i = flag.Bool("i", false, "interactive: show each line to be executed and ask the user whether to execute it.  Y or y will execute it, anything else will skip it.  Note that you just need to type one character.")
	f.Opt_n = flag.Bool("n", false, "no execution: print what would happen, but don't do it.")
	f.Opt_q = flag.Bool("q", false, "Turn bare glob qualifiers off:  now assumed by default, so this has no effect.")
	f.Opt_Q = flag.Bool("Q", false, "Force bare glob qualifiers on.  Don't turn this on unless you are actually using glob qualifiers in a pattern (see below).")
	f.Opt_s = flag.Bool("s", false, "symbolic, passed down to ln; only works with zln or z?? -L.")
	f.Opt_v = flag.Bool("v", false, "verbose: print line as it's being executed.")
	f.Opt_o = flag.String("o", "", "<optstring> <optstring> will be split into words and passed down verbatim to the cp, ln or mv called to perform the work.  It will probably begin with a `-'.")
	f.Opt_p = flag.String("p", "", "<program> Call <program> instead of cp, ln or mv.  Whatever it does, it should at least understand the form '<program> -- <oldname> <newname>', where <oldname> and <newname> are filenames generated. <program> will be split into words.")
	f.Opt_P = flag.String("P", "", "<program> As -p, but the program doesn't understand the \"--\" convention.  In this case the file names must already be sane.")
	f.Opt_w = flag.Bool("w", false, "Pick out wildcard parts of the pattern, as described above, and implicitly add parentheses for referring to them.")
	f.Opt_W = flag.Bool("W", false, "Just like -w, with the addition of turning wildcards in the replacement pattern into sequential ${1} .. ${N} references.")
	f.Opt_C = flag.Bool("C", false, "Force cp, respectively, regardless of the name of the function.")
	f.Opt_L = flag.Bool("L", false, "Force ln, respectively, regardless of the name of the function.")
	f.Opt_M = flag.Bool("M", false, "Force mv, respectively, regardless of the name of the function.")
	return f
}
