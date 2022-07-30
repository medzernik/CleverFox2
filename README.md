# CleverFox2
The second generation of the CleverFox bot

This is the main setup point for the Discord bot.

For the correct functioning, the core package group is required (main, config, logging and command packages)

You can configure using additional modules in the config or on the fly using the bot.

Additional modules will become available as development progresses.

# Modules

The module system is created so you can choose which parts of the bot are active.


## Core modules
The core modules cannot be disabled.

The core modules include:
- **Main** - the main module that initializes the bot.
- **Config** - the module that reads, stores and writes the configuration of the bot
- **Logging** - the logging module to a file, or stdout if the file isn't able to be created.
- **Command** - the interaction with the bot using Discord command interface. All the valid commands are always
  available in Discord, the enabled check is done at execution time.
- **Spinner** - a CLI spinner STDOUT compatible dynamic status updater.
- **Tviewsystem** - A CLI GUI (if a supported terminal is available).

These modules cannot be disabled, although they can be limited, they will always be used by some parts of the bot.

Tviewsystem disables itself when the bot fails to initialize the terminal interface.

Administration commands via the Tviewsystem GUI are extremely limited. A possible force enable setting may be added
later, though controlling administrative settings from within the bot should not work. The only interface for admin work
should be an admin user interacting with the bot from within Discord.

## Administration

Banning, kicking, purging
(This should only be available from within a server application command, not from Tviewsystem for privacy and security
reasons)

## COVID info
For now planned only for Slovakia, but possibly may include other countries from a different database.
