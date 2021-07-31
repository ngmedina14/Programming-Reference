# Templates for all repetative format

------------

> # GIT

# CHANGELOG.md

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2021-07-31

### Added

- New template are published

### Changed

- Start using "changelog" over "change log" since it's the common usage.

### Removed

- Section about "changelog" vs "CHANGELOG".

### Fixed

- Fix typos in recent README changes.

### Deprecated

- Usage of API From is replace with When

------------


# CONTRIBUTING.md

- Neil Medina [ngmedina14@gmail.com](https://github.com/ngmedina14 "ngmedina14@gmail.com")

------------


# README.md

# Project Name

Project Tagline /Brief Description

> - Badges (Create your own badge)(**change** your **"account"** and **"repository"** with **yours**)
> - https://pkg.go.dev/github.com/account/repository
> - https://goreportcard.com/report/github.com/account/repository
> - https://about.codecov.io

Personal Project of - Neil Medina [ngmedina14@gmail.com](https://github.com/ngmedina14 "ngmedina14@gmail.com") (*Credit Creator*)

For Documentation:

- [Link of Documentation](https://github.com/ngmedina14/ordering-system "Link of Documentation")

Social Media:

- [Facebook](https://www.facebook.com)

## Screenshots

### GUI OF THE SYSTEM

![Dashboard](https://cdn.corporatefinanceinstitute.com/assets/systems-thinking.jpeg)
&nbsp;

## Features

-  Feature 1
-  Feature 2

## Minimum requirements

| Operating System                   |                Architectures              |                                Notes                                                |
|------------------------------------|-------------------------------------------|-------------------------------------------------------------------------------------|
| FreeBSD 10.3 or later              |  amd64, 386                               | Debian GNU/kFreeBSD not supported                                                   |
| Linux 2.6.23 or later with glibc   |  amd64, 386, arm, arm64, s390x, ppc64le   | CentOS/RHEL 5.x not supported. Install from source for other libc.                  |
| macOS 10.10 or later               |  amd64                                    | Use the clang or gcc<sup>†</sup> that comes with Xcode<sup>‡</sup> for cgo support. |
| Windows 7, Server 2008 R2 or later |  amd64, 386                               | Use MinGW gcc<sup>†</sup>. No need for cygwin or msys.                              |

- <sup>†</sup> A C compiler is required only if you plan to use cgo.
- <sup>‡</sup> You only need to install the command line tools for Xcode. If you have already installed Xcode 4.3+, you can install it from the Components tab of the Downloads preferences panel.

### Hardware

- RAM - minimum 256MB
- CPU - minimum 2GHz

### Software

- Go Version 1.15 or later

## Installation

```bash
$ go get -u github.com/account/repository/...
```

some description of the installation

Prepare modules

```bash
$ go mod init

$ go mod tidy
```

Run your app (Linux, Apple macOS or Windows):

```bash
$ go build; ./program
```

In Windows:

```bash
> go build && program.exe
```

# Quick Reference

Other Things Here

------------

# .gitignore

>For Reference

https://www.atlassian.com/git/tutorials/saving-changes/gitignore

|  PATTERN | EXAMPLE MATCHES   | EXPLANATION  |
| :------------ | :------------ | :------------ |
|  **/logs | logs/debug.log<br>logs/monday/foo.bar<br>build/logs/debug.log | You can prepend a pattern with a double asterisk to match directories anywhere in the repository.|
|  **/logs/debug.log | logs/debug.log<br>build/logs/debug.log<br>but not<br>logs/build/debug.log | You can also use a double asterisk to match files based on their name and the name of their parent directory. |
|  *.log  |  debug.log<br>foo.log<br>.log<br>logs/debug.log  |  An asterisk is a wildcard that matches zero or more characters.
|  *.log<br>!important.log  |  debug.log<br>trace.log<br>but not<br>important.log<br>logs/important.log  |  Prepending an exclamation mark to a pattern negates it. If a file matches a pattern, but also matches a negating pattern defined later in the file, it will not be ignored.
|  *.log<br>!important/*.log<br>trace.*  |  debug.log<br>important/trace.log<br>but not<br>important/debug.log  |  Patterns defined after a negating pattern will re-ignore any previously negated files.
|  /debug.log  |  debug.log<br>but not<br>logs/debug.log  |  Prepending a slash matches files only in the repository root.
|  debug.log  |  debug.log<br>logs/debug.log  |  By default, patterns match files in any directory
|  debug?.log  |  debug0.log<br>debugg.log<br>but not<br>debug10.log  |  A question mark matches exactly one character.
|  debug[0-9].log  |  debug0.log<br>debug1.log<br>but not<br>debug10.log  |  Square brackets can also be used to match a single character from a specified range.
|  debug[01].log  |  debug0.log<br>debug1.log<br>but not<br>debug2.log<br>debug01.log  |  Square brackets match a single character form the specified set.
|  debug[!01].log  |  debug2.log<br>but not<br>debug0.log<br>debug1.log<br>debug01.log  |  An exclamation mark can be used to match any character except one from the specified set.
|  debug[a-z].log  |  debuga.log<br>debugb.log<br>but not<br>debug1.log  |  Ranges can be numeric or alphabetic.
|  logs  |  logs<br>logs/debug.log<br>logs/latest/foo.bar<br>build/logs<br>build/logs/debug.log  |  If you don't append a slash, the pattern will match both files and the contents of directories with that name. In the example matches on the left, both directories and files named logs are ignored
|  logs/  |  logs/debug.log<br>logs/latest/foo.bar<br>build/logs/foo.bar<br>build/logs/latest/debug.log  |  Appending a slash indicates the pattern is a directory. The entire contents of any directory in the repository matching that name – including all of its files and subdirectories – will be ignored
|  logs/<br>!logs/important.log  |  logs/debug.log<br>logs/important.log  |  Wait a minute! Shouldn't logs/important.log be negated in the example on the left<br><br>Nope! Due to a performance-related quirk in Git, you can not negate a file that is ignored due to a pattern matching a directory
|  logs/**/debug.log  |  logs/debug.log<br>logs/monday/debug.log<br>logs/monday/pm/debug.log  |  A double asterisk matches zero or more directories.
|  logs/*day/debug.log  |  logs/monday/debug.log<br>logs/tuesday/debug.log<br>but not<br>logs/latest/debug.log  |  Wildcards can be used in directory names as well.
|  logs/debug.log  |  logs/debug.log<br>but not<br>debug.log<br>build/logs/debug.log  |  Patterns specifying a file in a particular directory are relative to the repository root. (You can prepend a slash if you like, but it doesn't do anything special.)

------------

# SECURITY.md

------------

# LICENSE

------------



> # GO

# .go files

Todo: discuss about how to add Description in every function and type
https://blog.golang.org/godoc