# Meeseeks

## About this plugin

This plugin is a template and a functioning example for a basic JFrog CLI plugin.
This README shows the expected structure of your plugin's README.

## Installation with JFrog CLI

Installing the latest version:

`$ jfrog plugin install meeseeks`

Installing a specific version:

`$ jfrog plugin install meeseeks@version`

Uninstalling a plugin

`$ jfrog plugin uninstall meeseeks`

## Usage

### Commands

* info

  * Arguments:
    * addressee - The name of the person you would like to greet.
  * Flags:
        *shout: Makes output uppercase **[Default: false]**
        *repeat: Greets multiple times **[Default: 1]**
    * Example:

### Environment variables

* HELLO_FROG_GREET_PREFIX - Adds a prefix to every greet **[Default: New greeting: ]**

## Additional info

None.

## Release Notes

The release notes are available [here](RELEASE.md).
