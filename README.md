# ahkpm - The AutoHotKey Package Manager.

[![build](https://github.com/joshuacc/ahkpm/workflows/build/badge.svg?branch=main)](https://github.com/joshuacc/ahkpm/actions/workflows/build.yml?query=branch%3Amain) [![Maintainability](https://api.codeclimate.com/v1/badges/d879856f8da1f8b803a6/maintainability)](https://codeclimate.com/github/joshuacc/ahkpm/maintainability) [![codecov](https://codecov.io/gh/joshuacc/ahkpm/branch/main/graph/badge.svg?token=CHRB93N4U7)](https://codecov.io/gh/joshuacc/ahkpm)

[AutoHotKey][ahk] is a powerful tool for scripting automations on Windows,
but managing dependencies for your scripts is painful.
`ahkpm` intends to bring modern package management to AutoHotkey,
making it easier than ever to automate away the drudgery.

## Commands

```
Usage:
  ahkpm [command]

Available Commands:
  cache clean  Clears the package cache
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  init         Interactively create an ahkpm.json file in the current directory
  install      Installs either the specified package or all packages listed in ahkpm.json

Flags:
  -h, --help   help for ahkpm
```

## Installation

To install ahkpm:

1. Go to the [releases][releases] page and find the most recent version.
2. Download the `ahkpm-{version}.msi` file.
3. Open it on your Windows machine to launch the ahkpm installer.

## Basic usage

1. Open the command line and navigate to the directory which will contain your AutoHotKey script.
2. Run `ahkpm init` and answer the prompts to create an `ahkpm.json` file
3. Run `ahkpm install <package>@<version>`
   - The package can be any github repository in the form: `github.com/user/repo`
   - The version can be any of the following:
     - A valid [semantic version][semver] such as `1.0.0`
     - The prefix `tag:` followed by the name of a tag in the package's repository, such as `tag:beta2`
     - The prefix `branch:` followed by the name of a branch in the package's repository, such as `branch:main`
     - The prefix `commit:` followed by the hash of a commit in the package's repository, such as `commit:badcce14f8e828cda4d8ac404a12448700de1441`
     - Omitting the version is not yet supported
4. Add `#Include, %A_ScriptDir%` to the top of your script to set the current directory as the context for subsequent includes
5. Add `#Include, ahkpm-modules\github.com\user\repo\main-file.ahk` to your script
6. You can now use the package's functionality within your AutoHotKey script!

## ahkpm.json

```jsonc
{
  "name": "my-project",
  "version": "0.0.1",
  "description": "A brief description of what the package does",
  // URL for the package's git repository
  "repository": "github.com/user/my-project",
  // URL for the package's main website
  "website": "example.com",
  // The SPDX License identifier for the package's software license
  "license": "MIT",
  // URL for users to file bugs/issues for the package
  "issueTracker": "github.com/user/my-project/issues",
  // Information about the primary author of the package
  "author": {
    "name": "joshuacc",
    "email": "",
    "website": "joshuaclanton.dev"
  },
  // Lists all dependencies along with the required version
  "dependencies": {
    "github.com/user/repo1": "1.0.0",
    "github.com/user/repo2": "tag:beta2",
    "github.com/user/repo3": "branch:main",
    "github.com/user/repo4": "commit:badcce14f8e828cda4d8ac404a12448700de1441"
  }
}
```

## Current limitations

ahkpm is being actively developed, but it is still a young project.
As a result it has the following limitations.

- It only supports hosting and downloading of packages on GitHub, though other git hosts will be supported in the future.
- It doesn't (yet) resolve dependencies of dependencies, so it isn't (yet) useful for managing the dependencies of a reusable AHK library.
- It doesn't (yet) support specifying version ranges as you can in npm and other package managers.
- It is not (yet) conveniently packaged into a Windows installer
- It doesn't (yet) have a lockfile like other package managers, which means that your dependency's code can unexpectedly change out from under you when you run `ahkpm install` to reinstall it, especially if you do something like specifying a `branch:` version.

If you'd like to help remedy these limitations, consider contributing!

## Contributing to ahkpm

See the [contribution guidelines](./CONTRIBUTING.md)

[ahk]:https://www.autohotkey.com/
[semver]:https://semver.org/
[releases]:https://github.com/joshuacc/ahkpm/releases

## ahkpm.lock

```jsonc
{
    "lockFileVersion": 1,
    // This should exactly match ahkpm.json unless it has been edited manually
    "dependencies": {
      "github.com/joshuacc/mock-ahkpm-package-a": "branch:main",
    },
    "resolved": {
        "ahkpm-modules/github.com/joshuacc/mock-ahkpm-package-a": {
            "version": "branch:main",
            "gitSha": "c5b8f8d0d0d1e5c9a5f7f8b8b5c9a5f7f8b8b5c9a",
            "repository": "github.com/joshuacc/mock-ahkpm-package-a",
            "dependencies": {
                "github.com/joshuacc/mock-ahkpm-package-b": "^1.0.0"
            }
        },
        "ahkpm-modules/github.com/joshuacc/mock-ahkpm-package-a/ahkpm-modules/github.com/joshuacc/mock-ahkpm-package-b": {
            "version": "1.0.2",
            "repository": "github.com/joshuacc/mock-ahkpm-package-b",
            "gitSha": "c5b8f8d0d0d1e5c9a5f7f8b8b5c9a5f7f8b8b5c9a",
            "dependencies": {}
        }
    }
}
```