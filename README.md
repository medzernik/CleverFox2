# CleverFox2
The second generation of the CleverFox bot

This is the main setup point for the Discord bot.

For the correct functioning, the core package group is required (main, config, logging and command packages)

You can configure using additional modules in the config or on the fly using the bot.

Additional modules will become available as development progresses.

# Modules
## Core modules
The core modules cannot be disabled.

The core modules include:
- **Main** - the main module that initializes the bot.
- **Config** - the module that reads, stores and writes the configuration of the bot
- **Logging** - the logging module to a file, or stdout if the file isn't able to be created.
- **Command** - the interaction with the bot using Discord command interface. All the valid commands are always available in Discord, the enabled check is done at execution time.