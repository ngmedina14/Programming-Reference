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



> # GO

# .go files

Todo: discuss about how to add Description in every function and type
https://blog.golang.org/godoc