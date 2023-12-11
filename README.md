# Commit Msg Helper

CLI tool to help you write consistent commits

## Setup

You can install by download the binary in the release or manually installing it yourself.

### Release

Go to [releases](https://gitlab.eposnow.io/david.ribeiro/commit-msg-helper/-/releases) and download the latest binary according to your platform. If it's not there, follow the manual installation below.

Then be sure to put it somewhere where that's defined in your `$PATH` and rename it to `ch` so that everything works properly.

### Manual Installation

This requires version `1.23` of go.

Clone the repo and in the root of it, run `go install`.
This should install to the `/bin` folder of where you installed go.

For me, in Windows, it's at:
> C:\Users\\\<user>\go\bin

You can find out where it is by printing or checking the env variable `$GOPATH`.

## Recommended Flow

The recommended flow from setup to working is:
1. Run `go install` in the root of the repository or move the correct binary into a folder that's present on your `$PATH`
2. After install, setup a git alias to run the command through git with `ch aliases`. This will allow you to run `ch` with `git ch`
3. Lastly, run the command each time you need to commit `git ch`. The steps done for these command are documented above

If you want to remove the alias from git then you just need to run `ch cleanup`.

## Commands

- `aliases`
- `ch` ( which is just the root command by itself )
- `cleanup`

### aliases

Use this to set git aliases to make it easier in your git flow.

This will only be needed once.

Currently, only adds alias for the `ch` command.

### ch

This'll be your bread and butter.

This will ask for the commit type, ticket id (if any) and lastly a commit message.
Then at the end, auto commits what you have added.

For a multiline commit message, when you're asked for it, just write the new line with the flag `-m`.

For example:
>`Line 1 -m Line 2 -m Line 3 (and so on...)`

### cleanup

Reverts the aliases set.

## Env

You can edit the [`.env`](./.env) file to avoid having to type the same ticket id on every commit :)

Something like:
>`TICKET_PREFIX=HI-0001`

It also supports files named

- `.env.local`
- `.local.env`

In case `.env` is committed and can't be altered.

## Todo

- [x] Add multiple multi line support.

Parse all the `-m` options passed in the commit message and pass them correctly as args to the command call.\
Right now it just parses two lines, title and description (single line)

- [x] Add ticket prefix in env (TICKET_PREFIX=) for the cases where we're committing for the same ticket several times. This will avoid having to type it all the time

It would still show in the editor, but would have the env value as the default value (leaving it to the user to change it if needed)
