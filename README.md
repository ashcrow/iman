iman
====
Reading man pages from images

**This is a quick POC**: If it ends up being a helpful idea it will be fleshed out to be more than exec'ing commands :-)


In a nutshell, this command:

- Starts a container
- Copies the help.1 from the containers root filesystem
- Uses man to show the help file
- Destroys the container


Example
-------
```
[steve@elpis iman]$ sudo ./iman image-helpgen:1.0
IMAGE-HELPGEN(1)                                                                 May 2018                                                                IMAGE-HELPGEN(1)

NAME
       image-helpgen - Create container help pages

DESCRIPTION
       The image-helpgen command lets you create a container help page from either guided prompts or content from a Dockerfile. Output from image-helpgen is a help.md
       file (in Markdown format) and/or a help.1 file (in manpage format) that describes the container, its uses, and possible security issues.

       By placing the help.1 file in a container image's root directory (/help.1), it can be displayed by various container tools.

USAGE
       image-helpgen <command> [args]

OPTIONS
       guide
             Prompts for container image content (name, usage, etc.) and produces help.1 and help.md files.

       dockerfile
             Parses a Dockerfile and generates a help page template in Markdown format (help.md).

       man
             Generates container help page (help.1) in manpage format from a completed Markdown file (help.md).

       version
             Shows version information and exits.

CONTAINER HELP PAGE
       Whether using the guide or dockerfile option, the following sections are created in the resulting help page:

       * NAME: The NAME line is constructed from the values of “LABEL name”  + “LABEL summary” values in the Dockerfile.

       * DESCRIPTION: All commented lines (#) at the beginning of the Dockerfile are used as the help file’s DESCRIPTION. Add a line with just a # to have separate
       paragraphs. A line not beginning with a comment character (#) ends the description.

       * ENVIRONMENT VARIABLES: The variable name and default setting for all lines beginning with ENV in the Dockerfile are added to the ENVIRONMENT VARIABLES table.

       * SECURITY IMPLICATIONS: The security implications section is made up of the following subsections:

           * Ports: Port numbers on EXPOSE lines are added to the Ports table.
           * Volumes: Directories listed on VOLUME lines are added to the Volumes table.
           * Daemon: If “-d” is on the usage line, text notes the container runs as a daemon.
           * Expected Capabilities: Each --cap-add on the usage line adds an entry to the capabilities table.

       * SEE ALSO: The value of any “LABEL url” line from the Dockerfile is added to the SEE ALSO section.

       * Headings and footers: The header is created from this Dockerfile information: "LABEL name="  Month/Year "LABEL name="(2) and the footer is created from: ”LABEL
       maintainer” Container Image Pages “LABEL name=”(2)

EXAMPLES
           image-helpgen dockerfile -dockerfile Dockerfile
                   # Creates a help.md file from the Dockerfile in the current directory
                   # Descriptions need to be added manually (look for TODO lines)

           image-helpgen guide
             Image name: myownimage
                   # Creates a help.md and help.1 file from content you input from prompts

           image-helpgen man
                   # Produces a help.1 file from the help.md file in the current directory

FILES
       /etc/image-helpgen/template.tpl
             Template file used to create container help pages.

SEE ALSO
       https://github.com/ashcrow/image-helpgen

Steve Milner                                                                  User Commands                                                              IMAGE-HELPGEN(1)
```