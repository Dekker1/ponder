# Ponder
Ponder is a small command line tool that will help you manage your sheet music libraries. Ponder currently has support for sheet music music in the lilypond format and PDF documents in containing sheet music. The application's main purpose is the creation of songbooks without the hassle of having to include all the pieces of music manually.

## Prerequisites
Ponder requires a working lilypond and LaTeX (`latexmk` & `pdflatex`) installation.

## Installation
Ponder hasn't been released yet. So for now you'll have to build the application yourself. (This will require `git`, `go`, and the go dependency manager `glide`).

1. `git clone git@github.com:jjdekker/ponder.git $GOPATH/src/github.com/jjdekker/ponder`
2. `cd $GOPATH/src/github.com/jjdekker/ponder`
3. `glide install`
4. `go install`

Make sure that the created binary is on the executable path.

## Usage
If the information you're looking for isn't in the text below, try the `ponder help` command.

Ponder works from an repository folder, you can create one using the command `ponder init [PATH]`. Options for the repository are located in the `ponder.json` file. You can find all possible options in the code documentation.

### Adding Sheet Music
You can add sheet music to the book using the command `ponder add [PATH]...`. Each score has the following properties that will influence the compilation of the book:

- Name
- Categories

These properties will put into a separate JSON file in case of an PDF file, and are embedded in the lilypond files.

### Compiling
Once the music is added, ponder will automatically take these into account when compiling. You can compile scores using the `ponder compile` command. This will compile all scores in the repository and place them in the output directory.

You can also compile a songbook using the `ponder book` command. This will first compile all scores, and then compile a songbook for all scores in the repository.

## FAQ
**How do I add a page of text in between songs?**

There is no great way to do this yet. Hopefully we can add this soon. In the meantime, put your text in a pdf and pretend it's a song.
