# eclipse-workspace-themer

Set the theme for a given workspace for Eclipse IDE. Focuses on the Java IDE/JDT specifically, and only changes the text editor colours. Currently supplies a few different themes that I considered interesting. Uses [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss).

## Building

1. clone
2. `go build`
3. copy `eclipse-workspace-themer` onto `$PATH`

## Usage

1. `cd` into your Eclipse workspace
2. Run `eclipse-workspace-themer`, select theme with arrow keys, then press enter.

## Limitations
- Only themes the text editor portion; any more than this gets very complex very fast
- https://github.com/eclipse-jdt/eclipse.jdt.ui/issues/2971
- Check the issue tracker as well

## Adding Themes
It should be fairly self explanatory; add an entry to the map, the key is the name of the theme. Feel free to open a PR if you've implemented an interesting one that you want to share.

If you want more fine grain control around eg. colouring class and interface names differently, I'm open to that. Feel free to raise an issue about it.
