#!/usr/bin/env node

const { Command, Option } = require('commander');
const pkg = require('./package.json');
const ascii_txt_gen = require("ascii-text-generator");

const shortenCommand = require("./commands/shorten/app");
const unshortenCommand = require("./commands/unshorten/app");
const iviCommand = require("./commands/isValidInstance/app");

const program = new Command("quecto");

program
  .version(pkg.version)
  .addHelpText("before", ascii_txt_gen("Quecto", "2"))
  .description("Quecto is a CLI tool for use the Quecto API.")
  .addCommand(shortenCommand)
  .addCommand(unshortenCommand)
  .addCommand(iviCommand)
  .parseAsync().catch((e) => {
    console.log(e);
    process.exit(1)
  })