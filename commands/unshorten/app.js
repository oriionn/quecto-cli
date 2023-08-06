const {createCommand} = require("commander");
const axios = require("axios");
const colors = require("colors");

function unshorten(link, options) {

  if (!options.instance) options.instance = "https://s.oriondev.fr";
  if (!options.instance.startsWith("http")) options.instance = `https://${options.instance}`;

  axios.get(`${options.instance}/api/quectoCheck`).then((res) => {
    if (res.data.data.quecto) {
      let apiUrl = `${options.instance}/api/s/${link.split("/")[link.split("/").length - 1]}`;
      if (options.password) apiUrl += `?password=${options.password}`;
      axios.get(apiUrl, {}).then((res) => {
        console.log(`Unshortened link: ${res.data.data.original}`.brightGreen);
      }).catch((e) => {
        console.log("This link is not a from this instance.".brightRed);
      })
    } else {
      console.log("This instance is not a Quecto instance.".brightRed);
    }
  }).catch((e) => {
    console.log("This instance is not a Quecto instance.".brightRed);
  })
}

module.exports = createCommand("unshorten")
    .description("Unshorten a URL")
    .argument("<link>", "The link to unshorten")
    .option("-i, --instance <instance>", "The instance to use")
    .option("-p, --password <password>", "The password for shortened link")
    .action(unshorten)