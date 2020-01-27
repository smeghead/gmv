# gmv #

This is a multiple move based on pattern matching. here are some basic examples:

```bash
gmv '(*).txt' '$1.lis'
```

Rename foo.txt to foo.lis, etc.  The parenthesis is the thing that
gets replaced by the $1 (not the `*', as happens in mmv, and note the
`$', not `=', so that you need to quote both words).

```bash
gmv '(**/)(*).txt '$1$2.lis'
```

The same, but scanning through subdirectories.  The $1 becomes the full
path.  Note that you need to write it like this; you can't get away with
'(**/*).txt'.

```bash
gmv -w '**/*.txt' '$1$2.lis'
gmv -W '**/*.txt' '**/*.lis'
```

These are the lazy version of the one above; with -w, gmv inserts the
parentheses for you in the search pattern, and with -W it also inserts
the numbered variables for you in the replacement pattern.  The catch
in the first version is that you don't need the / in the replacement
pattern.  (It's not really a catch, since $1 can be empty.)  Note that
-W actually inserts ${1}, ${2}, etc., so it works even if you put a
number after a wildcard (such as gmv -W '*1.txt' '*2.txt').

```bash
gmv -C '**/(*).txt' ~/save/'$1'.lis
```

Copy, instead of move, all .txt files in subdirectories to .lis files
in the single directory `~/save'.  Note that the ~ was not quoted.
You can test things safely by using the `-n' (no, not now) option.
Clashes, where multiple files are renamed or copied to the same one, are
picked up.

## Options: ##

- -f  force overwriting of destination files.  Not currently passed
     down to the mv/cp/ln command due to vagaries of implementations
     (but you can use -o-f to do that).
- -i  interactive: show each line to be executed and ask the user whether
     to execute it.  Y or y will execute it, anything else will skip it.
     Note that you just need to type one character.
- -n  no execution: print what would happen, but don't do it.
- -q  Turn bare glob qualifiers off:  now assumed by default, so this
     has no effect.
- -Q  Force bare glob qualifiers on.  Don't turn this on unless you are
     actually using glob qualifiers in a pattern (see below).
- -s  symbolic, passed down to ln; only works with zln or z?? -L.
- -v  verbose: print line as it's being executed.
- -o <optstring>
     <optstring> will be split into words and passed down verbatim
     to the cp, ln or mv called to perform the work.  It will probably
     begin with a `-'.
- -p <program>
     Call <program> instead of cp, ln or mv.  Whatever it does, it should
     at least understand the form '<program> -- <oldname> <newname>',
     where <oldname> and <newname> are filenames generated. <program>
     will be split into words.
- -P <program>
     As -p, but the program doesn't understand the "--" convention.
     In this case the file names must already be sane.
- -w  Pick out wildcard parts of the pattern, as described above, and
     implicitly add parentheses for referring to them.
- -W  Just like -w, with the addition of turning wildcards in the
     replacement pattern into sequential ${1} .. ${N} references.
- -C
- -L
- -M  Force cp, ln or mv, respectively, regardless of the name of the
     function.
